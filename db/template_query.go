package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"reflect"
	"strconv"
	"strings"
)

func (this_ *Template[T]) QueryOne(ctx context.Context, querySql string, queryArgs []interface{}) (res T, err error) {
	stmt, err := this_.Service.GetDb().PrepareContext(ctx, querySql)
	if err != nil {
		util.Logger.Error("query one error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
		return
	}
	defer func() { _ = stmt.Close() }()

	rows, err := stmt.Query(queryArgs...)
	if err != nil {
		util.Logger.Error("query one error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
		return
	}
	defer func() { _ = rows.Close() }()

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		util.Logger.Error("query one error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
		return
	}

	var find bool
	for rows.Next() {
		if find {
			err = errors.New("has more rows by query one")
			util.Logger.Error("query one error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
			return
		}
		find = true
		one, values, fields := this_.getValues(columnTypes)
		err = rows.Scan(values...)
		if err != nil {
			util.Logger.Error("query one error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
			return
		}
		if this_.objValueType.Kind() == reflect.Map {
			err = this_.fullMapValues(one, values, columnTypes)
			if err != nil {
				util.Logger.Error("full map values error", zap.Error(err))
				return
			}
		} else {
			err = this_.fullStructValues(one, values, columnTypes, fields)
			if err != nil {
				util.Logger.Error("full struct values error", zap.Error(err))
				return
			}
		}

		if this_.objType.Kind() == reflect.Ptr {
			res = one.Interface().(T)
		} else {
			res = one.Elem().Interface().(T)
		}
	}
	return
}

func (this_ *Template[T]) QueryList(ctx context.Context, querySql string, queryArgs []interface{}) (res []T, err error) {

	stmt, err := this_.Service.GetDb().PrepareContext(ctx, querySql)
	if err != nil {
		util.Logger.Error("query list error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
		return
	}
	defer func() { _ = stmt.Close() }()

	rows, err := stmt.Query(queryArgs...)
	if err != nil {
		util.Logger.Error("query list error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
		return
	}
	defer func() { _ = rows.Close() }()
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		util.Logger.Error("query list error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
		return
	}
	for rows.Next() {
		one, values, fields := this_.getValues(columnTypes)
		err = rows.Scan(values...)
		if err != nil {
			util.Logger.Error("query list error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
			return
		}
		if this_.objValueType.Kind() == reflect.Map {
			err = this_.fullMapValues(one, values, columnTypes)
			if err != nil {
				util.Logger.Error("full map values error", zap.Error(err))
				return
			}
		} else {
			err = this_.fullStructValues(one, values, columnTypes, fields)
			if err != nil {
				util.Logger.Error("full struct values error", zap.Error(err))
				return
			}
		}
		if this_.objType.Kind() == reflect.Ptr {
			res = append(res, one.Interface().(T))
		} else {
			res = append(res, one.Elem().Interface().(T))
		}

	}
	return
}

type Page[T any] struct {
	PageSize   int `json:"pageSize"`
	PageNo     int `json:"pageNo"`
	TotalCount int `json:"totalCount"`
	TotalPage  int `json:"totalPage"`
	List       []T `json:"list"`
}

func (this_ *Template[T]) QueryCount(ctx context.Context, querySql string, queryArgs []interface{}) (count int, err error) {

	stmt, err := this_.Service.GetDb().PrepareContext(ctx, querySql)
	if err != nil {
		util.Logger.Error("query count error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
		return
	}
	defer func() { _ = stmt.Close() }()

	rows, err := stmt.Query(queryArgs...)
	if err != nil {
		util.Logger.Error("query count error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
		return
	}
	defer func() { _ = rows.Close() }()
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			util.Logger.Error("query count error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
			return
		}
	}

	return
}

func (this_ *Template[T]) QueryPage(ctx context.Context, querySql string, queryArgs []interface{}, pageSize int, pageNo int) (res []T, err error) {
	if pageSize < 1 {
		pageSize = 1
	}
	if pageNo < 1 {
		pageNo = 1
	}
	pageSql := this_.Service.GetDialect().PackPageSql(querySql, pageSize, pageNo)

	res, err = this_.QueryList(ctx, pageSql, queryArgs)
	if err != nil {
		return
	}
	return
}

func (this_ *Template[T]) QueryPageBean(ctx context.Context, querySql string, queryArgs []interface{}, pageSize int, pageNo int) (res *Page[T], err error) {
	if pageSize < 1 {
		pageSize = 1
	}
	if pageNo < 1 {
		pageNo = 1
	}

	countSql, err := dialect.FormatCountSql(querySql)
	if err != nil {
		util.Logger.Error("query page bean error", zap.Any("sql", querySql), zap.Any("args", queryArgs), zap.Error(err))
		return
	}
	res = &Page[T]{}
	res.PageSize = pageSize
	res.PageNo = pageNo
	res.TotalCount, err = this_.QueryCount(ctx, countSql, queryArgs)
	if err != nil {
		return
	}
	res.TotalPage = (res.TotalCount + res.PageSize - 1) / res.PageSize
	// 如果查询的页码 大于 总页码 则不查询
	if pageNo > res.TotalPage {
		return
	}
	pageSql := this_.Service.GetDialect().PackPageSql(querySql, pageSize, pageNo)

	res.List, err = this_.QueryList(ctx, pageSql, queryArgs)
	if err != nil {
		return
	}

	return
}

func (this_ *Template[T]) fullMapValues(res reflect.Value, values []any, columns []*sql.ColumnType) (err error) {
	objV := res
	for objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}
	data := map[string]interface{}{}
	for i, column := range columns {
		name := column.Name()
		//key := reflect.ValueOf(name)
		v := values[i]
		vv := worker.GetSqlValue(column, v)
		data[name] = vv
	}
	objV.Set(reflect.ValueOf(data))
	return
}

func (this_ *Template[T]) fullStructValues(res reflect.Value, values []any, columns []*sql.ColumnType, fields []*reflect.Value) (err error) {

	for i, column := range columns {
		field := fields[i]
		if field == nil {
			continue
		}
		v := values[i]
		if v == nil {
			continue
		}
		vv := worker.GetSqlValue(column, v)
		if vv == nil {
			continue
		}
		valueTypeOf := reflect.TypeOf(vv)
		columnValueType := ""
		fieldV := *field
		for fieldV.Kind() == reflect.Ptr {
			fieldV = fieldV.Elem()
		}
		fieldT := fieldV.Type()
		for fieldT.Kind() == reflect.Ptr {
			fieldT = fieldT.Elem()
		}
		fieldType := field.Type().String()
		if valueTypeOf != nil {
			columnValueType = valueTypeOf.String()
		}
		if columnValueType != fieldType {
			fieldKind := fieldT.Kind()
			if fieldKind == reflect.String {
				fieldV.SetString(util.GetStringValue(vv))
			} else if fieldKind >= reflect.Int && fieldKind <= reflect.Int64 {
				str := util.GetStringValue(vv)
				var num int64
				if str != "" {
					num, err = strconv.ParseInt(str, 10, 64)
					if err != nil {
						return
					}
				}
				fieldV.SetInt(num)
			} else if fieldKind >= reflect.Uint && fieldKind <= reflect.Uint64 {
				str := util.GetStringValue(vv)
				var num uint64
				if str != "" {
					num, err = strconv.ParseUint(str, 10, 64)
					if err != nil {
						return
					}
				}
				fieldV.SetUint(num)
			} else if fieldKind == reflect.Float32 || fieldKind == reflect.Float64 {

				str := util.GetStringValue(vv)
				var num float64
				if str != "" {
					num, err = strconv.ParseFloat(str, 64)
					if err != nil {
						return
					}
				}
				fieldV.SetFloat(num)
			} else {
				fieldV.Set(reflect.ValueOf(vv))
			}
		} else {
			fieldV.Set(reflect.ValueOf(vv))
		}
	}
	return
}
func (this_ *Template[T]) getValues(columns []*sql.ColumnType) (res reflect.Value, values []interface{}, fields []*reflect.Value) {
	if this_.objValueType.Kind() == reflect.Map {
		res = reflect.New(this_.objValueType)
		// 处理 map
		for range columns {
			values = append(values, new(interface{}))
		}
		return
	}
	res = reflect.New(this_.objValueType)
	objV := res
	for objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}
	structInfo := this_.structInfo
	if structInfo == nil {
		structInfo = this_.GetStructInfo(this_.objValueType)
		this_.structInfo = structInfo
	}

	for _, column := range columns {
		var fieldColumn *FieldColumn
		if this_.StrictColumnName {
			fieldColumn = structInfo.structColumnMap[column.Name()]
		} else {
			fieldColumn = structInfo.structColumnLower[strings.ToLower(column.Name())]
		}
		if fieldColumn != nil {
			fieldV := objV.Field(fieldColumn.Index)
			fieldV.Set(reflect.New(fieldV.Type()).Elem())
			values = append(values, new(interface{}))
			fields = append(fields, &fieldV)
		} else {
			values = append(values, new(interface{}))
			fields = append(fields, nil)
		}
	}
	return
}
