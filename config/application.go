package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	ServerConfig `yaml:"server"`
	MysqlConfig  `yaml:"mysql"`
}

func NewConfig() Config {
	config := &Config{}
	const configFileName = "application.yaml"

	file, err := os.Open(configFileName)
	if err != nil {
		log.Fatal("Failed to read ", configFileName)
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		log.Fatal("Failed to read ", configFileName)
	}

	return *config
}
