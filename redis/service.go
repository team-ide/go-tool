package redis

import (
	"golang.org/x/crypto/ssh"
	"strings"
)

type Config struct {
	Address   string      `json:"address"`
	Auth      string      `json:"auth"`
	Username  string      `json:"username"`
	CertPath  string      `json:"certPath"`
	Servers   []string    `json:"servers"`
	SSHClient *ssh.Client `json:"-"`
}

// New 创建Redis服务
func New(config *Config) (service IService, err error) {
	if !strings.Contains(config.Address, ",") && !strings.Contains(config.Address, ";") {
		service, err = NewRedisService(config)
	} else {
		if strings.Contains(config.Address, ",") {
			config.Servers = strings.Split(config.Address, ",")
		} else if strings.Contains(config.Address, ";") {
			config.Servers = strings.Split(config.Address, ";")
		} else {
			config.Servers = []string{config.Address}
		}
		service, err = NewClusterService(config)
	}
	return
}
