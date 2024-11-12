package db_type_sqlite

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_sqlite3"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := config.Dsn
			if dsn == "" {
				dsn = "file:" + config.DatabasePath + "?cache=shared"
				if config.Username != "" {
					dsn += "&_auth_user=" + config.Username
				}
				if config.Password != "" {
					dsn += "&_auth_pass=" + config.Password
				}
				if config.Username != "" && config.Password != "" {
					dsn += "&_auth=true"
				}
			}
			db, err = db_sqlite3.Open(dsn)
			return
		},
		DialectName: db_sqlite3.GetDialect(),
		Matches:     []string{"sqlite", "sqlite3"},
	})
	if err != nil {
		util.Logger.Error("init sqlite3 db error", zap.Error(err))
		panic("init sqlite3 db error:" + err.Error())
	}
}
