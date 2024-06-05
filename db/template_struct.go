package db

import (
	"database/sql"
	"github.com/team-ide/go-dialect/worker"
	"github.com/team-ide/go-tool/util"
	"reflect"
	"strings"
	"time"
)

func (this_ *TemplateOptions) getColumnValue(columnType *sql.ColumnType, value any) (res any, err error) {
	if this_.GetColumnValue != nil {
		res, err = this_.GetColumnValue(columnType, value)
		return
	}
	res = worker.GetSqlValue(columnType, value)
	return
}

func (this_ *TemplateOptions) setFieldValue(columnType *sql.ColumnType, field reflect.StructField, fieldValue reflect.Value, value any) (err error) {
	if this_.SetFieldValue != nil {
		err = this_.SetFieldValue(columnType, field, fieldValue, value)
		return
	}

	columnValue, err := this_.getColumnValue(columnType, value)
	if err != nil {
		return
	}
	if columnValue == nil {
		return
	}
	fieldType := field.Type
	valueOf := reflect.ValueOf(columnValue)
	valueType := valueOf.Type()
	if valueType == fieldType {
		fieldValue.Set(valueOf)
		return
	}
	var isPtr = fieldType.Kind() == reflect.Ptr
	if isPtr {
		fieldType = fieldType.Elem()
	}
	fieldTypeKind := fieldType.Kind()
	if fieldTypeKind == reflect.String {
		var str = util.GetStringValue(columnValue)
		if isPtr {
			fieldValue.Set(reflect.ValueOf(&str))
		} else {
			fieldValue.SetString(str)
		}
	} else if fieldTypeKind >= reflect.Int && fieldTypeKind <= reflect.Int64 {
		var num int64
		if num, err = util.ValueToInt64(columnValue); err != nil {
			return
		}
		if isPtr {
			switch fieldTypeKind {
			case reflect.Int:
				v := int(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Int8:
				v := int8(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Int16:
				v := int16(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Int32:
				v := int32(num)
				fieldValue.Set(reflect.ValueOf(&v))
			default:
				fieldValue.Set(reflect.ValueOf(&num))
			}
		} else {
			fieldValue.SetInt(num)
		}
	} else if fieldTypeKind >= reflect.Uint && fieldTypeKind <= reflect.Uint64 {
		var num uint64
		if num, err = util.ValueToUint64(columnValue); err != nil {
			return
		}
		if isPtr {
			switch fieldTypeKind {
			case reflect.Uint:
				v := uint(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Uint8:
				v := uint8(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Uint16:
				v := uint16(num)
				fieldValue.Set(reflect.ValueOf(&v))
			case reflect.Uint32:
				v := uint32(num)
				fieldValue.Set(reflect.ValueOf(&v))
			default:
				fieldValue.Set(reflect.ValueOf(&num))
			}
		} else {
			fieldValue.SetUint(num)
		}
	} else if fieldTypeKind == reflect.Float32 || fieldTypeKind == reflect.Float64 {
		var num float64
		if num, err = util.ValueToFloat64(columnValue); err != nil {
			return
		}
		if isPtr {
			switch fieldTypeKind {
			case reflect.Float32:
				v := float32(num)
				fieldValue.Set(reflect.ValueOf(&v))
			default:
				fieldValue.Set(reflect.ValueOf(&num))
			}
		} else {
			fieldValue.SetFloat(num)
		}
	} else if fieldTypeKind == reflect.Bool {
		var v bool
		v, ok := columnValue.(bool)
		if !ok {
			str := util.GetStringValue(columnValue)
			if str == "1" || str == "true" {
				v = true
			}
		}
		if isPtr {
			fieldValue.Set(reflect.ValueOf(&v))
		} else {
			fieldValue.SetBool(v)
		}
	} else if fieldType.String() == "time.Time" {
		var v time.Time
		v, ok := columnValue.(time.Time)
		if !ok {
			var num int64
			if num, err = util.ValueToInt64(columnValue); err != nil {
				return
			}
			v = time.UnixMilli(num)
		}
		if isPtr {
			fieldValue.Set(reflect.ValueOf(&v))
		} else {
			fieldValue.Set(reflect.ValueOf(v))
		}
	} else {
		if isPtr {
			fieldValue.Set(reflect.ValueOf(&columnValue))
		} else {
			fieldValue.Set(reflect.ValueOf(columnValue))
		}
	}
	return
}

func (this_ *TemplateOptions) fullMapValues(res reflect.Value, values []any, columns []*sql.ColumnType) (err error) {
	objV := res
	for objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}
	data := map[string]interface{}{}
	var columnValue any
	for i, column := range columns {
		columnValue, err = this_.getColumnValue(column, values[i])
		if err != nil {
			return
		}
		data[column.Name()] = columnValue
	}
	objV.Set(reflect.ValueOf(data))
	return
}

func (this_ *TemplateOptions) fullStructValues(res reflect.Value, values []any, columns []*sql.ColumnType, fieldColumns []*FieldColumn) (err error) {

	objV := res
	for objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}
	//data := map[string]interface{}{}
	for i, column := range columns {
		fieldColumn := fieldColumns[i]
		if fieldColumn == nil {
			continue
		}
		fieldV := objV.Field(fieldColumn.Index)
		err = this_.setFieldValue(column, fieldColumn.Field, fieldV, values[i])
		if err != nil {
			return
		}
	}
	//err = mapstructure.Decode(data, res.Interface())
	//if err != nil {
	//	panic(err)
	//}
	return
}
func (this_ *Template[T]) getValues(columns []*sql.ColumnType) (res reflect.Value, values []interface{}, fieldColumns []*FieldColumn) {
	if this_.objValueType.Kind() == reflect.Map {
		res = reflect.New(this_.objValueType)
		// 处理 map
		for range columns {
			values = append(values, new(interface{}))
		}
		return
	}
	res = reflect.New(this_.objValueType)
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
		values = append(values, new(interface{}))
		if fieldColumn != nil {
			fieldColumns = append(fieldColumns, fieldColumn)
		} else {
			fieldColumns = append(fieldColumns, nil)
		}
	}
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
		if str == "" && !this_.NotUseJsonTagName {
			str = field.Tag.Get("json")
		}
		if str == "" && !this_.NotUseFieldName {
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
