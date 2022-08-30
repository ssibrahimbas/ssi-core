package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/ssi-core/pkg/result"
)

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*result.Result); ok {
		return c.Status(e.Code).JSON(e)
	}
	if e, ok := err.(*result.DataResult); ok {
		return c.Status(e.Code).JSON(e)
	}
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	err = c.Status(code).JSON(result.Error(err.Error(), code))
	return err
}
