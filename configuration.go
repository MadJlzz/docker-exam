package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

const defaultConfigPath = "configs/app.yml"

type AppConfiguration struct {
	Logger LoggerConfiguration `yaml:"logger"`
	Server ServerConfiguration `yaml:"server"`
}

func NewAppConfiguration() *AppConfiguration {
	var cfg AppConfiguration

	filepath, ok := os.LookupEnv("APP_CONFIGURATION_FILE")
	if !ok {
		filepath = defaultConfigPath
	}

	f, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	err = yaml.NewDecoder(f).Decode(&cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
