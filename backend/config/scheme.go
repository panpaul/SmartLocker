package config

type Config struct {
	WebServer struct {
		Address   string `yaml:"Address"`
		Port      string `yaml:"Port"`
		JwtSecret string `yaml:"JwtSecret"`
	} `yaml:"WebServer,flow"`
	Database struct {
		Type     string `yaml:"Type"`
		Address  string `yaml:"Address"`
		Port     string `yaml:"Port"`
		User     string `yaml:"User"`
		Password string `yaml:"Password"`
		Database string `yaml:"Database"`
		Prefix   string `yaml:"Prefix"`
	} `yaml:"Database,flow"`
	Mode string `yaml:"Mode"`
}
