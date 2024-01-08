package datamove

import (
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/elasticsearch"
)

type Options struct {
	Key    string            `json:"key"`    // 任务的 key
	Dir    string            `json:"dir"`    // 任务过程中 生成文件的目录
	Source *DataSourceConfig `json:"source"` // 源 数据配置
	Target *DataSourceConfig `json:"target"` // 目标 数据配置

	AllOwner       bool       `json:"allOwner"`
	Owners         []*DbOwner `json:"owners"`
	SkipOwnerNames []string   `json:"skipOwnerNames"`

	OwnerName  string    `json:"ownerName"`
	TableName  string    `json:"tableName"`
	BySql      bool      `json:"bySql"` // 根据 SQL 语句导出
	SelectSql  string    `json:"selectSql"`
	ColumnList []*Column `json:"columnList"`

	IndexName string `json:"indexName"`

	DataList []map[string]interface{} `json:"dataList"`

	FilePath string `json:"filePath"`

	ErrorContinue bool  `json:"errorContinue"`
	BatchNumber   int64 `json:"batchNumber"`

	*dialect.ParamModel
}

type DataSourceConfig struct {
	Type string `json:"type"`

	// 数据库 配置
	DbConfig *db.Config `json:"-"`

	EsConfig *elasticsearch.Config `json:"-"`
}

type DbOwner struct {
	SourceName     string     `json:"sourceName"`
	TargetName     string     `json:"targetName"`
	SkipTableNames []string   `json:"skipTableNames"`
	AllTable       bool       `json:"allTable"`
	Tables         []*DbTable `json:"tables"`
	Username       string     `json:"username"`
	Password       string     `json:"password"`
	sourceService  db.IService
	targetService  db.IService
}

type DbTable struct {
	SourceName      string      `json:"sourceName"`
	TargetName      string      `json:"targetName"`
	Columns         []*DbColumn `json:"columns"`
	SkipColumnNames []string    `json:"skipColumnNames"`
	AllColumn       bool        `json:"allColumn"`
}

type DbColumn struct {
	*Column
	SourceName string `json:"sourceName"`
	TargetName string `json:"targetName"`
	Value      string `json:"value"`
}

func (this_ DataSourceConfig) IsDb() bool {
	return this_.Type == "database"
}

func (this_ DataSourceConfig) IsEs() bool {
	return this_.Type == "elasticsearch"
}

func (this_ DataSourceConfig) IsTxt() bool {
	return this_.Type == "txt"
}

func (this_ DataSourceConfig) IsExcel() bool {
	return this_.Type == "excel"
}

func (this_ DataSourceConfig) IsSql() bool {
	return this_.Type == "sql"
}

func (this_ Options) GetDialectParam() *dialect.ParamModel {
	if this_.ParamModel == nil {
		this_.ParamModel = &dialect.ParamModel{}
	}
	return this_.ParamModel
}
