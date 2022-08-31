package i18n

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestI18n_Http(t *testing.T) {
	t.Run("Should context have query language or Accept-Language", func(t *testing.T) {
		i := New("en")
		i.LoadLanguages("./locales_test", "en", "tr")
		fa := fiber.New()
		fa.Get("/query-and-accept", func(c *fiber.Ctx) error {
			l, a := i.GetLanguagesInContext(c)
			assert.Equal(t, "tr", l)
			assert.Equal(t, "en", a)
			return c.Status(fiber.StatusOK).SendString("OK")
		})
		req, err := http.NewRequest("GET", "/query-and-accept?lang=tr", nil)
		assert.NoError(t, err)
		resp, _ := fa.Test(req)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Should context have query language or Accept-Language", func(t *testing.T) {
		i := New("en")
		i.LoadLanguages("./locales_test", "en", "tr")
		fa := fiber.New()
		fa.Use(i.I18nMiddleware)
		fa.Get("/query-and-accept", func(c *fiber.Ctx) error {
			l := c.Locals("lang").(string)
			a := c.Locals("accept-language").(string)
			assert.Equal(t, "tr", l)
			assert.Equal(t, "en", a)
			return c.Status(fiber.StatusOK).JSON(map[string]string{l: l, a: a})
		}).Use(i.I18nMiddleware)
		req, err := http.NewRequest("GET", "/query-and-accept?lang=tr", nil)
		assert.NoError(t, err)
		resp, _ := fa.Test(req)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}
