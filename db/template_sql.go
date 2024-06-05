package db

import (
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"reflect"
	"strings"
)

func (this_ *TemplateOptions) GetInsertSql(table string, obj any, ignoreColumns ...string) (insertSql string, sqlArgs []any) {
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

func (this_ *TemplateOptions) GetBatchInsertSql(table string, many any, ignoreColumns ...string) (insertSql string, sqlArgsList [][]any) {

	columns, valuesList := this_.GetBatchColumnValues(many, ignoreColumns...)
	insertSql = "INSERT INTO " + this_.Dialect.TableNamePack(this_.DialectParam, table)
	insertSql += "(" + columns + ") VALUES ("
	valuesSql := strings.Repeat("?, ", len(valuesList[0]))
	valuesSql = strings.TrimSuffix(valuesSql, ", ")
	insertSql += valuesSql + ")"

	insertSql = this_.Dialect.ReplaceSqlVariable(insertSql, valuesList[0])
	sqlArgsList = valuesList
	return
}

func (this_ *TemplateOptions) GetUpdateSql(table string, obj any, whereParser *SqlParamParser, ignoreColumns ...string) (updateSql string, sqlArgs []any, err error) {
	columns, values := this_.GetColumnValues(obj, true, ignoreColumns...)
	updateSql = "UPDATE " + this_.Dialect.TableNamePack(this_.DialectParam, table)
	updateSql += " SET " + columns

	sqlArgs = values
	if whereParser != nil {
		var whereSql string
		var whereArgs []any
		whereSql, whereArgs, err = whereParser.Parse()
		if err != nil {
			util.Logger.Error("get update sql where sql parse error", zap.Error(err))
			return
		}
		updateSql += " WHERE " + whereSql
		sqlArgs = append(sqlArgs, whereArgs...)
	}
	updateSql = this_.Dialect.ReplaceSqlVariable(updateSql, sqlArgs)
	return
}

func (this_ *TemplateOptions) GetDeleteSql(table string, whereParser *SqlParamParser) (deleteSql string, sqlArgs []any, err error) {
	deleteSql = "DELETE FROM " + this_.Dialect.TableNamePack(this_.DialectParam, table)

	if whereParser != nil {
		var whereSql string
		var whereArgs []any
		whereSql, whereArgs, err = whereParser.Parse()
		if err != nil {
			util.Logger.Error("get delete sql where sql parse error", zap.Error(err))
			return
		}
		deleteSql += " WHERE " + whereSql
		sqlArgs = append(sqlArgs, whereArgs...)
	}
	deleteSql = this_.Dialect.ReplaceSqlVariable(deleteSql, sqlArgs)
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

func (this_ *TemplateOptions) GetColumnValues(obj any, isUpdate bool, ignoreColumns ...string) (columnsSql string, values []any) {
	if obj == nil {
		panic("GetColumnValues obj is null")
	}
	info := this_.GetStructInfo(reflect.TypeOf(obj))

	ignoreStr := strings.ToLower(strings.Join(ignoreColumns, ";") + ";")
	objV := reflect.ValueOf(obj)
	for objV.Kind() == reflect.Ptr {
		objV = objV.Elem()
	}
	columns := info.structColumns
	if info.isMap {
		columns = this_.GetMapColumns(obj)
	}
	for _, column := range columns {
		if strings.Contains(ignoreStr, strings.ToLower(column.ColumnName)+";") {
			continue
		}
		var vV reflect.Value
		if column.value != nil {
			vV = *column.value
		} else {
			vV = objV.Field(column.Index)
		}
		v := vV.Interface()
		isNull := this_.ValueIsNull(column, vV)
		if isNull && !isUpdate {
			continue
		}
		if len(columnsSql) > 0 {
			columnsSql += ", "
		}
		columnsSql += this_.Dialect.ColumnNamePack(this_.DialectParam, column.ColumnName)

		if this_.ValueIsNull(column, vV) {
			if isUpdate {
				columnsSql += " = NULL"
			}
		} else {
			if isUpdate {
				columnsSql += " = ?"
			}
			values = append(values, v)
		}
	}

	return
}

func (this_ *TemplateOptions) GetBatchColumnValues(many any, ignoreColumns ...string) (columnsSql string, valuesList [][]any) {
	if many == nil {
		panic("GetBatchColumnValues many is null")
	}
	manyV := reflect.ValueOf(many)
	for manyV.Kind() == reflect.Ptr {
		manyV = manyV.Elem()
	}
	if manyV.Kind() != reflect.Slice {
		panic("GetBatchColumnValues many not a array")
	}
	size := manyV.Len()
	if size == 0 {
		panic("GetBatchColumnValues many size is 0")
	}
	info := this_.GetStructInfo(reflect.TypeOf(manyV.Index(0).Interface()))

	ignoreStr := strings.ToLower(strings.Join(ignoreColumns, ";") + ";")
	for objIndex := 0; objIndex < size; objIndex++ {
		obj := manyV.Index(objIndex).Interface()
		objV := reflect.ValueOf(obj)
		for objV.Kind() == reflect.Ptr {
			objV = objV.Elem()
		}

		var values []any

		columns := info.structColumns
		if info.isMap {
			columns = this_.GetMapColumns(obj)
		}

		for _, column := range columns {
			if strings.Contains(ignoreStr, strings.ToLower(column.ColumnName)+";") {
				continue
			}
			if objIndex == 0 {
				if len(columnsSql) > 0 {
					columnsSql += ", "
				}
				columnsSql += this_.Dialect.ColumnNamePack(this_.DialectParam, column.ColumnName)
			}
			var vV reflect.Value
			if column.value != nil {
				vV = *column.value
			} else {
				vV = objV.Field(column.Index)
			}
			v := vV.Interface()
			if this_.ValueIsNull(column, vV) {
				values = append(values, nil)
			} else {
				values = append(values, v)
			}
		}
		valuesList = append(valuesList, values)
	}
	return
}

func (this_ *TemplateOptions) ValueIsNull(column *FieldColumn, vField reflect.Value) bool {

	v := vField.Interface()
	if v != nil {
		if !this_.StringEmptyNotUseNull && column.IsString && v == "" {
			return true
		} else if !this_.NumberZeroNotUseNull && column.IsNumber && vField.IsZero() {
			return true
		} else {
			return false
		}
	} else {
		return true
	}
}
