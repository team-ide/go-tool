package db

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_sqlite3"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := db_sqlite3.GetDSN(config.DatabasePath)
			db, err = db_sqlite3.Open(dsn)
			return
		},
		DialectName: db_sqlite3.GetDialect(),
		matches:     []string{"sqlite", "sqlite3"},
	})
	if err != nil {
		util.Logger.Error("init sqlite3 db error", zap.Error(err))
		panic("init sqlite3 db error:" + err.Error())
	}
}
