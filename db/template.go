package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"reflect"
	"sync"
	"time"
)

func WarpTemplate[T any](t T, opts *TemplateOptions) (res *Template[T]) {
	if opts.DialectParam == nil {
		opts.DialectParam = &dialect.ParamModel{}
	}
	if opts.Dialect == nil {
		if opts.Service != nil {
			opts.Dialect = opts.Service.GetDialect()
		} else {
			opts.Dialect, _ = dialect.NewDialect("mysql")
		}
	}
	res = &Template[T]{}
	res.TemplateOptions = opts
	res.t = t
	res.objType = reflect.TypeOf(t)

	res.objValueType = res.objType
	for res.objValueType.Kind() == reflect.Ptr {
		res.objValueType = res.objValueType.Elem()
	}
	return
}

type TemplateOptions struct {
	Dialect               dialect.Dialect
	DialectParam          *dialect.ParamModel
	Service               IService
	ColumnTagName         string `json:"columnTagName"`     // 字段 tag 注解名称 默认 `column:"xx"`
	NotUseJsonTagName     bool   `json:"notUseJsonTagName"` // 如果 字段 tag 注解 未找到 使用 json 注解
	NotUseFieldName       bool   `json:"notUseFieldName"`   // 如果 以上都未配置 使用 字段名称
	StrictColumnName      bool   `json:"strictColumnName"`  // 严格的字段大小写 默认 false `userId 将 匹配 Userid UserId USERID`
	structInfoCache       map[reflect.Type]*StructInfo
	structInfoCacheLock   sync.Mutex
	StringEmptyNotUseNull bool `json:"stringEmptyNotUseNull"` // 字段 string 类型 如果是 空字段 则设置为 null
	NumberZeroNotUseNull  bool `json:"numberZeroNotUseNull"`  // 数字 类型 如果是 0 则设置为 null

	GetColumnValue func(columnType *sql.ColumnType, value any) (res any, err error)                                             `json:"-"`
	SetFieldValue  func(columnType *sql.ColumnType, field reflect.StructField, fieldValue reflect.Value, value any) (err error) `json:"-"`
}

type StructInfo struct {
	isMap             bool
	structColumns     []*FieldColumn
	structColumnMap   map[string]*FieldColumn
	structColumnLower map[string]*FieldColumn
}

type FieldColumn struct {
	Field      reflect.StructField
	Index      int
	ColumnName string
	IsString   bool
	IsNumber   bool
	IsBool     bool
	value      *reflect.Value
}
type Template[T any] struct {
	*TemplateOptions
	t            T
	objType      reflect.Type
	objValueType reflect.Type
	structInfo   *StructInfo
}

func (this_ *Template[T]) SelectOne(ctx context.Context, sqlParamSql string, sqlParam any) (res T, err error) {
	selectSql, selectArgs, err := this_.SqlParamParser(sqlParamSql, sqlParam).Parse()
	if err != nil {
		util.Logger.Error("sql parse error", zap.Any("sql", sqlParamSql), zap.Any("param", sqlParam), zap.Error(err))
		return
	}
	res, err = this_.QueryOne(ctx, selectSql, selectArgs)
	return
}

func (this_ *Template[T]) SelectList(ctx context.Context, sqlParamSql string, sqlParam any) (res []T, err error) {
	selectSql, selectArgs, err := this_.SqlParamParser(sqlParamSql, sqlParam).Parse()
	if err != nil {
		util.Logger.Error("sql parse error", zap.Any("sql", sqlParamSql), zap.Any("param", sqlParam), zap.Error(err))
		return
	}
	res, err = this_.QueryList(ctx, selectSql, selectArgs)
	return
}

func (this_ *Template[T]) SelectPage(ctx context.Context, sqlParamSql string, sqlParam any, pageSize int, pageNo int) (res []T, err error) {
	selectSql, selectArgs, err := this_.SqlParamParser(sqlParamSql, sqlParam).Parse()
	if err != nil {
		util.Logger.Error("sql parse error", zap.Any("sql", sqlParamSql), zap.Any("param", sqlParam), zap.Error(err))
		return
	}
	res, err = this_.QueryPage(ctx, selectSql, selectArgs, pageSize, pageNo)
	return
}

func (this_ *Template[T]) SelectPageBean(ctx context.Context, sqlParamSql string, sqlParam any, pageSize int, pageNo int) (res *Page[T], err error) {
	selectSql, selectArgs, err := this_.SqlParamParser(sqlParamSql, sqlParam).Parse()
	if err != nil {
		util.Logger.Error("sql parse error", zap.Any("sql", sqlParamSql), zap.Any("param", sqlParam), zap.Error(err))
		return
	}
	res, err = this_.QueryPageBean(ctx, selectSql, selectArgs, pageSize, pageNo)
	return
}

func (this_ *Template[T]) Insert(ctx context.Context, table string, obj any, ignoreColumns ...string) (res int64, err error) {
	insertSql, insertArgs := this_.GetInsertSql(table, obj, ignoreColumns...)
	res, err = this_.Exec(ctx, insertSql, insertArgs)
	return
}

func (this_ *Template[T]) BatchInsert(ctx context.Context, table string, obj []any, ignoreColumns ...string) (res int64, err error) {
	insertSql, insertArgsList := this_.GetBatchInsertSql(table, obj, ignoreColumns...)
	res, err = this_.ExecList(ctx, insertSql, insertArgsList)
	return
}

func (this_ *Template[T]) Update(ctx context.Context, table string, obj any, whereSql string, whereParam any, ignoreColumns ...string) (res int64, err error) {
	if whereSql == "" {
		err = errors.New("更新数据必须传入条件")
		util.Logger.Error("update error", zap.Error(err))
		return
	}
	updateSql, updateArgs, err := this_.GetUpdateSql(table, obj, this_.SqlParamParser(whereSql, whereParam), ignoreColumns...)
	if err != nil {
		util.Logger.Error("get update sql error", zap.Error(err))
		return
	}
	res, err = this_.Exec(ctx, updateSql, updateArgs)
	return
}

func (this_ *Template[T]) Delete(ctx context.Context, table string, whereSql string, whereParam any) (res int64, err error) {
	if whereSql == "" {
		err = errors.New("删除据必须传入条件")
		util.Logger.Error("delete error", zap.Error(err))
		return
	}
	deleteSql, deleteArgs, err := this_.GetDeleteSql(table, this_.SqlParamParser(whereSql, whereParam))
	if err != nil {
		util.Logger.Error("get delete sql error", zap.Error(err))
		return
	}
	res, err = this_.Exec(ctx, deleteSql, deleteArgs)
	return
}

func (this_ *Template[T]) Exec(ctx context.Context, execSql string, execArgs []any) (res int64, err error) {
	sqlDB := this_.Service.GetDb()
	tx, txInfo, err := openTx(ctx, sqlDB)
	if err != nil {
		return
	}
	defer func() {
		_, e := txInfo.end(err != nil)
		if err == nil && e != nil {
			err = e
		}
	}()

	util.Logger.Debug("exec start", zap.Any("sql", execSql), zap.Any("args", execArgs))
	var startTime = time.Now()
	result, err := tx.ExecContext(ctx, execSql, execArgs...)
	if err != nil {
		util.Logger.Error("exec error", zap.Any("sql", execSql), zap.Any("args", execArgs), zap.Error(err))
		return
	}
	var endTime = time.Now()
	util.Logger.Debug("exec end", zap.Any("sql", execSql), zap.Any("args", execArgs), zap.Any("useTime", endTime.UnixMilli()-startTime.UnixMilli()))
	res, err = result.RowsAffected()
	if err != nil {
		util.Logger.Error("exec get rows affected error", zap.Any("sql", execSql), zap.Any("args", execArgs), zap.Error(err))
		return
	}
	return
}

func (this_ *Template[T]) ExecList(ctx context.Context, execSql string, execArgsList [][]any) (res int64, err error) {
	sqlDB := this_.Service.GetDb()
	tx, txInfo, err := openTx(ctx, sqlDB)
	if err != nil {
		return
	}
	defer func() {
		_, e := txInfo.end(err != nil)
		if err == nil && e != nil {
			err = e
		}
	}()

	stmt, err := tx.PrepareContext(ctx, execSql)
	if err != nil {
		util.Logger.Error("exec list error", zap.Any("sql", execSql), zap.Error(err))
		return
	}
	defer func() { _ = stmt.Close() }()

	var result sql.Result
	for _, execArgs := range execArgsList {

		result, err = stmt.Exec(execArgs...)
		if err != nil {
			util.Logger.Error("exec error", zap.Any("sql", execSql), zap.Any("args", execArgs), zap.Error(err))
			return
		}
		rowsAffected, _ := result.RowsAffected()
		res += rowsAffected
	}
	return
}
