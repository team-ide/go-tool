//go:build !darwin

package db_type_odbc

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_odbc"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := config.OdbcDsn
			db, err = db_odbc.Open(dsn)
			return
		},
		DialectName: db_odbc.GetDialect(),
		Matches:     []string{"odbc"},
	})
	if err != nil {
		util.Logger.Error("init GBase db error", zap.Error(err))
		panic("init GBase db error:" + err.Error())
	}
}
