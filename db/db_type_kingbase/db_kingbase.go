package db_type_kingbase

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_kingbase_v8r6"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := db_kingbase_v8r6.GetDSN(config.Username, config.Password, config.Host, config.Port, config.DbName)
			if config.Schema != "" {
				dsn += "&search_path=" + config.Schema
			}
			db, err = db_kingbase_v8r6.Open(dsn)
			return
		},
		DialectName: db_kingbase_v8r6.GetDialect(),
		Matches:     []string{"KingBase", "kb"},
	})
	if err != nil {
		util.Logger.Error("init KingBase db error", zap.Error(err))
		panic("init KingBase db error:" + err.Error())
	}
}
