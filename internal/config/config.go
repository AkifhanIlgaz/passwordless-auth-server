package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Redis struct {
		URL string `yaml:"url"`
	} `yaml:"redis"`
	Port int `yaml:"port"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config/config.yaml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return &config, nil
}
