package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Server struct {
		Host string
		Port string
	}
	Mysql struct {
		User     string
		Password string
		Database string
	}
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
