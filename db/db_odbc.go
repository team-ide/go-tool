//go:build !darwin

package db

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_odbc"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := config.OdbcDsn
			db, err = db_odbc.Open(dsn)
			return
		},
		DialectName: db_odbc.GetDialect(),
		matches:     []string{"odbc"},
	})
	if err != nil {
		util.Logger.Error("init GBase db error", zap.Error(err))
		panic("init GBase db error:" + err.Error())
	}
}
