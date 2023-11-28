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

type ExecuteOptions struct {
	SelectDataMax int  `json:"selectDataMax"`
	OpenProfiling bool `json:"openProfiling"`
}
type executeTask struct {
	config       Config
	databaseType *DatabaseType
	dia          dialect.Dialect
	*Param
	ownerName string
	*ExecuteOptions
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
	defer func() {
		// 如果 是 mysql 关闭 profiling
		if this_.dia.DialectType() == dialect.TypeMysql && this_.OpenProfiling {
			_, _ = exec("SET profiling = 0")
		}
	}()
	// 如果 是 mysql 开启 profiling
	if this_.dia.DialectType() == dialect.TypeMysql && this_.OpenProfiling {
		_, _ = exec("SET profiling = 1")
	}
	sqlList := this_.dia.SqlSplit(sqlContent)
	var lastQueryID int
	for _, executeSql := range sqlList {
		lastQueryID, executeData, err = this_.execExecuteSQL(lastQueryID, executeSql, query, exec)
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

func (this_ *executeTask) execExecuteSQL(lastQueryID int, executeSql string,
	query func(query string, args ...any) (*sql.Rows, error),
	exec func(query string, args ...any) (sql.Result, error),
) (queryID int, executeData map[string]interface{}, err error) {

	queryID = lastQueryID
	executeData = map[string]interface{}{}
	var startTime = util.GetNow()
	executeData["sql"] = executeSql
	executeData["startTime"] = util.GetFormatByTime(startTime)

	defer func() {

		var endTime = time.Now()
		executeData["endTime"] = util.GetFormatByTime(endTime)
		executeData["isEnd"] = true
		executeData["useTime"] = util.GetMilliByTime(endTime) - util.GetMilliByTime(startTime)
		if err != nil {
			executeData["error"] = err.Error()
		}

		// 如果 是 mysql 关闭 profiling
		if this_.dia.DialectType() == dialect.TypeMysql && this_.OpenProfiling {
			queryID, executeData["profiling"], _ = queryProfiling(lastQueryID, query)
		}
	}()

	str := strings.ToLower(executeSql)
	if strings.HasPrefix(str, "select") ||
		strings.HasPrefix(str, "show") ||
		strings.HasPrefix(str, "desc") ||
		strings.HasPrefix(str, "explain") {
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
		var columnList []map[string]interface{}
		var dataList []map[string]interface{}
		var dataSize int
		dataSize, columnList, dataList, err = RowsToListMap(rows, this_.SelectDataMax)
		if err != nil {
			return
		}

		executeData["columnList"] = columnList
		executeData["dataSize"] = dataSize

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

func queryProfiling(lastQueryID int, query func(query string, args ...any) (*sql.Rows, error)) (queryID int, profiling map[string]interface{}, err error) {
	queryID = lastQueryID
	var dataList []map[string]interface{}
	// 查询
	var rows *sql.Rows
	rows, err = query("SHOW PROFILES")
	if err != nil {
		return
	}
	_, _, dataList, err = RowsToListMap(rows, 0)
	_ = rows.Close()
	if err != nil {
		return
	}
	if len(dataList) == 0 {
		return
	}
	//util.Logger.Debug("SHOW PROFILES", zap.Any("lastQueryID", lastQueryID), zap.Any("PROFILES", dataList))

	var data map[string]interface{}
	for _, one := range dataList {
		if one["Query_ID"] == nil {
			continue
		}
		id := util.StringToInt(util.GetStringValue(one["Query_ID"]))
		if lastQueryID < id {
			queryID = id
			data = one
			break
		}
	}

	var columnList []map[string]interface{}
	if data == nil || data["Query_ID"] == nil {
		return
	}

	rows, err = query("SHOW PROFILE ALL FOR QUERY " + util.GetStringValue(data["Query_ID"]))
	if err != nil {
		return
	}
	_, columnList, dataList, err = RowsToListMap(rows, 0)
	_ = rows.Close()
	if err != nil {
		return
	}
	data["columnList"] = columnList
	data["profileDataList"] = dataList

	profiling = data
	return
}

func RowsToListMap(rows *sql.Rows, selectDataMax int) (dataSize int, columnList []map[string]interface{}, dataList []map[string]interface{}, err error) {
	var columnTypes []*sql.ColumnType
	columnTypes, err = rows.ColumnTypes()
	if err != nil {
		return
	}

	if len(columnTypes) > 0 {
		for _, columnType := range columnTypes {
			column := map[string]interface{}{}
			column["name"] = columnType.Name()
			column["type"] = columnType.DatabaseTypeName()
			columnList = append(columnList, column)
		}
	}
	for rows.Next() {
		if selectDataMax > 0 && dataSize >= selectDataMax {
			dataSize++
			continue
		}
		dataSize++
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
					v = util.GetMilliByTime(tV)
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
						item[name] = util.GetMilliByTime(tV)
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
	return
}
