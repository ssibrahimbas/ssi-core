package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/result"
)

type RequiredAuthConfig struct {
	I18n   *i18n.I18n
	MsgKey string
}

func NewRequiredAuth(cnf *RequiredAuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		u := c.Locals("user")
		if u == nil {
			l := c.Locals("lang").(string)
			a := c.Locals("accept-language").(string)
			return result.Error(cnf.I18n.Translate(cnf.MsgKey, l, a), fiber.StatusUnauthorized)
		}
		return c.Next()
	}
}
