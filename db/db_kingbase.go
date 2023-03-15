package db

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_kingbase_v8r6"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := db_kingbase_v8r6.GetDSN(config.Username, config.Password, config.Host, config.Port, config.DbName)
			if config.Schema != "" {
				dsn += "&search_path=" + config.Schema
			}
			db, err = db_kingbase_v8r6.Open(dsn)
			return
		},
		DialectName: db_kingbase_v8r6.GetDialect(),
		matches:     []string{"KingBase", "kb"},
	})
	if err != nil {
		util.Logger.Error("init KingBase db error", zap.Error(err))
		panic("init KingBase db error:" + err.Error())
	}
}
