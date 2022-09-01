package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/ssi-core/pkg/result"
	"go.mongodb.org/mongo-driver/mongo"
)

func errorHandler(cfg *Config) func(c *fiber.Ctx, err error) error {
	return func(c *fiber.Ctx, err error) error {
		code := fiber.StatusInternalServerError
		if e, ok := err.(*result.Result); ok {
			return c.Status(e.Code).JSON(e)
		}
		if e, ok := err.(*result.DataResult); ok {
			return c.Status(e.Code).JSON(e)
		}
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).JSON(result.Error(cfg.NFMsgKey, fiber.StatusNotFound))
		}
		if e, ok := err.(*fiber.Error); ok {
			code = e.Code
		}
		if cfg.DfMsgKey != "" {
			return c.Status(code).JSON(result.Error(cfg.DfMsgKey, code))
		}
		err = c.Status(code).JSON(result.Error(err.Error(), code))
		return err
	}
}
