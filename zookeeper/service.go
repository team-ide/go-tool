package zookeeper

import (
	"golang.org/x/crypto/ssh"
)

type Config struct {
	Address   string      `json:"address"`
	Username  string      `json:"username,omitempty"`
	Password  string      `json:"password,omitempty"`
	Timeout   int         `json:"timeout,omitempty"`
	SSHClient *ssh.Client `json:"-"`
}

// New 创建zookeeper客户端
func New(config *Config) (IService, error) {
	service := &Service{
		Config: config,
	}
	err := service.init(config.SSHClient)
	if err != nil {
		return nil, err
	}
	return service, nil
}
