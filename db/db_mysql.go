package db

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_mysql"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := db_mysql.GetDSN(config.Username, config.Password, config.Host, config.Port, config.Database)
			db, err = db_mysql.Open(dsn)
			return
		},
		DialectName: db_mysql.GetDialect(),
		matches:     []string{"mysql"},
	})
	if err != nil {
		util.Logger.Error("init mysql db error", zap.Error(err))
		panic("init mysql db error:" + err.Error())
	}
}
