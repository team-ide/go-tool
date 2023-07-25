package db_type_postgresql

import (
	"database/sql"
	"github.com/team-ide/go-driver/db_opengauss"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
)

func init() {
	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := db_opengauss.GetDSN(config.Username, config.Password, config.Host, config.Port, config.DbName)
			db, err = db_opengauss.Open(dsn)
			return
		},
		DialectName: db_opengauss.GetDialect(),
		Matches:     []string{"opengauss"},
	})
	if err != nil {
		util.Logger.Error("init OpenGauss db error", zap.Error(err))
		panic("init OpenGauss db error:" + err.Error())
	}
}
