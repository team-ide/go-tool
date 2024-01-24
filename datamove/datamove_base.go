package datamove

import (
	"github.com/team-ide/go-dialect/dialect"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/elasticsearch"
	"github.com/team-ide/go-tool/kafka"
	"github.com/team-ide/go-tool/redis"
	"github.com/team-ide/go-tool/util"
	"os"
)

type Options struct {
	Key  string            `json:"key,omitempty"`  // 任务的 key
	Dir  string            `json:"dir,omitempty"`  // 任务过程中 生成文件的目录
	From *DataSourceConfig `json:"from,omitempty"` // 源 数据配置
	To   *DataSourceConfig `json:"to,omitempty"`   // 目标 数据配置

	ErrorContinue bool  `json:"errorContinue,omitempty"`
	BatchNumber   int64 `json:"batchNumber,omitempty"`
}

func (this_ *Options) getFilePath(dirName string, fileName string, suffix string) (path string) {
	dir := this_.Dir
	if dirName != "" {
		dir = this_.Dir + dirName + "/"
		exists, _ := util.PathExists(dir)
		if !exists {
			_ = os.MkdirAll(dir, os.ModePerm)
		}
	}
	path = dir + fileName + "." + suffix

	return
}

func (this_ *DataSourceConfig) GetDialectParam() *dialect.ParamModel {
	if this_.ParamModel == nil {
		this_.ParamModel = &dialect.ParamModel{}
	}
	return this_.ParamModel
}

func (this_ *DataSourceConfig) GetFileName() string {
	if this_.FileName == "" {
		return "导出"
	}
	return this_.FileName
}

type DataSourceConfig struct {
	*dialect.ParamModel
	Type             string `json:"type,omitempty"`
	SqlFileMergeType string `json:"sqlFileMergeType,omitempty"` // SQL 的文件合并类型 如：one：一个文件， owner：每个库一个文件，table：每个表一个文件

	DataSourceSqlParam

	DataSourceTxtParam

	DataSourceScriptParam

	DataSourceRedisParam

	DataSourceExcelParam

	DataSourceKafkaParam

	DataSourceEsParam

	BySql bool `json:"bySql,omitempty"` // 根据 SQL 语句导出

	DataSourceDbParam

	DataSourceDataParam

	TxtFileType string `json:"txtFileType,omitempty"` //

	// 数据库 配置
	DbConfig *db.Config `json:"-"`

	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`

	EsConfig *elasticsearch.Config `json:"-"`

	RedisConfig *redis.Config `json:"-"`

	KafkaConfig *kafka.Config `json:"-"`

	FillColumn bool `json:"fillColumn,omitempty"` // 自动填充列

	ColumnList []*Column `json:"columnList,omitempty"`

	AllOwner       bool       `json:"allOwner,omitempty"`
	Owners         []*DbOwner `json:"owners,omitempty"`
	SkipOwnerNames []string   `json:"skipOwnerNames,omitempty"`

	FilePath    string `json:"filePath,omitempty"`
	ShouldOwner bool   `json:"shouldOwner,omitempty"` // 需要 建库
	ShouldTable bool   `json:"shouldTable,omitempty"` // 需要 建表

	FileNameSplice string `json:"fileNameSplice,omitempty"` // 文件名拼接字符 如：/ :库作为目录 表作为名称 默认
	FileName       string `json:"fileName,omitempty"`
	RowNumber      int64  `json:"rowNumber,omitempty"`

	dbService db.IService
	dia_      dialect.Dialect
}

func (this_ *DataSourceConfig) GetDialect() dialect.Dialect {
	if this_.dia_ == nil {
		if this_.DialectType == "" {
			return this_.dbService.GetDialect()
		} else {
			this_.dia_, _ = dialect.NewDialect(this_.DialectType)
		}
	}
	return this_.dia_
}

func (this_ *DataSourceConfig) GetTxtFileType() string {
	if this_.TxtFileType == "" {
		return "txt"
	}
	return this_.TxtFileType
}
func (this_ *DataSourceConfig) IsData() bool {
	return this_.Type == "data"
}
func (this_ *DataSourceConfig) IsDb() bool {
	return this_.Type == "database"
}
func (this_ *DataSourceConfig) IsEs() bool {
	return this_.Type == "elasticsearch"
}
func (this_ *DataSourceConfig) IsTxt() bool {
	return this_.Type == "txt"
}
func (this_ *DataSourceConfig) IsExcel() bool {
	return this_.Type == "excel"
}
func (this_ *DataSourceConfig) IsSql() bool {
	return this_.Type == "sql"
}
func (this_ *DataSourceConfig) IsKafka() bool {
	return this_.Type == "kafka"
}
func (this_ *DataSourceConfig) IsRedis() bool {
	return this_.Type == "redis"
}
func (this_ *DataSourceConfig) IsScript() bool {
	return this_.Type == "script"
}

type DbOwner struct {
	From           *dialect.OwnerModel `json:"from,omitempty"`
	To             *dialect.OwnerModel `json:"to,omitempty"`
	SkipTableNames []string            `json:"skipTableNames,omitempty"`
	AllTable       bool                `json:"allTable,omitempty"`
	Tables         []*DbTable          `json:"tables,omitempty"`
	fromService    db.IService
	toService      db.IService
	appended       bool
}

type DbTable struct {
	From            *dialect.TableModel `json:"from,omitempty"`
	To              *dialect.TableModel `json:"to,omitempty"`
	Columns         []*DbColumn         `json:"columns,omitempty"`
	SkipColumnNames []string            `json:"skipColumnNames,omitempty"`
	AllColumn       bool                `json:"allColumn,omitempty"`
	appended        bool
}

func (this_ *DbTable) GetToDialectTable() *dialect.TableModel {
	table := &dialect.TableModel{}
	table.TableName = this_.To.TableName

	for _, c := range this_.Columns {
		column := &dialect.ColumnModel{}
		column.ColumnName = c.To.ColumnName

		column.ColumnDataType = c.From.ColumnDataType
		column.ColumnComment = c.From.ColumnComment
		column.ColumnDefault = c.From.ColumnDefault
		column.ColumnExtra = c.From.ColumnExtra
		column.ColumnEnums = c.From.ColumnEnums
		column.ColumnLength = c.From.ColumnLength
		column.ColumnNotNull = c.From.ColumnNotNull
		column.ColumnCharacterSetName = c.From.ColumnCharacterSetName
		column.ColumnPrecision = c.From.ColumnPrecision
		column.ColumnScale = c.From.ColumnScale
		table.ColumnList = append(table.ColumnList, column)
	}
	table.PrimaryKeys = this_.From.PrimaryKeys
	table.IndexList = this_.From.IndexList
	for _, one := range table.IndexList {
		one.IndexName = ""
	}
	return table
}

type DbColumn struct {
	From  *dialect.ColumnModel `json:"from,omitempty"`
	To    *dialect.ColumnModel `json:"to,omitempty"`
	Value string               `json:"value,omitempty"`
}
