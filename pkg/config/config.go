package config

import (
	"github.com/golobby/dotenv"
	"github.com/golobby/dotenv/pkg/decoder"
	"os"
)

type Config struct {
	d *decoder.Decoder
}

func New(fn string) (*Config, error) {
	file, err := os.Open(fn)
	if err != nil {
		return nil, err
	}
	return &Config{
		d: dotenv.NewDecoder(file),
	}, nil
}

func (c *Config) Decode(o interface{}) error {
	return c.d.Decode(o)
}
