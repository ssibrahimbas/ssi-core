package config

import (
	"github.com/ssibrahimbas/ssi-core/pkg/helper"
	"go.deanishe.net/env"
)

func LoadConfig(c interface{}) {
	err := env.Bind(c)
	helper.CheckErr(err)
}
