package kafka

// Config kafka配置
type Config struct {
	Address  string `json:"address" yaml:"address"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	CertPath string `json:"certPath,omitempty" yaml:"certPath,omitempty"`
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
