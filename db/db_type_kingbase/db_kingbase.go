package db_type_kingbase

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/team-ide/go-driver/db_kingbase_v8r6"
	"github.com/team-ide/go-driver/driver/kingbase/v8r6/kingbase.com/gokb"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/db/db_dialer"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"net/url"
)

func init() {

	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			password := url.PathEscape(config.Password)
			dsn := fmt.Sprintf("kingbase://%s:%s@%s:%d/%s?sslmode=disable", config.Username, password, config.Host, config.Port, config.DbName)
			if config.Schema != "" {
				dsn += "&search_path=" + config.Schema
			}
			if config.SSHClient != nil {
				db = sql.OpenDB(NewDialerConnector(config.SSHClient, dsn))
			} else {
				db, err = db_kingbase_v8r6.Open(dsn)
			}
			return
		},
		DialectName: db_kingbase_v8r6.GetDialect(),
		Matches:     []string{"KingBase", "kb"},
	})
	if err != nil {
		util.Logger.Error("init KingBase db error", zap.Error(err))
		panic("init KingBase db error:" + err.Error())
	}
}

type DialerDriver struct {
	SSHClient *ssh.Client `json:"-"`
}

func (d *DialerDriver) Open(dsn string) (driver.Conn, error) {
	if d.SSHClient != nil {
		return gokb.DialOpen(db_dialer.NewSSHDialer(d.SSHClient), dsn)
	}
	return gokb.Open(dsn)
}

func NewDialerConnector(SSHClient *ssh.Client, dsn string) *DialerConnector {

	return &DialerConnector{
		dsn: dsn,
		driver: &DialerDriver{
			SSHClient: SSHClient,
		},
	}
}

type DialerConnector struct {
	dsn    string
	driver driver.Driver
}

func (t *DialerConnector) Connect(_ context.Context) (driver.Conn, error) {
	return t.driver.Open(t.dsn)
}

func (t *DialerConnector) Driver() driver.Driver {
	return t.driver
}
