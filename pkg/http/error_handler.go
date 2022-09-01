package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/ssi-core/pkg/result"
	"go.mongodb.org/mongo-driver/mongo"
)

func errorHandler(c *fiber.Ctx, err error) error {
	code := fiber.StatusInternalServerError
	if e, ok := err.(*result.Result); ok {
		return c.Status(e.Code).JSON(e)
	}
	if e, ok := err.(*result.DataResult); ok {
		return c.Status(e.Code).JSON(e)
	}
	if err == mongo.ErrNoDocuments {
		return c.Status(fiber.StatusNotFound).JSON(result.Error("not_found", fiber.StatusNotFound))
	}
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}
	err = c.Status(code).JSON(result.Error(err.Error(), code))
	return err
}
