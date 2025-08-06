package db_type_mysql

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/team-ide/go-driver/db_mysql"
	"github.com/team-ide/go-tool/db"
	"github.com/team-ide/go-tool/util"
	"go.uber.org/zap"
	"net"
	"os"
	"strings"
	"sync"
)

var DatabaseType *db.DatabaseType

func init() {
	DatabaseType = &db.DatabaseType{
		NewDb: func(config *db.Config) (db *sql.DB, err error) {
			dsn := db_mysql.GetDSN(config.Username, config.Password, config.Host, config.Port, config.Database)
			tlsConfig, err := registerTLSConfig(config)
			if err != nil {
				return
			}
			if config.SSHClient != nil {
				// 填写注册的mysql网络
				dsn = fmt.Sprintf("%s:%s@mysql+ssh(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
					config.Username, config.Password, config.Host, config.Port, config.Database)
				// 注册ssh代理
				mysql.RegisterDialContext("mysql+ssh", func(ctx context.Context, addr string) (net.Conn, error) {
					conn, e := config.SSHClient.Dial("tcp", addr)
					return &util.SSHChanConn{Conn: conn}, e
				})
			}
			if tlsConfig != "" {
				dsn += "&tls=" + tlsConfig
				util.Logger.Info("mysql use TLS Config", zap.Any("tls", tlsConfig))
			}
			if config.DsnAppend != "" {
				if !strings.HasPrefix(config.DsnAppend, "&") {
					dsn += "&"
				}
				dsn += config.DsnAppend
			}
			db, err = db_mysql.Open(dsn)
			return
		},
		DialectName: db_mysql.GetDialect(),
		Matches:     []string{"mysql"},
	}

	err := db.AddDatabaseType(DatabaseType)
	if err != nil {
		util.Logger.Error("init mysql db error", zap.Error(err))
		panic("init mysql db error:" + err.Error())
	}
}

var tlsCache = map[string]string{}
var tlsCacheLock sync.Mutex

func registerTLSConfig(config *db.Config) (name string, err error) {
	if config.TlsConfig == "" || config.TlsConfig == "skip-verify" || config.TlsConfig == "preferred" {
		name = config.TlsConfig
		return
	}

	tlsCacheLock.Lock()
	defer tlsCacheLock.Unlock()
	util.Logger.Info("mysql register TLS Config", zap.Any("TlsRootCert", config.TlsRootCert))
	util.Logger.Info("mysql register TLS Config", zap.Any("TlsClientCert", config.TlsClientCert))
	util.Logger.Info("mysql register TLS Config", zap.Any("TlsClientKey", config.TlsClientKey))
	key := config.TlsRootCert + "|" + config.TlsClientCert + "|" + config.TlsClientKey
	name, ok := tlsCache[key]
	if ok {
		return
	}
	name = util.GetUUID()
	rootCertPool := x509.NewCertPool()
	var bs []byte
	if config.TlsRootCert != "" {
		bs, err = os.ReadFile(config.TlsRootCert)
		if err != nil {
			return
		}
		if ok = rootCertPool.AppendCertsFromPEM(bs); !ok {
			err = errors.New("root cert error, append certs from PEM error")
			return
		}
	}
	var clientCert []tls.Certificate

	if config.TlsClientCert != "" || config.TlsClientKey != "" {
		var cert tls.Certificate
		cert, err = tls.LoadX509KeyPair(config.TlsClientCert, config.TlsClientKey)
		if err != nil {
			err = errors.New("client cert or key error, " + err.Error())
			return
		}
		clientCert = append(clientCert, cert)
	}
	err = mysql.RegisterTLSConfig(name, &tls.Config{
		RootCAs:      rootCertPool,
		Certificates: clientCert,
	})
	if err != nil {
		return
	}
	tlsCache[key] = name
	return
}
