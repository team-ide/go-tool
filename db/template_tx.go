package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"strings"
	"sync"
)

var (
	contextDbTxName = new(string)
)

func init() {
	*contextDbTxName = "db:template:tx"
}

func OpenTxContext(ctx context.Context) (res context.Context, err error) {
	res = ctx

	txInfo := getTxInfo(res)
	if txInfo == nil {
		txInfo = &TxInfo{
			ctx: res,
		}
		res = context.WithValue(res, contextDbTxName, txInfo)
		txInfo.ctx = res
	}
	txInfo.openNumber++
	return
}

func openTx(ctx context.Context, sqlDB *sql.DB) (tx *sql.Tx, txInfo *TxInfo, err error) {
	txInfo = getTxInfo(ctx)
	if txInfo == nil {
		txInfo = &TxInfo{
			ctx: ctx,
		}
	}
	tx, err = txInfo.openTx(sqlDB)
	if err != nil {
		return
	}
	return
}

func EndTxContext(ctx context.Context, rollback bool) (err error) {
	txInfo := getTxInfo(ctx)
	if txInfo == nil {
		err = errors.New("上下文中，事务不存在")
		util.Logger.Error("tx end error", zap.Error(err))
		return
	}
	_, err = txInfo.end(rollback)
	return
}

func getTx(ctx context.Context, sqlDB *sql.DB) (tx *sql.Tx) {
	if ctx.Value(sqlDB) == nil {
		return
	}
	tx, ok := ctx.Value(sqlDB).(*sql.Tx)
	if !ok {
		tx = nil
	}
	return
}

func getTxInfo(ctx context.Context) (txInfo *TxInfo) {
	find := ctx.Value(contextDbTxName)
	if find == nil {
		return
	}
	txInfo, ok := find.(*TxInfo)
	if !ok {
		txInfo = nil
	}
	return
}

type TxInfo struct {
	openNumber  int
	closeNumber int
	txs         []*sql.Tx
	dbs         []*sql.DB
	ctx         context.Context
	txLocker    sync.Mutex
}

func (this_ *TxInfo) openTx(sqlDB *sql.DB) (tx *sql.Tx, err error) {
	this_.txLocker.Lock()
	defer this_.txLocker.Unlock()

	for i, db_ := range this_.dbs {
		if db_ == sqlDB {
			tx = this_.txs[i]
			break
		}
	}
	this_.openNumber++
	if tx == nil {
		tx, err = sqlDB.BeginTx(this_.ctx, &sql.TxOptions{})
		if err != nil {
			return
		}
		this_.dbs = append(this_.dbs, sqlDB)
		this_.txs = append(this_.txs, tx)
	}
	return
}

func (this_ *TxInfo) end(rollback bool) (isEnd bool, err error) {
	this_.txLocker.Lock()
	defer this_.txLocker.Unlock()

	this_.closeNumber++
	if this_.openNumber != this_.closeNumber {
		util.Logger.Debug("事务结束，不做任何事情，开启和关闭次数不一致", zap.Any("open", this_.openNumber), zap.Any("close", this_.closeNumber))
		return
	}

	txs := this_.txs

	this_.txs = make([]*sql.Tx, 0)
	this_.dbs = make([]*sql.DB, 0)
	this_.openNumber = 0
	this_.closeNumber = 0
	if txs == nil || len(txs) == 0 {
		util.Logger.Debug("事务结束，未开启任何数据库事务")
		return
	}
	util.Logger.Debug("事务结束，执行事务关闭", zap.Any("rollback", rollback))
	var errs []string
	for _, tx := range txs {
		if rollback {
			util.Logger.Debug("事务结束，执行 事务 回滚")
			err = tx.Rollback()
			if err != nil {
				util.Logger.Error("事务结束，执行 事务 回滚 异常", zap.Error(err))
				errs = append(errs, err.Error())
			} else {
				util.Logger.Debug("事务结束，执行 事务 回滚 成功")
			}
		} else {
			util.Logger.Debug("事务结束，执行 事务 提交")
			err = tx.Commit()
			if err != nil {
				util.Logger.Error("事务结束，执行 事务 提交 异常", zap.Error(err))
				errs = append(errs, err.Error())
			} else {
				util.Logger.Debug("事务结束，执行 事务 提交 成功")
			}
		}
	}
	if len(errs) > 0 {
		err = errors.New("事务结束，执行 异常:" + strings.Join(errs, ";"))
		return
	}
	return
}
