package db

import (
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
)

type IService interface {
	Stop()
	GetDialectName() string
	Exec(sql string, args []interface{}) (rowsAffected int64, err error)
	Execs(sqlList []string, argsList [][]interface{}) (rowsAffected int64, err error)
	Count(sql string, args []interface{}) (count int64, err error)
	QueryOne(sql string, args []interface{}, one interface{}) (find bool, err error)
	Query(sql string, args []interface{}, list interface{}) (err error)
	QueryMap(sql string, args []interface{}) (list []map[string]interface{}, err error)
	QueryPage(sql string, args []interface{}, list interface{}, page *worker.Page) (err error)
	QueryMapPage(sql string, args []interface{}, page *worker.Page) (list []map[string]interface{}, err error)
	Info() (res interface{}, err error)
	OwnerCreateSql(param *Param, owner *dialect.OwnerModel) (sqlList []string, err error)
	OwnersSelect(param *Param) (owners []*dialect.OwnerModel, err error)
	TablesSelect(param *Param, ownerName string) (tables []*dialect.TableModel, err error)
	TableDetail(param *Param, ownerName string, tableName string) (tableDetail *dialect.TableModel, err error)
	OwnerCreate(param *Param, owner *dialect.OwnerModel) (created bool, err error)
	OwnerDelete(param *Param, ownerName string) (deleted bool, err error)
	OwnerDataTrim(param *Param, ownerName string) (err error)
	DDL(param *Param, ownerName string, tableName string) (sqlList []string, err error)
	Model(param *Param, ownerName string, tableName string) (content string, err error)
	TableCreate(param *Param, ownerName string, table *dialect.TableModel) (err error)
	TableDelete(param *Param, ownerName string, tableName string) (err error)
	TableDataTrim(param *Param, ownerName string, tableName string) (err error)
	TableCreateSql(param *Param, ownerName string, table *dialect.TableModel) (sqlList []string, err error)
	TableUpdateSql(param *Param, ownerName string, tableName string, updateTableParam *UpdateTableParam) (sqlList []string, err error)
	TableUpdate(param *Param, ownerName string, tableName string, updateTableParam *UpdateTableParam) (err error)
	DataListSql(param *Param, ownerName string, tableName string, columnList []*dialect.ColumnModel,
		insertDataList []map[string]interface{},
		updateDataList []map[string]interface{}, updateWhereDataList []map[string]interface{},
		deleteDataList []map[string]interface{},
	) (sqlList []string, err error)
	DataListExec(param *Param, ownerName string, tableName string, columnList []*dialect.ColumnModel,
		insertDataList []map[string]interface{},
		updateDataList []map[string]interface{}, updateWhereDataList []map[string]interface{},
		deleteDataList []map[string]interface{},
	) (err error)

	TableData(param *Param, ownerName string, tableName string, columnList []*dialect.ColumnModel, whereList []*dialect.Where, orderList []*dialect.Order, pageSize int, pageNo int) (dataListResult DataListResult, err error)
	GetTargetDialect(param *Param) (dia dialect.Dialect)
	ExecuteSQL(param *Param, ownerName string, sqlContent string) (executeList []map[string]interface{}, errStr string, err error)
	StartExport(param *Param, exportParam *worker.TaskExportParam) (task *worker.Task, err error)
	StartImport(param *Param, importParam *worker.TaskImportParam) (task *worker.Task, err error)
	StartSync(param *Param, syncParam *worker.TaskSyncParam) (task *worker.Task, err error)
}
