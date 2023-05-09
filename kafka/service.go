package kafka

// Config kafka配置
type Config struct {
	Address  string `json:"address"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	CertPath string `json:"certPath,omitempty"`
}

// New 创建kafka服务
func New(config *Config) (IService, error) {
	service := &Service{
		Config: config,
	}
	err := service.init()
	if err != nil {
		return nil, err
	}
	return service, nil
}
