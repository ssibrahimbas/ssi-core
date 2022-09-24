package config

import (
	"os"

	"github.com/golobby/dotenv"
	"github.com/spf13/viper"
	"github.com/ssibrahimbas/ssi-core/pkg/helper"
)

func LoadConfig(p string, c interface{}) {
	viper.AddConfigPath(p)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	helper.CheckErr(err)
	err = viper.Unmarshal(&c)
	helper.CheckErr(err)
}

func LoadEnv(p string, c interface{}) {
	file, err := os.Open(p)
	helper.CheckErr(err)
	err = dotenv.NewDecoder(file).Decode(&c)
	helper.CheckErr(err)
}
