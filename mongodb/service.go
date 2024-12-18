package mongodb

type Config struct {
	Address        string `json:"address" yaml:"address"`
	Username       string `json:"username,omitempty" yaml:"username,omitempty"`
	Password       string `json:"password,omitempty" yaml:"password,omitempty"`
	MinPoolSize    int    `json:"minPoolSize,omitempty" yaml:"minPoolSize,omitempty"`
	MaxPoolSize    int    `json:"maxPoolSize,omitempty" yaml:"maxPoolSize,omitempty"`
	ConnectTimeout int    `json:"connectTimeout,omitempty" yaml:"connectTimeout,omitempty"` // 客户端连接超时 单位 毫秒
	CertPath       string `json:"certPath,omitempty" yaml:"certPath,omitempty"`
}

// New 创建 mongodb 客户端
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
