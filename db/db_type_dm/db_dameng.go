package db_type_dm

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_dm"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"strings"
)

func init() {
	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := db_dm.GetDSN(config.Username, config.Password, config.Host, config.Port, config.Schema)
			if config.DsnAppend != "" {
				if !strings.HasPrefix(config.DsnAppend, "&") {
					dsn += "&"
				}
				dsn += config.DsnAppend
			}
			db, err = db_dm.Open(dsn)
			return
		},
		DialectName: db_dm.GetDialect(),
		Matches:     []string{"DaMeng", "dm"},
	})
	if err != nil {
		util.Logger.Error("init DM db error", zap.Error(err))
		panic("init DM db error:" + err.Error())
	}
}
