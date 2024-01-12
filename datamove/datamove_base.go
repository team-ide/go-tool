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
	Key  string            `json:"key"`  // 任务的 key
	Dir  string            `json:"dir"`  // 任务过程中 生成文件的目录
	From *DataSourceConfig `json:"from"` // 源 数据配置
	To   *DataSourceConfig `json:"to"`   // 目标 数据配置

	ErrorContinue bool  `json:"errorContinue"`
	BatchNumber   int64 `json:"batchNumber"`
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
	Type             string `json:"type"`
	SqlFileMergeType string `json:"sqlFileMergeType"` // SQL 的文件合并类型 如：one：一个文件， owner：每个库一个文件，table：每个表一个文件
	ShouldTrimSpace  bool   `json:"shouldTrimSpace"`  // 是否需要去除空白字符
	ColSeparator     string `json:"colSeparator"`     // 列 分隔符 默认 `,`
	ReplaceCol       string `json:"replaceCol"`       //
	ReplaceLine      string `json:"replaceLine"`      //
	TxtFileType      string `json:"txtFileType"`      //

	// 数据库 配置
	DbConfig    *db.Config `json:"-"`
	DialectType string     `json:"databaseType"`

	Username string `json:"username"`
	Password string `json:"password"`

	EsConfig *elasticsearch.Config `json:"-"`

	RedisConfig *redis.Config `json:"-"`

	KafkaConfig *kafka.Config `json:"-"`

	OwnerName string `json:"ownerName"`
	TableName string `json:"tableName"`
	BySql     bool   `json:"bySql"` // 根据 SQL 语句导出
	SelectSql string `json:"selectSql"`

	IndexName     string `json:"indexName"`
	IndexIdName   string `json:"indexIdName"`
	IndexIdScript string `json:"indexIdScript"`

	TopicName      string `json:"topicName"`
	TopicGroupName string `json:"topicGroupName"`
	TopicKey       string `json:"topicKey"`
	TopicValue     string `json:"topicValue"`

	DataList []map[string]interface{} `json:"dataList"`

	Total int64 `json:"total"`

	ColumnList []*Column `json:"columnList"`

	AllOwner       bool       `json:"allOwner"`
	Owners         []*DbOwner `json:"owners"`
	SkipOwnerNames []string   `json:"skipOwnerNames"`

	FilePath    string `json:"filePath"`
	ShouldOwner bool   `json:"shouldOwner"` // 需要 建库
	ShouldTable bool   `json:"shouldTable"` // 需要 建表

	FileNameSplice string `json:"fileNameSplice"` // 文件名拼接字符 如：/ :库作为目录 表作为名称 默认
	FileName       string `json:"fileName"`
	RowNumber      int64  `json:"rowNumber"`
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
	From           *dialect.OwnerModel `json:"from"`
	To             *dialect.OwnerModel `json:"to"`
	SkipTableNames []string            `json:"skipTableNames"`
	AllTable       bool                `json:"allTable"`
	Tables         []*DbTable          `json:"tables"`
	fromService    db.IService
	toService      db.IService
	appended       bool
}

type DbTable struct {
	From            *dialect.TableModel `json:"from"`
	To              *dialect.TableModel `json:"to"`
	Columns         []*DbColumn         `json:"columns"`
	SkipColumnNames []string            `json:"skipColumnNames"`
	AllColumn       bool                `json:"allColumn"`
	appended        bool

	IndexIdName   string `json:"indexIdName"`
	IndexIdScript string `json:"indexIdScript"`

	TopicGroupName string `json:"topicGroupName"`
	TopicKey       string `json:"topicKey"`
	TopicValue     string `json:"topicValue"`
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
	return table
}

type DbColumn struct {
	From  *dialect.ColumnModel `json:"from"`
	To    *dialect.ColumnModel `json:"to"`
	Value string               `json:"value"`
}
