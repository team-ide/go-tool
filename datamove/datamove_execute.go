package datamove

import (
	"errors"
	"fmt"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func (this_ *Executor) execute() (err error) {

	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
			util.Logger.Error("execute panic error", zap.Error(err))
		}
		if this_.From != nil && this_.From.DbConfig != nil && this_.From.DbConfig.SSHClient != nil {
			_ = this_.From.DbConfig.SSHClient.Close()
		}
		if this_.To != nil && this_.To.DbConfig != nil && this_.To.DbConfig.SSHClient != nil {
			_ = this_.To.DbConfig.SSHClient.Close()
		}

		if this_.From != nil && this_.From.RedisConfig != nil && this_.From.RedisConfig.SSHClient != nil {
			_ = this_.From.RedisConfig.SSHClient.Close()
		}
		if this_.To != nil && this_.To.RedisConfig != nil && this_.To.RedisConfig.SSHClient != nil {
			_ = this_.To.RedisConfig.SSHClient.Close()
		}
	}()

	if this_.From.IsDb() {
		if this_.To.IsSql() {
			err = this_.dbToSql()
		} else if this_.To.IsExcel() {
			err = this_.dbToExcel()
		} else if this_.To.IsTxt() {
			err = this_.dbToTxt()
		} else if this_.To.IsDb() {
			err = this_.dbToDb()
		} else if this_.To.IsEs() {
			err = this_.dbToEs()
		} else {
			err = errors.New(fmt.Sprintf("不支持的 目标 类型[%s]", this_.To.Type))
			util.Logger.Error("execute error", zap.Error(err))
			return
		}
	} else if this_.From.IsSql() {
		if this_.To.IsDb() {
			err = this_.sqlToDb()
		} else {
			err = errors.New(fmt.Sprintf("不支持的 目标 类型[%s]", this_.To.Type))
			util.Logger.Error("execute error", zap.Error(err))
			return
		}
	} else if this_.From.IsData() {
		if this_.To.IsSql() {
			err = this_.dataToSql()
		} else if this_.To.IsExcel() {
			err = this_.dataToExcel()
		} else if this_.To.IsTxt() {
			err = this_.dataToTxt()
		} else if this_.To.IsDb() {
			err = this_.dataToDb()
		} else if this_.To.IsEs() {
			err = this_.dataToEs()
		} else {
			err = errors.New(fmt.Sprintf("不支持的 目标 类型[%s]", this_.To.Type))
			util.Logger.Error("execute error", zap.Error(err))
			return
		}
	} else if this_.From.IsTxt() {
		if this_.To.IsSql() {
			err = this_.txtToSql()
		} else if this_.To.IsExcel() {
			err = this_.txtToExcel()
		} else if this_.To.IsTxt() {
			err = this_.txtToTxt()
		} else if this_.To.IsDb() {
			err = this_.txtToDb()
		} else if this_.To.IsEs() {
			err = this_.txtToEs()
		} else {
			err = errors.New(fmt.Sprintf("不支持的 目标 类型[%s]", this_.To.Type))
			util.Logger.Error("execute error", zap.Error(err))
			return
		}
	} else if this_.From.IsExcel() {
		if this_.To.IsSql() {
			err = this_.excelToSql()
		} else if this_.To.IsExcel() {
			err = this_.excelToExcel()
		} else if this_.To.IsTxt() {
			err = this_.excelToTxt()
		} else if this_.To.IsDb() {
			err = this_.excelToDb()
		} else if this_.To.IsEs() {
			err = this_.excelToEs()
		} else {
			err = errors.New(fmt.Sprintf("不支持的 目标 类型[%s]", this_.To.Type))
			util.Logger.Error("execute error", zap.Error(err))
			return
		}
	} else if this_.From.IsEs() {
		if this_.To.IsSql() {
			err = this_.esToSql()
		} else if this_.To.IsExcel() {
			err = this_.esToExcel()
		} else if this_.To.IsTxt() {
			err = this_.esToTxt()
		} else if this_.To.IsDb() {
			err = this_.esToDb()
		} else if this_.To.IsEs() {
			err = this_.esToEs()
		} else {
			err = errors.New(fmt.Sprintf("不支持的 目标 类型[%s]", this_.To.Type))
			util.Logger.Error("execute error", zap.Error(err))
			return
		}
	} else if this_.From.IsKafka() {
		if this_.To.IsSql() {
			err = this_.kafkaToSql()
		} else if this_.To.IsExcel() {
			err = this_.kafkaToExcel()
		} else if this_.To.IsTxt() {
			err = this_.kafkaToTxt()
		} else if this_.To.IsDb() {
			err = this_.kafkaToDb()
		} else if this_.To.IsEs() {
			err = this_.kafkaToEs()
		} else {
			err = errors.New(fmt.Sprintf("不支持的 目标 类型[%s]", this_.To.Type))
			util.Logger.Error("execute error", zap.Error(err))
			return
		}
	} else if this_.From.IsScript() {
		if this_.To.IsSql() {
			err = this_.scriptToSql()
		} else if this_.To.IsExcel() {
			err = this_.scriptToExcel()
		} else if this_.To.IsTxt() {
			err = this_.scriptToTxt()
		} else if this_.To.IsDb() {
			err = this_.scriptToDb()
		} else if this_.To.IsEs() {
			err = this_.scriptToEs()
		} else {
			err = errors.New(fmt.Sprintf("不支持的 目标 类型[%s]", this_.To.Type))
			util.Logger.Error("execute error", zap.Error(err))
			return
		}
	} else {
		err = errors.New(fmt.Sprintf("不支持的 源 类型[%s]", this_.From.Type))
		util.Logger.Error("execute error", zap.Error(err))
		return
	}
	return
}
