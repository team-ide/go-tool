//go:build !darwin && !arm64

package db_type_gbase

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_gbase"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func initDatabase() {
	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := config.OdbcDsn
			db, err = db_gbase.Open(dsn)
			return
		},
		DialectName: db_gbase.GetDialect(),
		Matches:     []string{"gbase"},
	})
	if err != nil {
		util.Logger.Error("init GBase db error", zap.Error(err))
		panic("init GBase db error:" + err.Error())
	}
}
