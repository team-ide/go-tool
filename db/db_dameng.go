package db

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_dm"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := db_dm.GetDSN(config.Username, config.Password, config.Host, config.Port)
			if config.Schema != "" {
				dsn += "&schema=" + config.Schema
			}
			db, err = db_dm.Open(dsn)
			return
		},
		DialectName: db_dm.GetDialect(),
		matches:     []string{"DaMeng", "dm"},
	})
	if err != nil {
		util.Logger.Error("init DaMeng db error", zap.Error(err))
		panic("init DaMeng db error:" + err.Error())
	}
}
