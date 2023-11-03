package db

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"os"
	"strings"
	"time"
)

type Config struct {
	Type            string `json:"type,omitempty"`
	Host            string `json:"host,omitempty"`
	Port            int    `json:"port,omitempty"`
	Database        string `json:"database,omitempty"`
	DbName          string `json:"dbName,omitempty"`
	Username        string `json:"username,omitempty"`
	Password        string `json:"password,omitempty"`
	OdbcDsn         string `json:"odbcDsn,omitempty"`
	OdbcDialectName string `json:"odbcDialectName,omitempty"`

	Schema       string      `json:"schema,omitempty"`
	Sid          string      `json:"sid,omitempty"`
	MaxIdleConn  int         `json:"maxIdleConn,omitempty"`
	MaxOpenConn  int         `json:"maxOpenConn,omitempty"`
	DatabasePath string      `json:"databasePath,omitempty"`
	SSHClient    *ssh.Client `json:"-"`
	Dsn          string      `json:"dsn,omitempty"`
}

// New 创建 db客户端
func New(config *Config) (IService, error) {
	return createService(config)
}

func createService(config *Config) (*Service, error) {
	service := &Service{
		config: config,
	}
	err := service.init()
	if err != nil {
		return nil, err
	}
	return service, nil
}

// Service 注册处理器在线信息等
type Service struct {
	config       *Config
	databaseType *DatabaseType
	db           *sql.DB
	dialect.Dialect
}

func (this_ *Service) GetConfig() Config {
	return *this_.config
}

func (this_ *Service) GetDialect() dialect.Dialect {
	return this_.Dialect
}

func (this_ *Service) init() (err error) {
	this_.databaseType = GetDatabaseType(this_.config.Type)
	if this_.databaseType == nil {
		err = errors.New("数据库类型[" + this_.config.Type + "]暂不支持")
		return
	}

	if this_.databaseType.DialectName == "odbc" {
		if this_.config.OdbcDialectName != "" {
			this_.Dialect, err = dialect.NewDialect(this_.config.OdbcDialectName)
			if err != nil {
				return
			}
		}
	}
	if this_.Dialect == nil {
		this_.Dialect = this_.databaseType.dia
	}
	this_.db, err = this_.databaseType.NewDb(this_.config)
	if err != nil {
		return
	}

	if this_.config.MaxIdleConn > 0 {
		this_.db.SetMaxIdleConns(this_.config.MaxIdleConn)
	}
	if this_.config.MaxOpenConn > 0 {
		this_.db.SetMaxOpenConns(this_.config.MaxOpenConn)
	}

	err = this_.db.Ping()
	if err != nil {
		return
	}
	return
}

func (this_ *Service) GetDb() *sql.DB {
	if this_ != nil && this_.db != nil {
		return this_.db
	}
	return nil
}

func (this_ *Service) Close() {
	if this_ != nil && this_.db != nil {
		_ = this_.db.Close()
	}
	return
}

func (this_ *Service) Exec(sql string, args []interface{}) (rowsAffected int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
		if err != nil {
			err = errors.New("Exec error sql:" + sql + ",error:" + err.Error())
		}
	}()

	rowsAffected, err = this_.Execs([]string{sql}, [][]interface{}{args})

	if err != nil {
		return
	}
	return
}

func (this_ *Service) Execs(sqlList []string, argsList [][]interface{}) (rowsAffected int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
		if err != nil {
			util.Logger.Error("Execs Error", zap.Any("sqlList", sqlList), zap.Any("argsList", argsList), zap.Error(err))
			err = errors.New("Execs error sql:" + strings.Join(sqlList, ";") + ",error:" + err.Error())
		}
	}()
	res, errSql, errArgs, err := worker.DoExecs(this_.db, sqlList, argsList)
	if err != nil {
		util.Logger.Error("Execs Error", zap.Any("sql", errSql), zap.Any("args", errArgs), zap.Error(err))
		return
	}
	for _, one := range res {
		rowsAffected_, _ := one.RowsAffected()
		rowsAffected += rowsAffected_
	}
	return
}

func (this_ *Service) Count(sql string, args []interface{}) (count int64, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
		if err != nil {
			util.Logger.Error("Count Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
			err = errors.New("Count error sql:" + sql + ",error:" + err.Error())
		}
	}()
	count_, err := worker.DoQueryCount(this_.db, sql, args)
	if err != nil {
		util.Logger.Error("Count Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
		return
	}
	count = int64(count_)
	return
}

func (this_ *Service) QueryOne(sql string, args []interface{}, one interface{}) (find bool, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
		if err != nil {
			util.Logger.Error("QueryOne Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
			err = errors.New("QueryOne error sql:" + sql + ",error:" + err.Error())
		}
	}()
	find, err = worker.DoQueryStruct(this_.db, sql, args, one)

	if err != nil {
		util.Logger.Error("QueryOne Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
		return
	}

	return
}

func (this_ *Service) Query(sql string, args []interface{}, list interface{}) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
		if err != nil {
			util.Logger.Error("Query Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
			err = errors.New("Query error sql:" + sql + ",error:" + err.Error())
		}
	}()
	err = worker.DoQueryStructs(this_.db, sql, args, list)

	if err != nil {
		util.Logger.Error("Query Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
		return
	}

	return
}

func (this_ *Service) QueryMap(sql string, args []interface{}) (list []map[string]interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
		if err != nil {
			util.Logger.Error("QueryMap Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
			err = errors.New("QueryMap error sql:" + sql + ",error:" + err.Error())
		}
	}()
	list, err = worker.DoQuery(this_.db, sql, args)

	if err != nil {
		util.Logger.Error("QueryMap Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
		return
	}

	return
}

func (this_ *Service) QueryPage(sql string, args []interface{}, list interface{}, page *worker.Page) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
		if err != nil {
			util.Logger.Error("QueryPage Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
			err = errors.New("QueryPage error sql:" + sql + ",error:" + err.Error())
		}
	}()
	err = worker.DoQueryPageStructs(this_.db, this_.Dialect, sql, args, page, list)

	if err != nil {
		util.Logger.Error("QueryPage Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
		return
	}

	return
}

func (this_ *Service) QueryMapPage(sql string, args []interface{}, page *worker.Page) (list []map[string]interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
		}
		if err != nil {
			util.Logger.Error("QueryMapPage Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
			err = errors.New("QueryMapPage error sql:" + sql + ",error:" + err.Error())
		}
	}()
	list, err = worker.DoQueryPage(this_.db, this_.Dialect, sql, args, page)

	if err != nil {
		util.Logger.Error("QueryMapPage Error", zap.Any("sql", sql), zap.Any("args", args), zap.Error(err))
		return
	}

	return
}

func (this_ *Service) Info() (res interface{}, err error) {
	return
}

func (this_ *Service) OwnersSelect(param *Param) (owners []*dialect.OwnerModel, err error) {
	owners, err = worker.OwnersSelect(this_.db, this_.Dialect, param.ParamModel)
	return
}

func (this_ *Service) TablesSelect(param *Param, ownerName string) (tables []*dialect.TableModel, err error) {
	tables, err = worker.TablesSelect(this_.db, this_.Dialect, param.ParamModel, ownerName)
	return
}

func (this_ *Service) TableDetail(param *Param, ownerName string, tableName string) (tableDetail *dialect.TableModel, err error) {
	tableDetail, err = worker.TableDetail(this_.db, this_.Dialect, param.ParamModel, ownerName, tableName, true)
	return
}

func (this_ *Service) OwnerCreateSql(param *Param, owner *dialect.OwnerModel) (sqlList []string, err error) {
	sqlList, err = this_.Dialect.OwnerCreateSql(param.ParamModel, owner)
	return
}

func (this_ *Service) OwnerCreate(param *Param, owner *dialect.OwnerModel) (created bool, err error) {
	created, err = worker.OwnerCreate(this_.db, this_.Dialect, param.ParamModel, owner)
	return
}

func (this_ *Service) OwnerDelete(param *Param, ownerName string) (deleted bool, err error) {
	deleted, err = worker.OwnerDelete(this_.db, this_.Dialect, param.ParamModel, ownerName)
	return
}

func (this_ *Service) OwnerDataTrim(param *Param, ownerName string) (err error) {
	tables, err := worker.TablesSelect(this_.db, this_.Dialect, param.ParamModel, ownerName)
	if err != nil {
		return
	}
	for _, table := range tables {
		sqlInfo := "DELETE FROM " + this_.OwnerTablePack(param.ParamModel, ownerName, table.TableName)
		_, _, _, err = worker.DoOwnerExecs(this_.Dialect, this_.db, ownerName, []string{sqlInfo}, [][]interface{}{nil})
		if err != nil {
			return
		}
	}
	return
}

func (this_ *Service) DDL(param *Param, ownerName string, tableName string) (sqlList []string, err error) {
	var sqlList_ []string
	if param.AppendOwnerCreateSql {
		var owner *dialect.OwnerModel
		owner, err = worker.OwnerSelect(this_.db, this_.Dialect, param.ParamModel, ownerName)
		if err != nil {
			return
		}
		if owner != nil {
			sqlList_, err = this_.GetTargetDialect(param).OwnerCreateSql(param.ParamModel, owner)
			if err != nil {
				return
			}
			sqlList = append(sqlList, sqlList_...)
		}
	}
	var tables []*dialect.TableModel
	if tableName != "" {
		var table *dialect.TableModel
		table, err = worker.TableDetail(this_.db, this_.Dialect, param.ParamModel, ownerName, tableName, false)
		if err != nil {
			return
		}
		if table != nil {
			tables = append(tables, table)
		}
	} else {
		tables, err = worker.TablesDetail(this_.db, this_.Dialect, param.ParamModel, ownerName, false)
		if err != nil {
			return
		}
	}
	for _, table := range tables {
		var appendOwnerName string
		if param.AppendOwnerName {
			appendOwnerName = ownerName
		}
		sqlList_, err = this_.GetTargetDialect(param).TableCreateSql(param.ParamModel, appendOwnerName, table)
		if err != nil {
			return
		}
		sqlList = append(sqlList, sqlList_...)
	}

	return
}

func (this_ *Service) Model(param *Param, ownerName string, tableName string) (content string, err error) {

	var tables []*dialect.TableModel
	if tableName != "" {
		var table *dialect.TableModel
		table, err = worker.TableDetail(this_.db, this_.Dialect, param.ParamModel, ownerName, tableName, false)
		if err != nil {
			return
		}
		if table != nil {
			tables = append(tables, table)
		}
	} else {
		tables, err = worker.TablesDetail(this_.db, this_.Dialect, param.ParamModel, ownerName, false)
		if err != nil {
			return
		}
	}
	gen := &modelGen{
		modelType: param.ModelType,
		tables:    tables,
	}
	content, err = gen.gen()

	return
}
func (this_ *Service) TableCreate(param *Param, ownerName string, table *dialect.TableModel) (err error) {
	workDb, err := newWorkDb(this_.databaseType, *this_.config, "", "", ownerName)
	if err != nil {
		util.Logger.Error("TableCreate new db pool error", zap.Error(err))
		return
	}
	defer func() { _ = workDb.Close() }()

	err = worker.TableCreate(workDb, this_.Dialect, param.ParamModel, ownerName, table)
	return
}

func (this_ *Service) TableDelete(param *Param, ownerName string, tableName string) (err error) {
	workDb, err := newWorkDb(this_.databaseType, *this_.config, "", "", ownerName)
	if err != nil {
		util.Logger.Error("TableDelete new db pool error", zap.Error(err))
		return
	}
	defer func() { _ = workDb.Close() }()

	err = worker.TableDelete(workDb, this_.Dialect, param.ParamModel, ownerName, tableName)
	return
}

func (this_ *Service) TableDataTrim(param *Param, ownerName string, tableName string) (err error) {
	sqlInfo := "DELETE FROM " + this_.OwnerTablePack(param.ParamModel, ownerName, tableName)
	_, _, _, err = worker.DoOwnerExecs(this_.Dialect, this_.db, ownerName, []string{sqlInfo}, [][]interface{}{nil})
	return
}

func (this_ *Service) TableCreateSql(param *Param, ownerName string, table *dialect.TableModel) (sqlList []string, err error) {
	var appendOwnerName string
	if param.AppendOwnerName {
		appendOwnerName = ownerName
	}
	sqlList, err = this_.GetTargetDialect(param).TableCreateSql(param.ParamModel, appendOwnerName, table)

	return
}

type UpdateTableParam struct {
	TableComment    string               `json:"tableComment"`
	OldTableComment string               `json:"oldTableComment"`
	ColumnList      []*UpdateTableColumn `json:"columnList"`
	IndexList       []*UpdateTableIndex  `json:"indexList"`
}

type UpdateTableColumn struct {
	*dialect.ColumnModel
	OldColumn *dialect.ColumnModel `json:"oldColumn"`
	Deleted   bool                 `json:"deleted"`
}
type UpdateTableIndex struct {
	*dialect.IndexModel
	OldIndex *dialect.IndexModel `json:"oldIndex"`
	Deleted  bool                `json:"deleted"`
}

func (this_ *Service) TableUpdateSql(param *Param, ownerName string, tableName string, updateTableParam *UpdateTableParam) (sqlList []string, err error) {
	var last *UpdateTableColumn
	for _, one := range updateTableParam.ColumnList {
		if last != nil {
			one.ColumnAfterColumn = last.ColumnName
		}
		last = one
	}

	var newPrimaryKeys []string
	var oldPrimaryKeys []string
	var sqlList_ []string
	for _, one := range updateTableParam.ColumnList {
		if one.PrimaryKey {
			newPrimaryKeys = append(newPrimaryKeys, one.ColumnName)
		}
		if one.OldColumn != nil {
			if one.OldColumn.PrimaryKey {
				oldPrimaryKeys = append(oldPrimaryKeys, one.OldColumn.ColumnName)
			}
		}
		if one.Deleted {
			if one.OldColumn != nil {
				sqlList_, err = this_.GetTargetDialect(param).ColumnDeleteSql(param.ParamModel, ownerName, tableName, one.OldColumn.ColumnName)
				if err != nil {
					return
				}
				sqlList = append(sqlList, sqlList_...)
			}
		} else if one.OldColumn == nil {
			sqlList_, err = this_.GetTargetDialect(param).ColumnAddSql(param.ParamModel, ownerName, tableName, one.ColumnModel)
			if err != nil {
				return
			}
			sqlList = append(sqlList, sqlList_...)
		} else {
			sqlList_, err = this_.GetTargetDialect(param).ColumnUpdateSql(param.ParamModel, ownerName, tableName, one.OldColumn, one.ColumnModel)
			if err != nil {
				return
			}
			sqlList = append(sqlList, sqlList_...)
		}
	}

	var primaryKeyChange bool
	for _, oldPrimaryKey := range oldPrimaryKeys {
		if dialect.StringsIndex(newPrimaryKeys, oldPrimaryKey) < 0 {
			primaryKeyChange = true
		}
	}
	if !primaryKeyChange {
		for _, newPrimaryKey := range newPrimaryKeys {
			if dialect.StringsIndex(oldPrimaryKeys, newPrimaryKey) < 0 {
				primaryKeyChange = true
			}
		}
	}
	if primaryKeyChange {
		if len(oldPrimaryKeys) > 0 {
			sqlList_, err = this_.GetTargetDialect(param).PrimaryKeyDeleteSql(param.ParamModel, ownerName, tableName)
			if err != nil {
				return
			}
			sqlList = append(sqlList, sqlList_...)
		}
		if len(newPrimaryKeys) > 0 {
			sqlList_, err = this_.GetTargetDialect(param).PrimaryKeyAddSql(param.ParamModel, ownerName, tableName, newPrimaryKeys)
			if err != nil {
				return
			}
			sqlList = append(sqlList, sqlList_...)
		}
	}

	for _, one := range updateTableParam.IndexList {

		if one.Deleted {
			if one.OldIndex != nil {
				sqlList_, err = this_.GetTargetDialect(param).IndexDeleteSql(param.ParamModel, ownerName, tableName, one.OldIndex.IndexName)
				if err != nil {
					return
				}
				sqlList = append(sqlList, sqlList_...)
			}
		} else if one.OldIndex == nil {
			sqlList_, err = this_.GetTargetDialect(param).IndexAddSql(param.ParamModel, ownerName, tableName, one.IndexModel)
			if err != nil {
				return
			}
			sqlList = append(sqlList, sqlList_...)
		} else {
			if one.IndexName != one.OldIndex.IndexName ||
				one.IndexType != one.OldIndex.IndexType ||
				one.IndexComment != one.OldIndex.IndexComment ||
				strings.Join(one.ColumnNames, ",") != strings.Join(one.OldIndex.ColumnNames, ",") {
				sqlList_, err = this_.GetTargetDialect(param).IndexDeleteSql(param.ParamModel, ownerName, tableName, one.OldIndex.IndexName)
				if err != nil {
					return
				}
				sqlList = append(sqlList, sqlList_...)
				sqlList_, err = this_.GetTargetDialect(param).IndexAddSql(param.ParamModel, ownerName, tableName, one.IndexModel)
				if err != nil {
					return
				}
				sqlList = append(sqlList, sqlList_...)
			}
		}
	}
	return
}

func (this_ *Service) TableUpdate(param *Param, ownerName string, tableName string, updateTableParam *UpdateTableParam) (err error) {
	sqlList, err := this_.TableUpdateSql(param, ownerName, tableName, updateTableParam)
	if err != nil {
		return
	}

	workDb, err := newWorkDb(this_.databaseType, *this_.config, "", "", ownerName)
	if err != nil {
		util.Logger.Error("TableUpdate new db pool error", zap.Error(err))
		return
	}
	defer func() { _ = workDb.Close() }()

	_, errSql, _, err := worker.DoOwnerExecs(this_.Dialect, workDb, ownerName, sqlList, nil)
	if err != nil {
		err = errors.New("sql:" + errSql + ",error:" + err.Error())
		util.Logger.Error("TableUpdate error:", zap.Error(err))
		return
	}
	return
}

func (this_ *Service) DataListSql(param *Param, ownerName string, tableName string, columnList []*dialect.ColumnModel,
	insertDataList []map[string]interface{},
	updateDataList []map[string]interface{}, updateWhereDataList []map[string]interface{},
	deleteDataList []map[string]interface{},
) (sqlList []string, err error) {
	var appendOwnerName string
	if param.AppendOwnerName {
		appendOwnerName = ownerName
	}

	dia := this_.GetTargetDialect(param)
	var sqlList_ []string
	//var valuesList [][]interface{}
	if len(insertDataList) > 0 {
		param.ParamModel.AppendSqlValue = new(bool)
		*param.ParamModel.AppendSqlValue = true
		sqlList_, _, _, _, err = dia.DataListInsertSql(param.ParamModel, appendOwnerName, tableName, columnList, insertDataList)
		if err != nil {
			return
		}
		sqlList = append(sqlList, sqlList_...)
	}
	if len(updateDataList) > 0 {
		param.ParamModel.AppendSqlValue = new(bool)
		*param.ParamModel.AppendSqlValue = true
		sqlList_, _, err = dia.DataListUpdateSql(param.ParamModel, appendOwnerName, tableName, columnList, updateDataList, updateWhereDataList)
		if err != nil {
			return
		}
		sqlList = append(sqlList, sqlList_...)
	}
	if len(deleteDataList) > 0 {
		param.ParamModel.AppendSqlValue = new(bool)
		*param.ParamModel.AppendSqlValue = true
		sqlList_, _, err = dia.DataListDeleteSql(param.ParamModel, appendOwnerName, tableName, columnList, deleteDataList)
		if err != nil {
			return
		}
		sqlList = append(sqlList, sqlList_...)
	}
	return
}

func (this_ *Service) DataListExec(param *Param, ownerName string, tableName string, columnList []*dialect.ColumnModel,
	insertDataList []map[string]interface{},
	updateDataList []map[string]interface{}, updateWhereDataList []map[string]interface{},
	deleteDataList []map[string]interface{},
) (err error) {
	var appendOwnerName = ownerName

	dia := this_.GetTargetDialect(param)
	var sqlList []string
	var valuesList [][]interface{}

	var sqlList_ []string
	var valuesList_ [][]interface{}
	if len(insertDataList) > 0 {
		sqlList_, valuesList_, _, _, err = dia.DataListInsertSql(param.ParamModel, appendOwnerName, tableName, columnList, insertDataList)
		if err != nil {
			return
		}
		sqlList = append(sqlList, sqlList_...)
		valuesList = append(valuesList, valuesList_...)
	}
	if len(updateDataList) > 0 {
		sqlList_, valuesList_, err = dia.DataListUpdateSql(param.ParamModel, appendOwnerName, tableName, columnList, updateDataList, updateWhereDataList)
		if err != nil {
			return
		}
		sqlList = append(sqlList, sqlList_...)
		valuesList = append(valuesList, valuesList_...)
	}
	if len(deleteDataList) > 0 {
		sqlList_, valuesList_, err = dia.DataListDeleteSql(param.ParamModel, appendOwnerName, tableName, columnList, deleteDataList)
		if err != nil {
			return
		}
		sqlList = append(sqlList, sqlList_...)
		valuesList = append(valuesList, valuesList_...)
	}
	_, errSql, errArgs, err := worker.DoOwnerExecs(this_.Dialect, this_.db, appendOwnerName, sqlList, valuesList)
	if err != nil {
		util.Logger.Error("DataListExec error", zap.Any("errSql", errSql), zap.Any("errArgs", errArgs), zap.Error(err))
		return
	}

	return
}

type DataListResult struct {
	Sql      string                   `json:"sql"`
	Total    int                      `json:"total"`
	Args     []interface{}            `json:"args"`
	DataList []map[string]interface{} `json:"dataList"`
}

func (this_ *Service) TableData(param *Param, ownerName string, tableName string, columnList []*dialect.ColumnModel, whereList []*dialect.Where, orderList []*dialect.Order, pageSize int, pageNo int) (dataListResult DataListResult, err error) {

	param.AppendSqlValue = nil
	selectSql, values, err := this_.Dialect.DataListSelectSql(param.ParamModel, ownerName, tableName, columnList, whereList, orderList)
	if err != nil {
		return
	}
	param.AppendSqlValue = new(bool)
	*param.AppendSqlValue = true
	selectTextSql, _, err := this_.Dialect.DataListSelectSql(param.ParamModel, ownerName, tableName, columnList, whereList, orderList)
	if err != nil {
		return
	}

	page := worker.NewPage()
	page.PageSize = pageSize
	page.PageNo = pageNo
	listMap, err := this_.QueryMapPage(selectSql, values, page)
	if err != nil {
		return
	}
	var names []string
	for _, column := range columnList {
		names = append(names, strings.ToLower(column.ColumnName))
	}
	for _, item := range listMap {
		for name, v := range item {
			if v == nil {
				continue
			}
			if util.StringIndexOf(names, strings.ToLower(name)) < 0 {
				delete(item, name)
				continue
			}
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
	dataListResult.Sql = this_.PackPageSql(selectTextSql, page.PageSize, page.PageNo)
	dataListResult.Args = values
	dataListResult.Total = page.TotalCount
	dataListResult.DataList = listMap
	return
}

func (this_ *Service) GetTargetDialect(param *Param) (dia dialect.Dialect) {
	if param != nil && param.TargetDatabaseType != "" {
		t := GetDatabaseType(param.TargetDatabaseType)
		if t != nil {
			return t.dia
		}
	}
	return this_.Dialect
}

func (this_ *Service) ExecuteSQL(param *Param, ownerName string, sqlContent string, options *ExecuteOptions) (executeList []map[string]interface{}, errStr string, err error) {

	if options == nil {
		options = &ExecuteOptions{}
	}
	task := &executeTask{
		ExecuteOptions: options,
		config:         *this_.config,
		databaseType:   this_.databaseType,
		dia:            this_.GetTargetDialect(param),
		ownerName:      ownerName,
		Param:          param,
	}
	executeList, errStr, err = task.run(sqlContent)
	return
}

var (
	FileUploadDir string
)

func (this_ *Service) StartExport(param *Param, exportParam *worker.TaskExportParam) (task *worker.Task, err error) {
	downloadPath := "export/" + util.GetUUID()
	var exportDir string
	exportDir, err = util.GetTempDir()
	if err != nil {
		return
	}
	exportDir += downloadPath
	exportParam.Dir = exportDir
	exportParam.DataSourceType = worker.GetDataSource(param.ExportType)
	exportParam.OnProgress = func(progress *worker.TaskProgress) {
		util.Logger.Info("export task on progress", zap.Any("progress", progress))
		progress.OnError = func(err error) {
			util.Logger.Error("export task progress error", zap.Any("progress", progress), zap.Error(err))
		}
	}

	exportParam.FormatIndexName = func(ownerName string, tableName string, index *dialect.IndexModel) (indexNameFormat string) {
		if index.IndexName != "" && !param.FormatIndexName {
			indexNameFormat = index.IndexName
			return
		}
		if ownerName != "" {
			indexNameFormat += ownerName + "_"
			//indexNameFormat += sortName(ownerName, 4) + "_"
		}
		if tableName != "" {
			indexNameFormat += tableName + "_"
			//indexNameFormat += sortName(tableName, 4) + "_"
		}
		if index.IndexType != "" && !strings.EqualFold(index.IndexType, "index") {
			indexNameFormat += index.IndexType + "_"
			//indexNameFormat += sortName(index.IndexType, 4) + "_"
		}
		indexNameFormat += strings.Join(index.ColumnNames, "_")
		if this_.Dialect.DialectType() == dialect.TypeOracle {
			indexNameFormat = sortName(indexNameFormat, 30)
		}
		return
	}

	targetDialect := this_.GetTargetDialect(param)
	task_ := worker.NewTaskExport(this_.db, this_.Dialect, targetDialect, exportParam)
	task_.Param = param.ParamModel
	task_.Extend = map[string]interface{}{
		"downloadPath": "",
	}
	go func() {
		defer func() {
			if err != nil {
				_ = os.RemoveAll(exportDir)
				return
			}
			if exportParam.IsDataListExport {
				fs, _ := os.ReadDir(exportDir)
				for _, f := range fs {
					if f.IsDir() {
						continue
					}
					if strings.HasPrefix(f.Name(), "数据列表导出") {
						task_.Extend["dirPath"] = exportDir
						task_.Extend["downloadPath"] = exportDir + "/" + f.Name()
					}
				}
			} else {
				err = util.Zip(exportDir, exportDir+".zip")
				if err != nil {
					util.Logger.Error("export file zip error", zap.Any("exportDir", exportDir), zap.Error(err))
					return
				}
				task_.Extend["dirPath"] = exportDir
				task_.Extend["zipPath"] = exportDir + ".zip"
				task_.Extend["downloadPath"] = downloadPath + ".zip"
			}
		}()
		err = task_.Start()
		if err != nil {
			return
		}
	}()
	time.Sleep(time.Millisecond * 100)
	if err != nil {
		worker.ClearTask(task_.TaskId)
		return
	}
	task = task_.Task
	return
}

func (this_ *Service) StartImport(param *Param, importParam *worker.TaskImportParam) (task *worker.Task, err error) {
	for _, owner := range importParam.Owners {
		if owner.Path != "" {
			owner.Path = FileUploadDir + owner.Path
		}
		for _, table := range owner.Tables {
			if table.Path != "" {
				table.Path = FileUploadDir + table.Path
			}
		}
	}

	importParam.DataSourceType = worker.GetDataSource(param.ImportType)
	importParam.OnProgress = func(progress *worker.TaskProgress) {
		util.Logger.Info("import task on progress", zap.Any("progress", progress))
		progress.OnError = func(err error) {
			util.Logger.Error("import task progress error", zap.Any("progress", progress), zap.Error(err))
		}
	}
	var workDbs []*sql.DB
	databaseType := this_.databaseType
	config := *this_.config

	task_ := worker.NewTaskImport(this_.db, this_.Dialect,
		func(owner *worker.TaskImportOwner) (workDb *sql.DB, err error) {
			ownerName := owner.Name
			username := owner.Username
			password := owner.Password
			workDb, err = newWorkDb(databaseType, config, username, password, ownerName)
			if err != nil {
				util.Logger.Error("import new db pool error", zap.Error(err))
				return
			}
			workDbs = append(workDbs, workDb)
			return
		},
		importParam,
	)
	task_.Param = param.ParamModel

	go func() {

		defer func() {
			for _, workDb := range workDbs {
				_ = workDb.Close()
			}
		}()

		err = task_.Start()
		if err != nil {
			return
		}
	}()
	time.Sleep(time.Millisecond * 100)
	if err != nil {
		worker.ClearTask(task_.TaskId)
		return
	}
	task = task_.Task
	return
}

func (this_ *Service) StartSync(param *Param, syncParam *worker.TaskSyncParam) (task *worker.Task, err error) {

	targetDatabaseWorker, err := createService(param.TargetDatabaseConfig)
	if err != nil {
		return
	}
	syncParam.OnProgress = func(progress *worker.TaskProgress) {
		util.Logger.Info("sync task on progress", zap.Any("progress", progress))
		progress.OnError = func(err error) {
			util.Logger.Error("sync task progress error", zap.Any("progress", progress), zap.Error(err))
		}
	}
	syncParam.FormatIndexName = func(ownerName string, tableName string, index *dialect.IndexModel) (indexNameFormat string) {
		if index.IndexName != "" && !param.FormatIndexName {
			indexNameFormat = index.IndexName
			return
		}
		if ownerName != "" {
			indexNameFormat += ownerName + "_"
			//indexNameFormat += sortName(ownerName, 4) + "_"
		}
		if tableName != "" {
			indexNameFormat += tableName + "_"
			//indexNameFormat += sortName(tableName, 4) + "_"
		}
		if index.IndexType != "" && !strings.EqualFold(index.IndexType, "index") {
			indexNameFormat += index.IndexType + "_"
			//indexNameFormat += sortName(index.IndexType, 4) + "_"
		}
		indexNameFormat += strings.Join(index.ColumnNames, "_")
		if targetDatabaseWorker.Dialect.DialectType() == dialect.TypeOracle {
			indexNameFormat = sortName(indexNameFormat, 30)
		}
		return
	}
	var workDbs []*sql.DB

	task_ := worker.NewTaskSync(this_.db, this_.Dialect, targetDatabaseWorker.db, targetDatabaseWorker.Dialect,
		func(owner *worker.TaskSyncOwner) (workDb *sql.DB, err error) {
			ownerName := owner.TargetName
			if ownerName == "" {
				ownerName = owner.SourceName
			}
			username := owner.Username
			password := owner.Password
			workDb, err = newWorkDb(targetDatabaseWorker.databaseType, *param.TargetDatabaseConfig, username, password, ownerName)
			if err != nil {
				util.Logger.Error("sync new db pool error", zap.Error(err))
				return
			}
			workDbs = append(workDbs, workDb)
			return
		},
		syncParam,
	)
	task_.Param = param.ParamModel

	go func() {

		defer func() {
			for _, workDb := range workDbs {
				_ = workDb.Close()
			}
			_ = targetDatabaseWorker.db.Close()
		}()

		err = task_.Start()
		if err != nil {
			return
		}
	}()
	time.Sleep(time.Millisecond * 100)
	if err != nil {
		worker.ClearTask(task_.TaskId)
		return
	}
	task = task_.Task
	return
}

type Param struct {
	*dialect.ParamModel
	TargetDatabaseType   string  `json:"targetDatabaseType"`
	AppendOwnerCreateSql bool    `json:"appendOwnerCreateSql"`
	AppendOwnerName      bool    `json:"appendOwnerName"`
	OpenTransaction      bool    `json:"openTransaction"`
	ErrorContinue        bool    `json:"errorContinue"`
	ExecUsername         string  `json:"execUsername"`
	ExecPassword         string  `json:"execPassword"`
	ExportType           string  `json:"exportType"`
	ImportType           string  `json:"importType"`
	TargetDatabaseConfig *Config `json:"targetDatabaseConfig"`
	ModelType            string  `json:"modelType"`

	FormatIndexName bool `json:"formatIndexName"`
}

func sortName(name string, size int) (res string) {
	name = strings.TrimSpace(name)
	if len(name) <= size {
		res = name
		return
	}
	if strings.Contains(name, "_") {
		ss := strings.Split(name, "_")
		var names []string

		for _, s := range ss {
			if strings.TrimSpace(s) == "" {
				continue
			}
			names = append(names, strings.TrimSpace(s))
		}
		rSize := size / len(names)

		for i, s := range ss {
			if len(res) >= size {
				break
			}
			if i < len(ss)-1 {
				if rSize >= len(s)-1 {
					res += s + "_"
				} else {
					res += s[0:rSize-1] + "_"
				}
			} else {
				if rSize >= len(s) {
					res += s
				} else {
					res += s[0:rSize]
				}
			}
		}
	} else {
		res += name[0:size]
	}
	return
}

func newWorkDb(databaseType *DatabaseType, config Config, username string, password string, ownerName string) (workDb *sql.DB, err error) {
	config.MaxIdleConn = 3
	config.MaxOpenConn = 3
	if username != "" {
		config.Username = username
	}
	if password != "" {
		config.Password = password
	}
	switch databaseType.dia.DialectType() {
	case dialect.TypeMysql:
		config.Database = ownerName
		break
	case dialect.TypeGBase:

		if ownerName != "" {
			var keyL = len("db=")
			index := strings.Index(strings.ToLower(config.OdbcDsn), "db=")

			if index < 0 {
				index = strings.Index(strings.ToLower(config.OdbcDsn), "database=")
				keyL = len("database=")
			}
			if index >= 0 {
				beforeStr := config.OdbcDsn[0 : index+keyL]
				afterStr := ""
				str := config.OdbcDsn[index+keyL:]
				index = strings.Index(str, ";")
				if index >= 0 {
					afterStr = str[index:]
				}
				config.OdbcDsn = beforeStr + ownerName + afterStr
			}
		}

		break
	default:
		config.Schema = ownerName
		break
	}
	workDb, err = databaseType.NewDb(&config)
	return
}
