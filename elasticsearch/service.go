package elasticsearch

type Config struct {
	Url      string `json:"url,omitempty" yaml:"url,omitempty"`
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
	CertPath string `json:"certPath,omitempty" yaml:"certPath,omitempty"`
}

func New(config *Config) (IService, error) {
	service := &V7Service{
		Config: config,
	}
	err := service.init()
	return service, err
}
