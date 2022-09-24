package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	type Config struct {
		Label string `mapstructure:"LABEL"`
	}
	c := Config{}
	LoadConfig(".", &c)
	assert.Equal(t, "test", c.Label)
}

func TestLoadEnv(t *testing.T) {
	type Config struct {
		Label string `env:"LABEL"`
	}
	c := Config{}
	LoadEnv("app.env", &c)
	assert.Equal(t, "test", c.Label)
}
