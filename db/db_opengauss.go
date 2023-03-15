package db

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_opengauss"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := db_opengauss.GetDSN(config.Username, config.Password, config.Host, config.Port, config.DbName)
			db, err = db_opengauss.Open(dsn)
			return
		},
		DialectName: db_opengauss.GetDialect(),
		matches:     []string{"opengauss"},
	})
	if err != nil {
		util.Logger.Error("init OpenGauss db error", zap.Error(err))
		panic("init OpenGauss db error:" + err.Error())
	}
}
