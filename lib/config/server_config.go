package config

import (
	"github.com/pulkit-tanwar/dispense-quotes/lib/constants"
)

type Config struct {
	Env     string
	Host    string
	Port    int
	APIPath string
}

func NewConfig(env, host string, port int, apiPath string) *Config {
	return &Config{
		Env:     env,
		Host:    host,
		Port:    port,
		APIPath: apiPath,
	}
}

func DefaultConfig() *Config {
	return &Config{
		Env:     constants.DefaultEnv,
		Host:    constants.DefaultHost,
		Port:    constants.DefaultPort,
		APIPath: constants.DefaultAPIPath,
	}
}
