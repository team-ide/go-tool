package redis

import "strings"

type Config struct {
	Address  string `json:"address"`
	Auth     string `json:"auth"`
	Username string `json:"username"`
	CertPath string `json:"certPath"`
}

// New 创建Redis服务
func New(config Config) (service IService, err error) {
	if !strings.Contains(config.Address, ",") && !strings.Contains(config.Address, ";") {
		service, err = NewRedisService(config.Address, config.Username, config.Auth, config.CertPath)
	} else {
		var servers []string
		if strings.Contains(config.Address, ",") {
			servers = strings.Split(config.Address, ",")
		} else if strings.Contains(config.Address, ";") {
			servers = strings.Split(config.Address, ";")
		} else {
			servers = []string{config.Address}
		}
		service, err = NewClusterService(servers, config.Username, config.Auth, config.CertPath)
	}
	return
}
