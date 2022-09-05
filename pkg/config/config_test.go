package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type Config struct {
		Label string `mapstructure:"LABEL"`
	}
	c := Config{}
	LoadConfig(".", &c)
	assert.Equal(t, "test", c.Label)
}
