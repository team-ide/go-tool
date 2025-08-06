package db_type_opengauss

import (
	"context"
	"database/sql"
	"database/sql/driver"
	openGauss "gitee.com/opengauss/openGauss-connector-go-pq"
	"github.com/team-ide/go-driver/db_opengauss"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/db/db_dialer"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
	"strings"
)

func init() {
	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := db_opengauss.GetDSN(config.Username, config.Password, config.Host, config.Port, config.DbName)
			if config.DsnAppend != "" {
				if !strings.HasPrefix(config.DsnAppend, "&") {
					dsn += "&"
				}
				dsn += config.DsnAppend
			}
			if config.SSHClient != nil {
				db = sql.OpenDB(NewDialerConnector(config.SSHClient, dsn))
			} else {
				db, err = db_opengauss.Open(dsn)
			}
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

type DialerDriver struct {
	SSHClient *ssh.Client `json:"-"`
}

func (d *DialerDriver) Open(dsn string) (driver.Conn, error) {
	if d.SSHClient != nil {
		return openGauss.DialOpen(db_dialer.NewSSHDialer(d.SSHClient), dsn)
	}
	return openGauss.Open(dsn)
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
