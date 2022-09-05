package config

import (
	"github.com/spf13/viper"
)

func LoadConfig(p string, c interface{}) interface{} {
	viper.AddConfigPath(p)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	viper.ReadInConfig()
	viper.Unmarshal(&c)
	return c
}
