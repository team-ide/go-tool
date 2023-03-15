//go:build !darwin

package db

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_oracle"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := db_oracle.GetDSN(config.Username, config.Password, config.Host, config.Port, config.Sid)
			db, err = db_oracle.Open(dsn)
			return
		},
		DialectName: db_oracle.GetDialect(),
		matches:     []string{"oracle"},
	})
	if err != nil {
		util.Logger.Error("init oracle db error", zap.Error(err))
		panic("init oracle db error:" + err.Error())
	}
}
