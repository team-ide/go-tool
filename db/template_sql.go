package db

import (
	"reflect"
	"strings"
)

type InsertBuilder struct {
}

func (this_ *Template[T]) GetInsertSql(table string, obj T, ignoreColumns ...string) (insertSql string, sqlArgs []any) {
	columns, values := this_.GetColumnValues(obj, false, ignoreColumns...)
	insertSql = "INSERT INTO " + this_.Dialect.TableNamePack(this_.DialectParam, table)
	insertSql += "(" + columns + ") VALUES ("
	valuesSql := strings.Repeat("?, ", len(values))
	valuesSql = strings.TrimSuffix(valuesSql, ", ")
	insertSql += valuesSql + ")"

	insertSql = this_.Dialect.ReplaceSqlVariable(insertSql, sqlArgs)
	sqlArgs = values
	return
}

func (this_ *Template[T]) GetBatchInsertSql(table string, objs []T, ignoreColumns ...string) (insertSql string, sqlArgsList [][]any) {
	columns, valuesList := this_.GetBatchColumnValues(objs, ignoreColumns...)
	insertSql = "INSERT INTO " + this_.Dialect.TableNamePack(this_.DialectParam, table)
	insertSql += "(" + columns + ") VALUES ("
	valuesSql := strings.Repeat("?, ", len(valuesList[0]))
	valuesSql = strings.TrimSuffix(valuesSql, ", ")
	insertSql += valuesSql + ")"

	insertSql = this_.Dialect.ReplaceSqlVariable(insertSql, valuesList[0])
	sqlArgsList = valuesList
	return
}

func (this_ *Template[T]) GetUpdateSql(table string, obj T, ignoreColumns ...string) (updateSql string, sqlArgs []any) {
	columns, values := this_.GetColumnValues(obj, true, ignoreColumns...)
	updateSql = "UPDATE " + this_.Dialect.TableNamePack(this_.DialectParam, table)
	updateSql += " SET " + columns

	updateSql = this_.Dialect.ReplaceSqlVariable(updateSql, sqlArgs)
	sqlArgs = values
	return
}

func (this_ *Template[T]) GetColumnsSql() (columnsSql string) {
	info := this_.GetStructInfo(this_.objValueType)

	for _, column := range info.structColumns {
		if len(columnsSql) > 0 {
			columnsSql += ", "
		}
		columnsSql += this_.Dialect.ColumnNamePack(this_.DialectParam, column.ColumnName)
	}
	return
}

func (this_ *Template[T]) GetColumnValues(obj T, isUpdate bool, ignoreColumns ...string) (columnsSql string, values []any) {
	info := this_.GetStructInfo(reflect.TypeOf(obj))

	ignoreStr := strings.ToLower(strings.Join(ignoreColumns, ";") + ";")
	objV := reflect.ValueOf(obj)
	for objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}

	for _, column := range info.structColumns {
		if strings.Contains(ignoreStr, strings.ToLower(column.ColumnName)+";") {
			continue
		}
		if len(columnsSql) > 0 {
			columnsSql += ", "
		}
		columnsSql += this_.Dialect.ColumnNamePack(this_.DialectParam, column.ColumnName)
		if isUpdate {
			columnsSql += " = ?"
		}
		vField := objV.Field(column.Index)
		v := vField.Interface()
		if this_.ValueIsNull(column, vField) {
			values = append(values, nil)
		} else {
			values = append(values, v)
		}

	}
	return
}

func (this_ *Template[T]) GetBatchColumnValues(objs []T, ignoreColumns ...string) (columnsSql string, valuesList [][]any) {
	if objs == nil || len(objs) == 0 {
		panic("GetColumnValues objs is null")
	}
	info := this_.GetStructInfo(reflect.TypeOf(objs[0]))

	ignoreStr := strings.ToLower(strings.Join(ignoreColumns, ";") + ";")
	for objIndex, obj := range objs {
		objV := reflect.ValueOf(obj)
		for objV.Kind() == reflect.Ptr {
			objV = objV.Elem()
		}
		var values []any

		for _, column := range info.structColumns {
			if strings.Contains(ignoreStr, strings.ToLower(column.ColumnName)+";") {
				continue
			}
			if objIndex == 0 {
				if len(columnsSql) > 0 {
					columnsSql += ", "
				}
				columnsSql += this_.Dialect.ColumnNamePack(this_.DialectParam, column.ColumnName)
			}
			vField := objV.Field(column.Index)
			v := vField.Interface()
			if this_.ValueIsNull(column, vField) {
				values = append(values, nil)
			} else {
				values = append(values, v)
			}
		}
		valuesList = append(valuesList, values)
	}
	return
}

func (this_ *Template[T]) ValueIsNull(column *FieldColumn, vField reflect.Value) bool {

	v := vField.Interface()
	if v != nil {
		if this_.StringEmptyUseNull && column.IsString && v == "" {
			return true
		} else if this_.NumberZeroUseNull && column.IsNumber && vField.IsZero() {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}
