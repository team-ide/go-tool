package db

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_postgresql"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := db_postgresql.GetDSN(config.Username, config.Password, config.Host, config.Port, config.DbName)
			db, err = db_postgresql.Open(dsn)
			return
		},
		DialectName: db_postgresql.GetDialect(),
		matches:     []string{"postgresql", "ps"},
	})
	if err != nil {
		util.Logger.Error("init postgresql db error", zap.Error(err))
		panic("init postgresql db error:" + err.Error())
	}
}
