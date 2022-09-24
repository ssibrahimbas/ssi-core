## Ssi Core / I18n

This package has been developed to localize the response messages in the microservice applications.

Works with `.toml`.

### Create New I18n

type:

```go
type I18n interface {
	LoadLanguages(localeDir string, languages ...string)
	Translate(key string, languages ...string) string
    TranslateWithParams(key string, params interface{}, languages ...string) string
    GetLanguagesInContext(c *fiber.Ctx) (string, string)
    I18nMiddleware(c *fiber.Ctx) error
}

func New(fb string) *I18n
```

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/i18n"
)

func main() {
    i := i18n.New("tr")
}
```

### Load Languages

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/i18n"
)

func main() {
    i := i18n.New("tr")
    i.LoadLanguages("./locales", "tr", "en")
    // ./locales/tr.toml
    // ./locales/en.toml
}
```

### Translate

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/i18n"
    "fmt"
)

func main() {
    i := i18n.New("tr")
    i.LoadLanguages("./locales", "tr", "en")
    fb := i.Translate("hello_world") // translate with fallback language
    fmt.Println(fb) // Merhaba, Dünya!
    
    tr := i.Translate("hello_world", "tr") // translate with tr language
    fmt.Println(tr) // Merhaba, Dünya!

    en := i.Translate("hello_world", "en") // translate with en language
    fmt.Println(en) // Hello, World!
}
```

### Translate With Params

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/i18n"
    "fmt"
)

func main() {
    i := i18n.New("tr")
    i.LoadLanguages("./locales", "tr", "en")
    fb := i.TranslateWithParams("hello_world", map[string]interface{}{"name": "Sami"}) // translate with fallback language
    fmt.Println(fb) // Merhaba, Sami!
    
    tr := i.TranslateWithParams("hello_world", map[string]interface{}{"name": "Sami"}, "tr") // translate with tr language
    fmt.Println(tr) // Merhaba, Sami!

    en := i.TranslateWithParams("hello_world", map[string]interface{}{"name": "Sami"}, "en") // translate with en language
    fmt.Println(en) // Hello, Sami!
}
```

### Get Languages In Context

```go
package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/ssibrahimbas/ssi-core/pkg/http"
    "github.com/ssibrahimbas/ssi-core/pkg/i18n"
    "fmt"
)

func main() {
    i := i18n.New("tr")
    i.LoadLanguages("./locales", "tr", "en")
    h := http.New(i)
	h.App.Use(h.i18n.I18nMiddleware)
    h.App.Use(i.I18nMiddleware)
    h.App.Get("/", func(c *fiber.Ctx) error {
        fb, _ := h.i18n.GetLanguagesInContext(c) // get languages in context with fallback language
        fmt.Println(fb) // tr
        return c.SendString("Hello, World!")
    })
    h.Listen(":3000")
}
```

### I18n Middleware

use i18n in fiber

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
	h.App.Use(h.i18n.I18nMiddleware)
    h.App.Get("/", func(c *fiber.Ctx) error {
        l, a := h.i18n.GetLanguagesInContext(c)
        return h.i18n.Translate("hello_world", l, a)
    })
    log.Fatal(h.Listen(":3000"))
}
```