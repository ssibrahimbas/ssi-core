package i18n

import "github.com/gofiber/fiber/v2"

func (i *I18n) GetLanguagesInContext(c *fiber.Ctx) (string, string) {
	l := c.Query("lang", i.fb)
	a := c.Get("Accept-Language", i.fb)
	return l, a
}

func (i *I18n) I18nMiddleware(c *fiber.Ctx) error {
	l, a := i.GetLanguagesInContext(c)
	c.Locals("lang", l)
	c.Locals("accept-language", a)
	return c.Next()
}
