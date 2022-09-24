## Ssi Core

It is a microservice core package developed with Golang.

It is a repository that treats repetitive code as a basis in an already developed microservice application. It has been made public as an example. You can report the missing aspects to me as an `issue` or send a `PR`.

### Technologies

- [Golang](https://go.dev/) `for the language`
- [Fiber](https://docs.gofiber.io/) `for the http`
- [JetStream by Nats](https://docs.nats.io/nats-concepts/jetstream) `for the message broker`
- [Mongodb](https://www.mongodb.com/docs/) `for the database`
- [JWT](https://jwt.io/) `for the authentication`
- [Go Playground Validator](https://github.com/go-playground/validator) `for the validation`
- [Go I18n v2](https://github.com/nicksnyder/go-i18n) `for the localization`

### Packages

- [auth](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/auth)
- [cipher](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/cipher)
- [config](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/config)
- [db](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/db)
- [helper](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/helper)
- [hook](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/hook)
- [http](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/http)
- [i18n](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/i18n)
- [jwt](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/jwt)
- [nats](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/nats)
- [result](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/result)
- [validator](https://github.com/ssibrahimbas/ssi-core/tree/main/pkg/validator)

### Installation

```bash
go get github.com/ssibrahimbas.com/ssi-core
```

### Usage

```go
package main

import(
	"github.com/ssibrahimbas/ssi-core/pkg/http"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"log"
)

func main() {
    i := i18n.New("tr")
	i.LoadLanguages("./locales", "tr", "en")
	h := http.New(i)
    h.Post("/login", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{"message": "Hello, World!"})
    })
	log.Fatal(a.Http.Listen(":8080"))
}
```

### License

[MIT](https://choosealicense.com/licenses/mit/)

### Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

### Authors

- [Sami Salih İbrahimbaş](https://github.com/ssibrahimbas)