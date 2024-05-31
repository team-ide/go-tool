package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
	"reflect"
	"strings"
)

func (this_ *Template[T]) QueryOne(querySql string, queryArgs []interface{}) (res T, err error) {
	ctx := context.Background()
	stmt, err := this_.Service.GetDb().PrepareContext(ctx, querySql)
	if err != nil {
		return
	}
	defer func() { _ = stmt.Close() }()

	rows, err := stmt.Query(queryArgs...)
	if err != nil {
		return
	}
	defer func() { _ = rows.Close() }()

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return
	}

	var find bool
	for rows.Next() {
		if find {
			err = errors.New("has more rows by query one")
			return
		}
		find = true
		one, values, fields := this_.GetStructValues(columnTypes)
		err = rows.Scan(values...)
		if err != nil {
			return
		}
		for i, value := range values {
			if fields[i] != nil {
				(*fields[i]).Set(reflect.ValueOf(value).Elem())
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

func (this_ *Template[T]) Query(querySql string, queryArgs []interface{}) (res []T, err error) {
	ctx := context.Background()

	stmt, err := this_.Service.GetDb().PrepareContext(ctx, querySql)
	if err != nil {
		return
	}
	defer func() { _ = stmt.Close() }()

	rows, err := stmt.Query(queryArgs...)
	if err != nil {
		return
	}
	defer func() { _ = rows.Close() }()
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return
	}
	for rows.Next() {
		one, values, fields := this_.GetStructValues(columnTypes)
		err = rows.Scan(values...)
		if err != nil {
			return
		}
		for i, value := range values {
			if fields[i] != nil {
				(*fields[i]).Set(reflect.ValueOf(value).Elem())
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
	*worker.Page
	List []T `json:"list"`
}

func (this_ *Template[T]) QueryPage(querySql string, queryArgs []interface{}, page *worker.Page) (res *Page[T], err error) {
	if page.PageSize < 1 {
		page.PageSize = 1
	}
	if page.PageNo < 1 {
		page.PageNo = 1
	}
	pageSize := page.PageSize
	pageNo := page.PageNo

	countSql, err := dialect.FormatCountSql(querySql)
	if err != nil {
		return
	}
	page.TotalCount, err = worker.DoQueryCount(this_.Service.GetDb(), countSql, queryArgs)
	if err != nil {
		return
	}
	page.TotalPage = (page.TotalCount + page.PageSize - 1) / page.PageSize
	// 如果查询的页码 大于 总页码 则不查询
	if pageNo > page.TotalPage {
		return
	}
	pageSql := this_.Service.GetDialect().PackPageSql(querySql, pageSize, pageNo)

	list, err := this_.Query(pageSql, queryArgs)
	if err != nil {
		return
	}

	res = &Page[T]{}

	res.Page = page
	res.List = list
	return
}

func (this_ *Template[T]) GetStructValues(columns []*sql.ColumnType) (res reflect.Value, values []interface{}, fields []*reflect.Value) {
	res = reflect.New(this_.objValueType)
	objV := res
	for objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}
	structColumns := this_.structColumns
	structColumnsLower := this_.structColumnsLower
	if structColumns == nil {
		structColumnsLower, structColumns = this_.GetStructColumn(this_.objValueType)
		this_.structColumns = structColumns
		this_.structColumnsLower = structColumnsLower
	}

	for _, column := range columns {
		var fieldColumn *FieldColumn
		if this_.StrictColumnName {
			fieldColumn = structColumns[column.Name()]
		} else {
			fieldColumn = structColumnsLower[strings.ToLower(column.Name())]
		}
		if fieldColumn != nil {
			fieldV := objV.Field(fieldColumn.Index)
			fieldVV := reflect.New(fieldV.Type())
			values = append(values, fieldVV.Interface())
			fields = append(fields, &fieldV)
		} else {
			values = append(values, new(interface{}))
			fields = append(fields, nil)
		}
	}
	return
}

type FieldColumn struct {
	reflect.StructField
	Index int
}

func (this_ *TemplateOptions) GetStructColumn(tOf reflect.Type) (structColumns map[string]*FieldColumn, structColumnsLower map[string]*FieldColumn) {
	this_.structCacheLock.Lock()
	defer this_.structCacheLock.Unlock()
	var ok bool
	structColumns, ok = this_.structColumnsCache[tOf]
	structColumnsLower, ok = this_.structColumnsLowerCache[tOf]
	if ok {
		//fmt.Println("find from cache")
		return
	}
	structColumns = map[string]*FieldColumn{}
	structColumnsLower = map[string]*FieldColumn{}
	for i := 0; i < tOf.NumField(); i++ {
		field := tOf.Field(i)
		fieldColumn := &FieldColumn{
			StructField: field,
			Index:       i,
		}
		var str string
		var key string
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
			key = ss[0]
		}
		if key != "" {
			structColumns[key] = fieldColumn
			structColumnsLower[strings.ToLower(key)] = fieldColumn
		}
	}
	this_.structColumnsCache[tOf] = structColumns
	this_.structColumnsLowerCache[tOf] = structColumnsLower
	return
}
