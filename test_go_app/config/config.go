package config

type Config struct {
	Database DatabaseConfig `yaml:"Database"`
}

type DatabaseConfig struct {
	Host string `yaml:"host"`
	Port uint32 `yaml:"port"`
	Name string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
