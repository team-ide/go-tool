package db

import (
	"database/sql"
	"strings"

	"github.com/team-ide/go-dialect/dialect"
)

type DatabaseType struct {
	DialectName string `json:"dialectName"`
	NewDb       func(config *Config) (db *sql.DB, err error)
	dia         dialect.Dialect
	Matches     []string
}

func (this_ *DatabaseType) init() (err error) {
	this_.dia, err = dialect.NewDialect(this_.DialectName)
	if err != nil {
		return
	}
	return
}

var (
	DatabaseTypes []*DatabaseType
)

func AddDatabaseType(databaseType *DatabaseType) (err error) {
	err = databaseType.init()
	if err != nil {
		return
	}
	DatabaseTypes = append(DatabaseTypes, databaseType)
	return
}

func GetDatabaseType(databaseType string) *DatabaseType {
	for _, one := range DatabaseTypes {
		if strings.EqualFold(databaseType, one.DialectName) {
			return one
		}
		for _, match := range one.Matches {
			if strings.EqualFold(databaseType, match) {
				return one
			}
		}
	}
	return nil
}

type SqlConditionalOperation struct {
	Value string `json:"value,omitempty"`
	Text  string `json:"text,omitempty"`
}

var (
	SqlConditionalOperations []*SqlConditionalOperation
)

func init() {
	SqlConditionalOperations = []*SqlConditionalOperation{
		{Text: "等于", Value: "="},
		{Text: "不等于", Value: "<>"},
		{Text: "大于", Value: ">"},
		{Text: "大于或等于", Value: ">="},
		{Text: "小于", Value: "<"},
		{Text: "小于或等于", Value: "<="},
		{Text: "包含", Value: "like"},
		{Text: "不包含", Value: "not like"},
		{Text: "开始以", Value: "like start"},
		{Text: "开始不是以", Value: "not like start"},
		{Text: "结束以", Value: "like end"},
		{Text: "结束不是以", Value: "not like end"},
		{Text: "是null", Value: "is null"},
		{Text: "不是null", Value: "is not null"},
		{Text: "是空", Value: "is empty"},
		{Text: "不是空", Value: "is not empty"},
		{Text: "介于", Value: "between"},
		{Text: "不介于", Value: "not between"},
		{Text: "在列表", Value: "in"},
		{Text: "不在列表", Value: "not in"},
		{Text: "自定义", Value: "custom"},
	}
}
func GetSqlConditionalOperations() []*SqlConditionalOperation {
	return SqlConditionalOperations
}
