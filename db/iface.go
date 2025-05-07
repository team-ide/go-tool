package db

import (
	"database/sql"
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-dialect/worker"
	"github.com/team-ide/go-tool/util"
)

type IService interface {
	// Close 关闭 db 客户端
	Close()
	// GetConfig 获取 数据库 配置
	GetConfig() Config
	// GetDialect 获取 方言
	GetDialect() dialect.Dialect
	// GetDb 获取 sql.DB
	GetDb() *sql.DB

	// Info 查询数据库信息
	Info() (res interface{}, err error)

	// Exec 执行 SQL
	Exec(sql string, args []interface{}) (rowsAffected int64, err error)
	// Execs 批量执行 SQL
	Execs(sqlList []string, argsList [][]interface{}) (rowsAffected int64, err error)
	// Count 统计查询 SQL
	Count(sql string, args []interface{}) (count int64, err error)
	// QueryOne 查询 单个 结构体 SQL  需要 传入 结构体 接收
	QueryOne(sql string, args []interface{}, one interface{}) (find bool, err error)
	// Query 查询 列表 结构体 SQL  需要 传入 结构体 接收
	Query(sql string, args []interface{}, list interface{}) (err error)
	// QueryMap 查询 列表 map SQL  返回 list < map >
	QueryMap(sql string, args []interface{}) (list []map[string]interface{}, err error)
	// QueryPage 分页查询 列表 结构体 SQL  需要 传入 结构体 接收
	QueryPage(sql string, args []interface{}, list interface{}, page *worker.Page) (err error)
	// QueryMapPage 分页查询 列表 map 列表 SQL  返回 list < map >
	QueryMapPage(sql string, args []interface{}, page *worker.Page) (list []map[string]interface{}, err error)

	// OwnerCreateSql 创建 库|模式|用户等 SQL 取决于哪种数据库
	OwnerCreateSql(param *Param, owner *dialect.OwnerModel) (sqlList []string, err error)
	// OwnersSelect 查询 库|模式|用户等 取决于哪种数据库
	OwnersSelect(param *Param) (owners []*dialect.OwnerModel, err error)

	// TablesSelect 查询 库|模式|用户 下所有表
	TablesSelect(param *Param, ownerName string) (tables []*dialect.TableModel, err error)
	// TableDetail 查询 某个表明细
	TableDetail(param *Param, ownerName string, tableName string) (tableDetail *dialect.TableModel, err error)

	// OwnerCreate 创建 库|模式|用户等 取决于哪种数据库
	OwnerCreate(param *Param, owner *dialect.OwnerModel) (created bool, err error)
	// OwnerDelete 删除 库|模式|用户等 取决于哪种数据库
	OwnerDelete(param *Param, ownerName string) (deleted bool, err error)
	// OwnerDataTrim 清空 库|模式|用户等 数据 取决于哪种数据库
	OwnerDataTrim(param *Param, ownerName string) (err error)
	// DDL 查看 库 或 表 SQL语句
	DDL(param *Param, ownerName string, tableName string) (sqlList []string, err error)
	// Model 将 表 转换某些语言的模型
	Model(param *Param, ownerName string, tableName string) (content string, err error)

	// TableCreate 创建表
	TableCreate(param *Param, ownerName string, table *dialect.TableModel) (err error)
	// TableDelete 删除表
	TableDelete(param *Param, ownerName string, tableName string) (err error)
	// TableDataTrim 表数据清理
	TableDataTrim(param *Param, ownerName string, tableName string) (err error)
	// TableCreateSql 获取建表 SQL
	TableCreateSql(param *Param, ownerName string, table *dialect.TableModel) (sqlList []string, err error)
	// TableUpdateSql 获取更新表的 SQL
	TableUpdateSql(param *Param, ownerName string, tableName string, updateTableParam *UpdateTableParam) (sqlList []string, err error)
	// TableUpdate 更新表
	TableUpdate(param *Param, ownerName string, tableName string, updateTableParam *UpdateTableParam) (err error)

	// DataListSql 将 需要 新增、修改、删除的 数据 转为 SQL
	DataListSql(param *Param, ownerName string, tableName string, columnList []*dialect.ColumnModel,
		insertDataList []map[string]interface{},
		updateDataList []map[string]interface{}, updateWhereDataList []map[string]interface{},
		deleteDataList []map[string]interface{},
	) (sqlList []string, err error)
	// DataListExec 执行 需要 新增、修改、删除的 数据
	DataListExec(param *Param, ownerName string, tableName string, columnList []*dialect.ColumnModel,
		insertDataList []map[string]interface{},
		updateDataList []map[string]interface{}, updateWhereDataList []map[string]interface{},
		deleteDataList []map[string]interface{},
	) (info *ExecuteInfo, err error)

	// TableData 根据 一些 条件 查询 表数据
	TableData(param *Param, ownerName string, tableName string, columnList []*dialect.ColumnModel, whereList []*dialect.Where, orderList []*dialect.Order, pageSize int, pageNo int) (dataListResult DataListResult, err error)

	// GetTargetDialect 获取 目标数据库方言
	GetTargetDialect(param *Param) (dia dialect.Dialect)

	// ExecuteSQL 执行某段 SQL
	ExecuteSQL(param *Param, ownerName string, sqlContent string, options *ExecuteOptions) (executeList []map[string]interface{}, errStr string, err error)

	// StartExport 开始 导出
	StartExport(param *Param, exportParam *worker.TaskExportParam) (task *worker.Task, err error)
	// StartImport 开始 导入
	StartImport(param *Param, importParam *worker.TaskImportParam) (task *worker.Task, err error)
	// StartSync 开始 同步
	StartSync(param *Param, syncParam *worker.TaskSyncParam) (task *worker.Task, err error)

	// NewTestExecutor 新建测试任务
	NewTestExecutor(options *TestTaskOptions) (testExecutor *TestExecutor, err error)
}

func ToOwnerModel(data interface{}) (res *dialect.OwnerModel) {
	res = &dialect.OwnerModel{}
	var bs []byte
	switch tV := data.(type) {
	case dialect.OwnerModel:
		res = &tV
		return
	case *dialect.OwnerModel:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}

func ToTableModel(data interface{}) (res *dialect.TableModel) {
	res = &dialect.TableModel{}
	var bs []byte
	switch tV := data.(type) {
	case dialect.TableModel:
		res = &tV
		return
	case *dialect.TableModel:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}

func ToColumnModel(data interface{}) (res *dialect.ColumnModel) {
	res = &dialect.ColumnModel{}
	var bs []byte
	switch tV := data.(type) {
	case dialect.ColumnModel:
		res = &tV
		return
	case *dialect.ColumnModel:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}

func ToIndexModel(data interface{}) (res *dialect.IndexModel) {
	res = &dialect.IndexModel{}
	var bs []byte
	switch tV := data.(type) {
	case dialect.IndexModel:
		res = &tV
		return
	case *dialect.IndexModel:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}

func ToPrimaryKeyModel(data interface{}) (res *dialect.PrimaryKeyModel) {
	res = &dialect.PrimaryKeyModel{}
	var bs []byte
	switch tV := data.(type) {
	case dialect.PrimaryKeyModel:
		res = &tV
		return
	case *dialect.PrimaryKeyModel:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}

func ToPage(data interface{}) (res *worker.Page) {
	res = &worker.Page{}
	var bs []byte
	switch tV := data.(type) {
	case worker.Page:
		res = &tV
		return
	case *worker.Page:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}

func ToTaskExportParam(data interface{}) (res *worker.TaskExportParam) {
	res = &worker.TaskExportParam{}
	var bs []byte
	switch tV := data.(type) {
	case worker.TaskExportParam:
		res = &tV
		return
	case *worker.TaskExportParam:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}
func ToTaskImportParam(data interface{}) (res *worker.TaskImportParam) {
	res = &worker.TaskImportParam{}
	var bs []byte
	switch tV := data.(type) {
	case worker.TaskImportParam:
		res = &tV
		return
	case *worker.TaskImportParam:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}

func ToTaskSyncParam(data interface{}) (res *worker.TaskSyncParam) {
	res = &worker.TaskSyncParam{}
	var bs []byte
	switch tV := data.(type) {
	case worker.TaskSyncParam:
		res = &tV
		return
	case *worker.TaskSyncParam:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}

func ToConfig(data interface{}) (res *Config) {
	res = &Config{}
	var bs []byte
	switch tV := data.(type) {
	case Config:
		res = &tV
		return
	case *Config:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}

func ToParam(data interface{}) (res *Param) {
	res = &Param{}
	var bs []byte
	switch tV := data.(type) {
	case Param:
		res = &tV
		return
	case *Param:
		res = tV
		return
	case string:
		bs = []byte(tV)
	case *string:
		if tV != nil {
			bs = []byte(*tV)
		}
	case []byte:
		bs = tV
	default:
		_ = util.ObjToObjByJson(data, res)
		return
	}
	_ = util.JSONDecodeUseNumber(bs, res)
	return
}
