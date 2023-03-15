//go:build !darwin

package db

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_gbase"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := config.OdbcDsn
			db, err = db_gbase.Open(dsn)
			return
		},
		DialectName: db_gbase.GetDialect(),
		matches:     []string{"gbase"},
	})
	if err != nil {
		util.Logger.Error("init GBase db error", zap.Error(err))
		panic("init GBase db error:" + err.Error())
	}
}
