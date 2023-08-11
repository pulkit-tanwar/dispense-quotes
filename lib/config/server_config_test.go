package config_test

import (
	"testing"

	"github.com/pulkit-tanwar/dispense-quotes/lib/config"
	"github.com/pulkit-tanwar/dispense-quotes/lib/constants"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	c := config.NewConfig("staging", "example.com", 4000, "/api/v1/ping")
	assert.Equal(t, "staging", c.Env)
	assert.Equal(t, "/api/v1/ping", c.APIPath)
	assert.Equal(t, "example.com", c.Host)
	assert.Equal(t, 4000, c.Port)
}

func TestDefaultConfig(t *testing.T) {
	c := config.DefaultConfig()
	assert.Equal(t, constants.DefaultAPIPath, c.APIPath)
	assert.Equal(t, constants.DefaultEnv, c.Env)
	assert.Equal(t, constants.DefaultHost, c.Host)
	assert.Equal(t, constants.DefaultPort, c.Port)
}
