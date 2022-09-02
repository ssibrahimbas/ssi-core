package auth

import (
	"github.com/gofiber/fiber/v2"
	orgJwt "github.com/golang-jwt/jwt"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/jwt"
	"github.com/ssibrahimbas/ssi-core/pkg/result"
)

type CurrentUser struct {
	ID    string `json:"uuid"`
	Email string `json:"email"`
}

type CurrentUserConfig struct {
	Jwt    *jwt.Jwt
	I18n   *i18n.I18n
	MsgKey string
}

func NewCurrentUser(cnf *CurrentUserConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		t := c.Cookies("token")
		if t == "" {
			return c.Next()
		}
		res, err := cnf.Jwt.Parse(t)
		if err != nil {
			l := c.Locals("lang").(string)
			a := c.Locals("accept-language").(string)
			return result.Error(cnf.I18n.Translate(cnf.MsgKey, l, a), fiber.StatusUnauthorized)
		}
		p := res.Claims.(orgJwt.MapClaims)["payload"]
		c.Locals("user", p)
		return c.Next()
	}
}
