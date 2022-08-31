package auth

import (
	"github.com/gofiber/fiber/v2"
	ssiHttp "github.com/ssibrahimbas/ssi-core/pkg/http"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/jwt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestAuth_Module(t *testing.T) {
	h := ssiHttp.New()
	fa := h.App
	i := i18n.New("en")
	i.LoadLanguages("./locales_test", "en")
	j := jwt.New("secret")
	fa.Use(i.I18nMiddleware)
	token, err := j.Sign("test")
	assert.NoError(t, err)
	t.Run("CurrentUser Middleware Testing", func(t *testing.T) {
		g := fa.Group("/current-user")
		g.Use(NewCurrentUser(&CurrentUserConfig{
			Jwt:    j,
			I18n:   i,
			MsgKey: "auth_unauthorized",
		}))
		t.Run("Should pass validation if token is valid", func(t *testing.T) {
			g.Get("/", func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusOK).SendString("Hello World")
			})
			req, err := http.NewRequest("GET", "/current-user/", nil)
			req.Header.Add("Authorization", token)
			assert.NoError(t, err)
			resp, _ := fa.Test(req)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		})
		t.Run("Should return unauthorized if token is invalid", func(t *testing.T) {
			g.Get("/", func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusOK).SendString("Hello World")
			}).Use(i.I18nMiddleware)
			req, err := http.NewRequest("GET", "/current-user/", nil)
			req.Header.Add("Authorization", "invalid token")
			assert.NoError(t, err)
			resp, _ := fa.Test(req)
			resp.Body.Close()
			assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
		})
		t.Run("Should pass if token is empty", func(t *testing.T) {
			g.Get("/", func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusOK).SendString("Hello World")
			}).Use(i.I18nMiddleware)
			req, err := http.NewRequest("GET", "/current-user/", nil)
			assert.NoError(t, err)
			resp, _ := fa.Test(req)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		})
	})
	t.Run("RequiredAuth Middleware Testing", func(t *testing.T) {
		g := fa.Group("/required-auth")
		g.Use(NewCurrentUser(&CurrentUserConfig{
			Jwt:    j,
			I18n:   i,
			MsgKey: "auth_unauthorized",
		}))
		g.Use(NewRequiredAuth(&RequiredAuthConfig{
			I18n:   i,
			MsgKey: "auth_unauthorized",
		}))
		t.Run("Should pass validation if user is there", func(t *testing.T) {
			g.Get("/", func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusOK).SendString("Hello World")
			})
			req, err := http.NewRequest("GET", "/required-auth/", nil)
			req.Header.Add("Authorization", token)
			assert.NoError(t, err)
			resp, _ := fa.Test(req)
			assert.Equal(t, fiber.StatusOK, resp.StatusCode)
		})
		t.Run("Should return unauthorized if user is not there", func(t *testing.T) {
			g.Get("/", func(c *fiber.Ctx) error {
				return c.Status(fiber.StatusOK).SendString("Hello World")
			}).Use(i.I18nMiddleware)
			req, err := http.NewRequest("GET", "/required-auth/", nil)
			assert.NoError(t, err)
			resp, _ := fa.Test(req)
			assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
		})
	})
}