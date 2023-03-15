//go:build !darwin

package db

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_shentong"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := db_shentong.GetDSN(config.Username, config.Password, config.Host, config.Port, config.DbName)
			db, err = db_shentong.Open(dsn)
			return
		},
		DialectName: db_shentong.GetDialect(),
		matches:     []string{"ShenTong", "st"},
	})
	if err != nil {
		util.Logger.Error("init ShenTong db error", zap.Error(err))
		panic("init ShenTong db error:" + err.Error())
	}
}
