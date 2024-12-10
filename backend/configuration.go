package main

import (
	"gopkg.in/yaml.v3"
	"os"
)

const defaultConfigPath = "configs/app.yml"

type AppConfiguration struct {
	Logger   LoggerConfiguration   `yaml:"logger"`
	Server   ServerConfiguration   `yaml:"server"`
	Database DatabaseConfiguration `yaml:"db"`
}

func NewAppConfiguration() *AppConfiguration {
	var cfg AppConfiguration

	filepath, ok := os.LookupEnv("APP_CONFIGURATION_FILE")
	if !ok {
		filepath = defaultConfigPath
	}

	b, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	// we expand environment variable if used in the file to avoid passing
	// secrets in clear.
	expCfg := os.ExpandEnv(string(b))

	err = yaml.Unmarshal([]byte(expCfg), &cfg)
	if err != nil {
		panic(err)
	}

	return &cfg
}
