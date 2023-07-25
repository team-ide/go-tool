//go:build !darwin

package db_type_oracle

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_oracle"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := db_oracle.GetDSN(config.Username, config.Password, config.Host, config.Port, config.Sid)
			db, err = db_oracle.Open(dsn)
			return
		},
		DialectName: db_oracle.GetDialect(),
		Matches:     []string{"oracle"},
	})
	if err != nil {
		util.Logger.Error("init oracle db error", zap.Error(err))
		panic("init oracle db error:" + err.Error())
	}
}
