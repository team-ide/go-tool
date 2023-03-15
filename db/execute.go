package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"strings"
	"time"
)

type executeTask struct {
	config       Config
	databaseType *DatabaseType
	dia          dialect.Dialect
	*Param
	ownerName string
}

func (this_ *executeTask) run(sqlContent string) (executeList []map[string]interface{}, errStr string, err error) {
	var executeData map[string]interface{}
	var query func(query string, args ...any) (*sql.Rows, error)
	var exec func(query string, args ...any) (sql.Result, error)

	workDb, err := newWorkDb(this_.databaseType, this_.config, this_.ExecUsername, this_.ExecPassword, this_.ownerName)
	if err != nil {
		util.Logger.Error("ExecuteSQL new db pool error", zap.Error(err))
		return
	}
	defer func() {
		_ = workDb.Close()
	}()
	cxt := context.Background()
	conn, err := workDb.Conn(cxt)
	if err != nil {
		util.Logger.Error("ExecuteSQL Conn error", zap.Error(err))
		return
	}
	defer func() {
		_ = conn.Close()
	}()
	var hasError bool
	if this_.OpenTransaction {
		var tx *sql.Tx
		tx, err = conn.BeginTx(cxt, nil)
		if err != nil {
			util.Logger.Error("ExecuteSQL BeginTx error", zap.Error(err))
			return
		}
		defer func() {
			if hasError {
				err = tx.Rollback()
			} else {
				err = tx.Commit()
				if err != nil && strings.Contains(err.Error(), "Not in transaction") {
					err = nil
				}
			}
		}()
		query = tx.Query
		exec = tx.Exec
	} else {
		query = func(query string, args ...any) (*sql.Rows, error) {
			return conn.QueryContext(cxt, query, args...)
		}
		exec = func(query string, args ...any) (sql.Result, error) {
			return conn.ExecContext(cxt, query, args...)
		}
	}
	sqlList := this_.dia.SqlSplit(sqlContent)
	for _, executeSql := range sqlList {
		executeData, err = this_.execExecuteSQL(executeSql, query, exec)
		executeList = append(executeList, executeData)
		if err != nil {
			util.Logger.Error("ExecuteSQL execExecuteSQL error", zap.Any("executeSql", executeSql), zap.Error(err))
			errStr = err.Error()
			hasError = true
			if !this_.ErrorContinue {
				return
			}
			err = nil
		}
	}
	return
}

func (this_ *executeTask) execExecuteSQL(executeSql string,
	query func(query string, args ...any) (*sql.Rows, error),
	exec func(query string, args ...any) (sql.Result, error),
) (executeData map[string]interface{}, err error) {

	executeData = map[string]interface{}{}
	var startTime = util.GetNow()
	executeData["sql"] = executeSql
	executeData["startTime"] = util.GetFormatByTime(startTime)

	defer func() {
		var endTime = time.Now()
		executeData["endTime"] = util.GetFormatByTime(endTime)
		executeData["isEnd"] = true
		executeData["useTime"] = util.GetTimeByTime(endTime) - util.GetTimeByTime(startTime)
		if err != nil {
			executeData["error"] = err.Error()
			return
		}
	}()

	str := strings.ToLower(executeSql)
	if strings.HasPrefix(str, "select") ||
		strings.HasPrefix(str, "show") ||
		strings.HasPrefix(str, "desc") {
		executeData["isSelect"] = true
		// 查询
		var rows *sql.Rows
		rows, err = query(executeSql)
		if err != nil {
			return
		}
		defer func() {
			_ = rows.Close()
		}()
		var columnTypes []*sql.ColumnType
		columnTypes, err = rows.ColumnTypes()
		if err != nil {
			return
		}

		var columnList []map[string]interface{}
		if len(columnTypes) > 0 {
			for _, columnType := range columnTypes {
				column := map[string]interface{}{}
				column["name"] = columnType.Name()
				column["type"] = columnType.DatabaseTypeName()
				columnList = append(columnList, column)
			}
		}
		executeData["columnList"] = columnList
		var dataList []map[string]interface{}
		for rows.Next() {
			cache := worker.GetSqlValueCache(columnTypes) //临时存储每行数据
			err = rows.Scan(cache...)
			if err != nil {
				return
			}
			item := make(map[string]interface{})
			for index, data := range cache {
				name := columnTypes[index].Name()
				var v interface{}
				switch tV := data.(type) {
				case time.Time:
					if tV.IsZero() {
						v = nil
					} else {
						v = util.GetTimeByTime(tV)
					}
				default:
					v = worker.GetSqlValue(columnTypes[index], data)
				}
				if v == nil {
					item[name] = v
				} else {
					switch tV := v.(type) {
					case time.Time:
						if tV.IsZero() {
							item[name] = nil
						} else {
							item[name] = util.GetTimeByTime(tV)
						}
					case float64:
						if tV >= float64(9007199254740991) || tV <= float64(-9007199254740991) {
							item[name] = fmt.Sprintf("%f", tV)
						} else {
							item[name] = tV
						}
					case int64:
						if tV >= int64(9007199254740991) || tV <= int64(-9007199254740991) {
							item[name] = fmt.Sprintf("%d", tV)
						} else {
							item[name] = tV
						}
					case int, int8, int16, int32:
						item[name] = tV
					default:
						item[name] = fmt.Sprint(tV)
					}
				}
			}
			dataList = append(dataList, item)
		}
		executeData["dataList"] = dataList
	} else if strings.HasPrefix(str, "insert") {
		executeData["isInsert"] = true
		var result sql.Result
		result, err = exec(executeSql)
		if err != nil {
			return
		}
		executeData["rowsAffected"], _ = result.RowsAffected()
	} else if strings.HasPrefix(str, "update") {
		executeData["isUpdate"] = true
		var result sql.Result
		result, err = exec(executeSql)
		if err != nil {
			return
		}
		executeData["rowsAffected"], _ = result.RowsAffected()
	} else if strings.HasPrefix(str, "delete") {
		executeData["isDelete"] = true
		var result sql.Result
		result, err = exec(executeSql)
		if err != nil {
			return
		}
		executeData["rowsAffected"], _ = result.RowsAffected()
	} else {
		executeData["isExec"] = true
		var result sql.Result
		result, err = exec(executeSql)
		if err != nil {
			return
		}
		executeData["rowsAffected"], _ = result.RowsAffected()
	}

	return
}
