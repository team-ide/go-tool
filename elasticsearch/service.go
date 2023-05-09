package elasticsearch

type Config struct {
	Url      string `json:"url,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	CertPath string `json:"certPath,omitempty"`
}

func New(config *Config) (IService, error) {
	service := &V7Service{
		Config: config,
	}
	err := service.init()
	return service, err
}
