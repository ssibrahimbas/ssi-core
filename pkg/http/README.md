## Ssi Core / Http

This package uses and customizes  the fiber application to be used in microservice applications. It was developed to not do them individually in every microservice.

### Type

```go
type Client struct {
	App *fiber.App
}
```

### Config 

```go
type Config struct {
	NFMsgKey string
	DfMsgKey string
}
```

### Functions

```go
func New(i18n *i18n.I18n, config ...Config) *Client
func (h *Client) Listen(p string) error
```

### Example

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/ssibrahimbas/ssi-core/pkg/http"
    "github.com/ssibrahimbas/ssi-core/pkg/i18n"
    "log"
)

func main() {
    i := i18n.New("tr")
	i.LoadLanguages("./locales", "tr", "en")
    h := http.New(i)
    h.App.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("Hello, World!")
    })
    log.Fatal(h.Listen(":3000"))
}
```