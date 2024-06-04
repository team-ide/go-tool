package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"reflect"
	"strings"
	"sync"
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
	if opts.contextDbTxName == "" {
		opts.contextDbTxName = "db:tx"
	}
	if opts.contextDbTxOpenNumberName == "" {
		opts.contextDbTxOpenNumberName = "db:tx:open:number"
	}
	if opts.contextDbTxCloseNumberName == "" {
		opts.contextDbTxCloseNumberName = "db:tx:close:number"
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
	Dialect             dialect.Dialect
	DialectParam        *dialect.ParamModel
	Service             IService
	ColumnTagName       string `json:"columnTagName"`    // 字段 tag 注解名称 默认 `column:"xx"`
	UseJsonTagName      bool   `json:"useJsonTagName"`   // 如果 字段 tag 注解 未找到 使用 json 注解 默认为 false
	UseFieldName        bool   `json:"useFieldName"`     // 如果 以上都未配置 使用 字段名称 默认为 false
	StrictColumnName    bool   `json:"strictColumnName"` // 严格的字段大小写 默认 false `userId 将 匹配 Userid UserId USERID`
	structInfoCache     map[reflect.Type]*StructInfo
	structInfoCacheLock sync.Mutex
	StringEmptyUseNull  bool `json:"stringEmptyUseNull"`
	NumberZeroUseNull   bool `json:"numberZeroUseNull"`

	contextDbTxName            string
	contextDbTxOpenNumberName  string
	contextDbTxCloseNumberName string
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
	ctx, err = this_.OpenTxContext(ctx)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = this_.TxRollback(ctx)
		} else {
			_ = this_.TxCommit(ctx)
		}
	}()
	tx := this_.getTx(ctx)

	result, err := tx.ExecContext(ctx, execSql, execArgs...)
	if err != nil {
		util.Logger.Error("exec error", zap.Any("sql", execSql), zap.Any("args", execArgs), zap.Error(err))
		return
	}
	res, err = result.RowsAffected()
	if err != nil {
		util.Logger.Error("exec get rows affected error", zap.Any("sql", execSql), zap.Any("args", execArgs), zap.Error(err))
		return
	}
	return
}

func (this_ *Template[T]) ExecList(ctx context.Context, execSql string, execArgsList [][]any) (res int64, err error) {
	ctx, err = this_.OpenTxContext(ctx)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			_ = this_.TxRollback(ctx)
		} else {
			_ = this_.TxCommit(ctx)
		}
	}()
	tx := this_.getTx(ctx)

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

func (this_ *Template[T]) OpenTxContext(ctx context.Context) (res context.Context, err error) {
	res = ctx
	if this_.getTx(res) == nil {
		var tx *sql.Tx
		tx, err = this_.Service.GetDb().BeginTx(res, &sql.TxOptions{})
		if err != nil {
			util.Logger.Error("begin tx error", zap.Error(err))
			return
		}
		res = context.WithValue(res, this_.contextDbTxName, tx)
	}
	if this_.getTxOpenNumber(res) == nil {
		var num = new(int)
		*num = 1
		res = context.WithValue(res, this_.contextDbTxOpenNumberName, num)
	} else {
		*this_.getTxOpenNumber(res)++
	}

	if this_.getTxCloseNumber(res) == nil {
		var num = new(int)
		*num = 0
		res = context.WithValue(res, this_.contextDbTxCloseNumberName, num)
	}
	return
}

func (this_ *Template[T]) TxCommit(ctx context.Context) (err error) {
	tx := this_.getTx(ctx)
	if tx == nil {
		err = errors.New("上下文中，事务不存在")
		util.Logger.Error("tx commit error", zap.Error(err))
		return
	}
	openNum := this_.getTxOpenNumber(ctx)
	if openNum == nil {
		err = errors.New("上下文中，打开次数不存在")
		util.Logger.Error("tx commit error", zap.Error(err))
		return
	}
	closeNum := this_.getTxCloseNumber(ctx)
	if closeNum == nil {
		err = errors.New("上下文中，关闭次数不存在")
		util.Logger.Error("tx commit error", zap.Error(err))
		return
	}
	*closeNum++
	if *openNum == *closeNum {
		util.Logger.Debug("tx commit start")
		err = tx.Commit()
		if err != nil {
			util.Logger.Error("tx commit error", zap.Error(err))
			return
		} else {
			util.Logger.Debug("tx commit success")
		}
	}
	return
}

func (this_ *Template[T]) TxRollback(ctx context.Context) (err error) {
	tx := this_.getTx(ctx)
	if tx == nil {
		err = errors.New("上下文中，事务不存在")
		util.Logger.Error("tx rollback error", zap.Error(err))
		return
	}
	openNum := this_.getTxOpenNumber(ctx)
	if openNum == nil {
		err = errors.New("上下文中，打开次数不存在")
		util.Logger.Error("tx rollback error", zap.Error(err))
		return
	}
	closeNum := this_.getTxCloseNumber(ctx)
	if closeNum == nil {
		err = errors.New("上下文中，关闭次数不存在")
		util.Logger.Error("tx rollback error", zap.Error(err))
		return
	}
	*closeNum++
	if *openNum == *closeNum {
		util.Logger.Debug("tx rollback start")
		err = tx.Rollback()
		if err != nil {
			util.Logger.Error("tx rollback error", zap.Error(err))
			return
		} else {
			util.Logger.Debug("tx rollback success")
		}
	}
	return
}

func (this_ *Template[T]) getTx(ctx context.Context) (tx *sql.Tx) {
	if ctx.Value(this_.contextDbTxName) == nil {
		return
	}
	tx, ok := ctx.Value(this_.contextDbTxName).(*sql.Tx)
	if !ok {
		tx = nil
	}
	return
}

func (this_ *Template[T]) getTxOpenNumber(ctx context.Context) (res *int) {
	if ctx.Value(this_.contextDbTxOpenNumberName) == nil {
		return
	}
	num, ok := ctx.Value(this_.contextDbTxOpenNumberName).(*int)
	if !ok {
		return
	}
	res = num
	return
}

func (this_ *Template[T]) getTxCloseNumber(ctx context.Context) (res *int) {
	if ctx.Value(this_.contextDbTxCloseNumberName) == nil {
		return
	}
	num, ok := ctx.Value(this_.contextDbTxCloseNumberName).(*int)
	if !ok {
		return
	}
	res = num
	return
}

func (this_ *TemplateOptions) GetStructInfo(structType reflect.Type) (info *StructInfo) {
	for structType.Kind() == reflect.Ptr {
		structType = structType.Elem()
	}
	if structType.Kind() == reflect.Map {
		info = &StructInfo{
			isMap: true,
		}
		return
	}

	this_.structInfoCacheLock.Lock()
	if this_.structInfoCache == nil {
		this_.structInfoCache = map[reflect.Type]*StructInfo{}
	}
	var ok bool
	info, ok = this_.structInfoCache[structType]
	if ok {
		this_.structInfoCacheLock.Unlock()
		return
	}
	defer this_.structInfoCacheLock.Unlock()
	info = &StructInfo{}
	this_.structInfoCache[structType] = info
	info.structColumnMap = map[string]*FieldColumn{}
	info.structColumnLower = map[string]*FieldColumn{}
	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		fieldColumn := &FieldColumn{
			Field: field,
			Index: i,
		}
		var str string
		var columnName string
		var tag = this_.ColumnTagName
		if tag == "" {
			tag = "column"
		}
		str = field.Tag.Get(tag)
		if str == "" && this_.UseJsonTagName {
			str = field.Tag.Get("json")
		}
		if str == "" && this_.UseFieldName {
			str = field.Name
		}
		if str != "" && str != "-" {
			ss := strings.Split(str, ",")
			columnName = ss[0]
		}
		if columnName != "" {
			fT := field.Type
			fTKind := fT.Kind()
			for fTKind == reflect.Ptr {
				fT = fT.Elem()
				fTKind = fT.Kind()
			}
			fieldColumn.IsString = fTKind == reflect.String
			fieldColumn.IsNumber = fTKind >= reflect.Int && fTKind <= reflect.Uint64
			fieldColumn.IsBool = fTKind == reflect.Bool
			fieldColumn.ColumnName = columnName
			info.structColumns = append(info.structColumns, fieldColumn)
			info.structColumnMap[columnName] = fieldColumn
			info.structColumnLower[strings.ToLower(columnName)] = fieldColumn
		}
	}
	return
}

func (this_ *TemplateOptions) GetMapColumns(v any) (fieldColumns []*FieldColumn) {
	objV := reflect.ValueOf(v)
	for objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}

	for _, kV := range objV.MapKeys() {
		if kV.Type().Kind() != reflect.String {
			continue
		}
		k := kV.String()
		vV := objV.MapIndex(kV)
		fieldColumn := &FieldColumn{
			ColumnName: k,
			value:      &vV,
		}
		fT := vV.Type()
		fTKind := fT.Kind()
		for fTKind == reflect.Ptr {
			fT = fT.Elem()
			fTKind = fT.Kind()
		}
		this_.fullFieldValue(fieldColumn, fTKind)

		fieldColumns = append(fieldColumns, fieldColumn)
	}
	return
}

func (this_ *TemplateOptions) fullFieldValue(fieldColumn *FieldColumn, fTKind reflect.Kind) {
	fieldColumn.IsString = fTKind == reflect.String
	fieldColumn.IsNumber = fTKind >= reflect.Int && fTKind <= reflect.Uint64
	fieldColumn.IsBool = fTKind == reflect.Bool

}
