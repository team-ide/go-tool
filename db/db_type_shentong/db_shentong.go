//go:build !darwin

package db_type_shentong

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_shentong"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"strings"
)

func initDatabase() {
	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := db_shentong.GetDSN(config.Username, config.Password, config.Host, config.Port, config.DbName)
			if config.DsnAppend != "" {
				if !strings.HasPrefix(config.DsnAppend, "&") {
					dsn += "&"
				}
				dsn += config.DsnAppend
			}
			db, err = db_shentong.Open(dsn)
			return
		},
		DialectName: db_shentong.GetDialect(),
		Matches:     []string{"ShenTong", "st"},
	})
	if err != nil {
		util.Logger.Error("init ShenTong db error", zap.Error(err))
		panic("init ShenTong db error:" + err.Error())
	}
}
