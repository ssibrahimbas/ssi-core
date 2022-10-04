package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	type Config struct {
		Label string `env:"LABEL"`
	}
	os.Setenv("LABEL", "test")
	c := Config{}
	LoadConfig(&c)
	assert.Equal(t, "test", c.Label)
}
