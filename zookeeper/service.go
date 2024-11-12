package zookeeper

import (
	"golang.org/x/crypto/ssh"
)

type Config struct {
	Address           string      `json:"address" yaml:"address"`
	Username          string      `json:"username,omitempty" yaml:"username,omitempty"`
	Password          string      `json:"password,omitempty" yaml:"password,omitempty"`
	SessionTimeout    int         `json:"sessionTimeout,omitempty" yaml:"sessionTimeout,omitempty"`       // 会话超时 单位 毫秒
	ConnectionTimeout int         `json:"connectionTimeout,omitempty" yaml:"connectionTimeout,omitempty"` // 客户端连接超时 单位 毫秒
	SSHClient         *ssh.Client `json:"-"`
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
