package db

import (
	"database/sql"
	"strings"

	"github.com/team-ide/go-dialect/dialect"
)

type DatabaseType struct {
	DialectName string `json:"dialectName"`
	newDb       func(config *Config) (db *sql.DB, err error)
	dia         dialect.Dialect
	matches     []string
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

func addDatabaseType(databaseType *DatabaseType) (err error) {
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
		for _, match := range one.matches {
			if strings.EqualFold(databaseType, match) {
				return one
			}
		}
	}
	return nil
}
