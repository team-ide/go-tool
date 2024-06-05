package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

func DoQueryStructs(db *sql.DB, sqlInfo string, args []interface{}, list interface{}) (err error) {
	ctx := context.Background()

	stmt, err := db.PrepareContext(ctx, sqlInfo)
	if err != nil {
		return
	}
	defer func() { _ = stmt.Close() }()

	rows, err := stmt.Query(args...)
	if err != nil {
		return
	}
	defer func() { _ = rows.Close() }()
	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return
	}
	listVOf := reflect.ValueOf(list).Elem()
	listStrType := GetListStructType(list)
	for rows.Next() {
		var values []interface{}
		for range columnTypes {
			values = append(values, new(interface{}))
		}
		err = rows.Scan(values...)
		if err != nil {
			return
		}

		item := make(map[string]interface{})
		for index, data := range values {
			item[columnTypes[index].Name()] = worker.GetSqlValue(columnTypes[index], data)
		}
		listStrValue := reflect.New(listStrType)
		SetStructColumnValues(item, listStrValue.Elem())
		listVOf = reflect.Append(listVOf, listStrValue)
	}
	reflect.ValueOf(list).Elem().Set(listVOf)
	return
}

func DoQueryStruct(db *sql.DB, sqlInfo string, args []interface{}, str interface{}) (find bool, err error) {
	ctx := context.Background()
	stmt, err := db.PrepareContext(ctx, sqlInfo)
	if err != nil {
		return
	}
	defer func() { _ = stmt.Close() }()

	rows, err := stmt.Query(args...)
	if err != nil {
		return
	}
	defer func() { _ = rows.Close() }()

	columnTypes, err := rows.ColumnTypes()
	if err != nil {
		return
	}
	strVOf := reflect.ValueOf(str)

	var isBase bool
	switch str.(type) {
	case *int, *int8, *int16, *int32, *int64, *float32, *float64:
		isBase = true
		break
	}
	for rows.Next() {
		if find {
			err = errors.New("has more rows by query one")
			return
		}
		find = true
		var values []interface{}
		if isBase {
			values = []interface{}{str}
		} else {
			for range columnTypes {
				values = append(values, new(interface{}))
			}
		}
		err = rows.Scan(values...)
		if err != nil {
			return
		}
		if isBase {
			continue
		}
		item := make(map[string]interface{})
		for index, data := range values {
			item[columnTypes[index].Name()] = worker.GetSqlValue(columnTypes[index], data)
		}
		SetStructColumnValues(item, strVOf.Elem())
	}
	return
}

var (
	structFieldMapCache  = map[reflect.Type]map[string]reflect.StructField{}
	structColumnMapCache = map[reflect.Type]map[string]reflect.StructField{}
	structMapLock        sync.Mutex
)

func getStructColumn(tOf reflect.Type) (structFieldMap map[string]reflect.StructField, structColumnMap map[string]reflect.StructField) {
	structMapLock.Lock()
	defer structMapLock.Unlock()
	structFieldMap, ok := structFieldMapCache[tOf]
	structColumnMap = structColumnMapCache[tOf]
	if ok {
		//fmt.Println("find from cache")
		return
	}
	structFieldMap = map[string]reflect.StructField{}
	structColumnMap = map[string]reflect.StructField{}
	for i := 0; i < tOf.NumField(); i++ {
		field := tOf.Field(i)
		structFieldMap[field.Name] = field
		str := field.Tag.Get("column")
		if str != "" && str != "-" {
			ss := strings.Split(str, ",")
			structColumnMap[ss[0]] = field
		} else {
			str = field.Tag.Get("json")
			if str != "" && str != "-" {
				ss := strings.Split(str, ",")
				structColumnMap[ss[0]] = field
			}
		}
	}
	structFieldMapCache[tOf] = structFieldMap
	structColumnMapCache[tOf] = structColumnMap
	return
}
func SetStructColumnValues(columnValueMap map[string]interface{}, strValue reflect.Value) {
	if len(columnValueMap) == 0 {
		return
	}
	tOf := strValue.Type()

	_, structColumnMap := getStructColumn(tOf)

	for columnName, columnValue := range columnValueMap {
		field, find := structColumnMap[columnName]
		if !find {
			field, find = structColumnMap[columnName]
		}
		if !find {
			continue
		}
		valueTypeOf := reflect.TypeOf(columnValue)
		columnValueType := ""
		fieldType := field.Type.String()
		if valueTypeOf != nil {
			columnValueType = valueTypeOf.String()
		}
		if columnValueType != fieldType {
			switch fieldType {
			case "string":
				columnValue = dialect.GetStringValue(columnValue)
				break
			case "int8", "int16", "int32", "int64", "int":
				str := dialect.GetStringValue(columnValue)
				var num int64
				if str != "" {
					num, _ = dialect.StringToInt64(str)
				}
				if fieldType == "int8" {
					columnValue = int8(num)
				} else if fieldType == "int16" {
					columnValue = int16(num)
				} else if fieldType == "int32" {
					columnValue = int32(num)
				} else if fieldType == "int64" {
					columnValue = num
				} else if fieldType == "int" {
					columnValue = int(num)
				}
				break
			case "uint8", "uint16", "uint32", "uint64", "uint":
				str := dialect.GetStringValue(columnValue)
				var num uint64
				if str != "" {
					num, _ = dialect.StringToUint64(str)
				}
				if fieldType == "uint8" {
					columnValue = uint8(num)
				} else if fieldType == "uint16" {
					columnValue = uint16(num)
				} else if fieldType == "uint32" {
					columnValue = uint32(num)
				} else if fieldType == "uint64" {
					columnValue = num
				} else if fieldType == "uint" {
					columnValue = uint(num)
				}
				break
			case "float32", "float64":
				str := dialect.GetStringValue(columnValue)
				var num float64
				if str != "" {
					num, _ = strconv.ParseFloat(str, 64)
				}
				if fieldType == "float32" {
					columnValue = float32(num)
				} else if fieldType == "float64" {
					columnValue = num
				}
				break
			case "time.Time":
				if columnValue == nil || columnValue == 0 {
					columnValue = time.Time{}
					break
				}
				valueOf := reflect.ValueOf(columnValue)
				if valueOf.IsNil() || valueOf.IsZero() {
					columnValue = time.Time{}
				}
				break
			}
		}

		valueOf := reflect.ValueOf(columnValue)
		strValue.FieldByName(field.Name).Set(valueOf)
	}
	return
}

func GetListStructType(list interface{}) reflect.Type {
	vOf := reflect.ValueOf(list)
	if vOf.Kind() == reflect.Ptr {
		return GetListStructType(vOf.Elem().Interface())
	}
	tOf := reflect.TypeOf(list).Elem()
	if tOf.Kind() == reflect.Ptr { //指针类型获取真正type需要调用Elem
		tOf = tOf.Elem()
	}
	return tOf
}

func DoQueryPageStructs(db *sql.DB, dia dialect.Dialect, sqlInfo string, args []interface{}, page *worker.Page, list interface{}) (err error) {
	if page.PageSize < 1 {
		page.PageSize = 1
	}
	if page.PageNo < 1 {
		page.PageNo = 1
	}
	pageSize := page.PageSize
	pageNo := page.PageNo

	countSql, err := dialect.FormatCountSql(sqlInfo)
	if err != nil {
		return
	}
	page.TotalCount, err = worker.DoQueryCount(db, countSql, args)
	if err != nil {
		return
	}
	page.TotalPage = (page.TotalCount + page.PageSize - 1) / page.PageSize
	// 如果查询的页码 大于 总页码 则不查询
	if pageNo > page.TotalPage {
		return
	}
	pageSql := dia.PackPageSql(sqlInfo, pageSize, pageNo)

	err = DoQueryStructs(db, pageSql, args, list)
	if err != nil {
		return
	}

	return
}
