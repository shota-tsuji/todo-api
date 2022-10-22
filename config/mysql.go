package config

type MysqlConfig struct {
	User     string
	Password string
	Database string
}

func NewMysqlConfig(config Config) MysqlConfig {
	return config.MysqlConfig
}
