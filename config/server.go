package config

type ServerConfig struct {
	Host string
	Port string
}

func NewServerConfig(config Config) ServerConfig {
	return config.ServerConfig
}
