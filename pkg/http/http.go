package http

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type Client struct {
	App *fiber.App
}

type Config struct {
	NFMsgKey string
	DfMsgKey string
}

var defaultConfig = Config{
	NFMsgKey: "not_found",
	DfMsgKey: "",
}

func New(config ...Config) *Client {
	var cfg = defaultConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler(&cfg),
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
