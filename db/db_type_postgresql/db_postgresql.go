package db_type_postgresql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"github.com/lib/pq"
	"github.com/team-ide/go-driver/db_postgresql"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/db/db_dialer"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

func init() {
	err := db.AddDatabaseType(&db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := db_postgresql.GetDSN(config.Username, config.Password, config.Host, config.Port, config.DbName)

			if config.SSHClient != nil {
				db = sql.OpenDB(NewDialerConnector(config.SSHClient, dsn))
			} else {
				db, err = db_postgresql.Open(dsn)
			}
			return
		},
		DialectName: db_postgresql.GetDialect(),
		Matches:     []string{"postgresql", "ps"},
	})
	if err != nil {
		util.Logger.Error("init postgresql db error", zap.Error(err))
		panic("init postgresql db error:" + err.Error())
	}
}

type DialerDriver struct {
	SSHClient *ssh.Client `json:"-"`
}

func (d *DialerDriver) Open(dsn string) (driver.Conn, error) {
	if d.SSHClient != nil {
		return pq.DialOpen(db_dialer.NewSSHDialer(d.SSHClient), dsn)
	}
	return pq.Open(dsn)
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
