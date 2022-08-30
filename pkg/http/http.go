package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Client struct {
	App *fiber.App
}

func New() *Client {
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
	})
	app.Use(recover.New())
	return &Client{
		App: app,
	}
}

func (h *Client) Listen(p string) error {
	return h.App.Listen(p)
}
