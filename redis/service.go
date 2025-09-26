package redis

import (
	"strings"

	"golang.org/x/crypto/ssh"
)

type Config struct {
	Address            string      `json:"address" yaml:"address"`
	Auth               string      `json:"auth" yaml:"auth"`
	Username           string      `json:"username" yaml:"username"`
	CertPath           string      `json:"certPath" yaml:"certPath"`
	Servers            []string    `json:"servers" yaml:"servers"`
	SSHClient          *ssh.Client `json:"-"`
	ThrowNotFoundErr   bool        `json:"throwNotFoundErr" yaml:"throwNotFoundErr"`     // 是否 抛出 值 不存在异常
	InsecureSkipVerify bool        `json:"insecureSkipVerify" yaml:"insecureSkipVerify"` // TLS 跳过验证
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
