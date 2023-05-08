package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/team-ide/go-driver/db_mysql"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"net"
)

func init() {
	err := addDatabaseType(&DatabaseType{
		newDb: func(config *Config) (db *sql.DB, err error) {
			dsn := db_mysql.GetDSN(config.Username, config.Password, config.Host, config.Port, config.Database)
			if config.SSHClient != nil {
				// 填写注册的mysql网络
				dsn = fmt.Sprintf("%s:%s@mysql+ssh(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
					config.Username, config.Password, config.Host, config.Port, config.Database)
				// 注册ssh代理
				mysql.RegisterDialContext("mysql+ssh", func(ctx context.Context, addr string) (net.Conn, error) {
					return config.SSHClient.Dial("tcp", addr)
				})
			}
			db, err = db_mysql.Open(dsn)
			return
		},
		DialectName: db_mysql.GetDialect(),
		matches:     []string{"mysql"},
	})
	if err != nil {
		util.Logger.Error("init mysql db error", zap.Error(err))
		panic("init mysql db error:" + err.Error())
	}
}
