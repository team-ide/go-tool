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

	FilePath         string `json:"filePath"`
	FileSuffix       string `json:"fileSuffix"`
	SqlFileMergeType string `json:"sqlFileMergeType"` // SQL 的文件合并类型 如：one：一个文件， owner：每个库一个文件，table：每个表一个文件
	ShouldOwner      bool   `json:"shouldOwner"`      // 需要 建库
	ShouldTable      bool   `json:"shouldTable"`      // 需要 建表

	FileNameSplice string `json:"fileNameSplice"` // 文件名拼接字符 如：/ :库作为目录 表作为名称 默认
	FileName       string `json:"fileName"`

	ErrorContinue bool  `json:"errorContinue"`
	BatchNumber   int64 `json:"batchNumber"`

	ColSeparator string `json:"colSeparator"` // 列 分隔符 默认 `,`

	ReplaceSeparators map[string]string `json:"replaceSeparators"` // 替换字符，如将：`\n` 替换为 `|:-n-:|`，`,` 替换为 `|:-，-:|`，写入时候 将 key 替换为 value，读取时候将 value 替换为 key
	ShouldTrimSpace   bool              `json:"shouldTrimSpace"`   // 是否需要去除空白字符

	*dialect.ParamModel
}

func (this_ *Options) GetColSeparator() string {
	if this_.ColSeparator == "" {
		return ","
	}
	return this_.ColSeparator
}

func (this_ *Options) GetDialectParam() *dialect.ParamModel {
	if this_.ParamModel == nil {
		this_.ParamModel = &dialect.ParamModel{}
	}
	return this_.ParamModel
}

func (this_ *Options) GetFileSuffix() string {
	if this_.FileSuffix == "" {
		return "txt"
	}
	return this_.FileSuffix
}

func (this_ *Options) GetFileName() string {
	if this_.FileName == "" {
		return "导出"
	}
	return this_.FileName
}

type DataSourceConfig struct {
	Type string `json:"type"`

	// 数据库 配置
	DbConfig    *db.Config `json:"-"`
	DialectType string     `json:"databaseType"`

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
	appended       bool
}

type DbTable struct {
	SourceName      string      `json:"sourceName"`
	TargetName      string      `json:"targetName"`
	Columns         []*DbColumn `json:"columns"`
	SkipColumnNames []string    `json:"skipColumnNames"`
	AllColumn       bool        `json:"allColumn"`
	table           *dialect.TableModel
	appended        bool
}

func (this_ *DbTable) GetTargetDialectTable() *dialect.TableModel {
	table := &dialect.TableModel{}
	table.TableName = this_.TargetName

	for _, c := range this_.Columns {
		column := &dialect.ColumnModel{}
		column.ColumnName = c.TargetName
		if c.Column != nil && c.Column.ColumnModel != nil {
			column.ColumnDataType = c.Column.ColumnModel.ColumnDataType
			column.ColumnComment = c.Column.ColumnModel.ColumnComment
			column.ColumnDefault = c.Column.ColumnModel.ColumnDefault
			column.ColumnExtra = c.Column.ColumnModel.ColumnExtra
			column.ColumnEnums = c.Column.ColumnModel.ColumnEnums
			column.ColumnLength = c.Column.ColumnModel.ColumnLength
			column.ColumnNotNull = c.Column.ColumnModel.ColumnNotNull
			column.ColumnCharacterSetName = c.Column.ColumnModel.ColumnCharacterSetName
			column.ColumnPrecision = c.Column.ColumnModel.ColumnPrecision
			column.ColumnScale = c.Column.ColumnModel.ColumnScale
		}
		table.ColumnList = append(table.ColumnList, column)
	}
	if this_.table != nil {
		table.PrimaryKeys = this_.table.PrimaryKeys
		table.IndexList = this_.table.IndexList
	}
	return table
}

type DbColumn struct {
	*Column
	SourceName string `json:"sourceName"`
	TargetName string `json:"targetName"`
	Value      string `json:"value"`
}

func (this_ DataSourceConfig) IsData() bool {
	return this_.Type == "data"
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
