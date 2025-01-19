package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Redis Redis `yaml:"redis"`
	OTP   OTP   `yaml:"otp"`
	Port  int   `yaml:"port"`
}

type Redis struct {
	URL string `yaml:"url"`
}

type OTP struct {
	Length int `yaml:"length"`
	Ttl    int `yaml:"ttl"`
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
