## Ssi Core / Auth

This package has been developed to perform independent authorization in microservices.

You can find many authentication strategies in microservice applications. We chose this strategy because our microservices are truly independent from each other and it is the most suitable choice for the project we are using. If a different strategy is more suitable for your project, you can skip it.

## Middlewares

### CurrentUserMiddleware

If there is a token in the request's cookie, it decodes it and adds the current user to the fiber context. If there is no token, it continues. If you want to enforce authorization then use `requiredAuth` after this middleware.

params:

```go
type CurrentUserConfig struct {
	Jwt    *jwt.Jwt
	I18n   *i18n.I18n
	MsgKey string
}
```

example:

```go
// handler.go
package internal

import(
    "github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/ssi-core/pkg/auth"
	"github.com/ssibrahimbas/ssi-core/pkg/http"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/jwt"
)

type Handler struct {
	s    *Service
	h    *http.Client
	i18n *i18n.I18n
	jwt  *jwt.Jwt
}

func (h *Handler) InitAllVersions() {
	h.initV1()
}

func (h *Handler) initV1() {
	v1 := h.h.App.Group("api/auth/v1")
	v1.Use(h.i18n.I18nMiddleware)
	v1.Post("/login", h.Login)
	v1.Post("/register", h.Register)
	v1.Post("/logout", h.currentUser(), h.requiredAuth(), h.Logout)
	v1.Post("/refresh", h.currentUser(), h.requiredAuth(), h.ExtendToken)
}

func (h *Handler) currentUser() fiber.Handler {
	return auth.NewCurrentUser(&auth.CurrentUserConfig{
		I18n:   h.i18n,
		Jwt:    h.jwt,
		MsgKey: "auth_unauthorized",
	})
}

func (h *Handler) requiredAuth() fiber.Handler {
	return auth.NewRequiredAuth(&auth.RequiredAuthConfig{
		I18n:   h.i18n,
		MsgKey: "auth_unauthenticated",
	})
}
```

### RequiredAuthMiddleware

This middleware is used to check the authorization of the user. It is used in the routes that require authorization.

Important Note: It should be used after `CurrentUserMiddleware`.

params:

```go
type RequiredAuthConfig struct {
	I18n   *i18n.I18n
	MsgKey string
}
```

example:

```go
// handler.go
package internal

import(
    "github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/ssi-core/pkg/auth"
	"github.com/ssibrahimbas/ssi-core/pkg/http"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/ssibrahimbas/ssi-core/pkg/jwt"
)

type Handler struct {
	s    *Service
	h    *http.Client
	i18n *i18n.I18n
	jwt  *jwt.Jwt
}

func (h *Handler) InitAllVersions() {
	h.initV1()
}

func (h *Handler) initV1() {
	v1 := h.h.App.Group("api/auth/v1")
	v1.Use(h.i18n.I18nMiddleware)
	v1.Post("/login", h.Login)
	v1.Post("/register", h.Register)
	v1.Post("/logout", h.currentUser(), h.requiredAuth(), h.Logout)
	v1.Post("/refresh", h.currentUser(), h.requiredAuth(), h.ExtendToken)
}

func (h *Handler) currentUser() fiber.Handler {
	return auth.NewCurrentUser(&auth.CurrentUserConfig{
		I18n:   h.i18n,
		Jwt:    h.jwt,
		MsgKey: "auth_unauthorized",
	})
}

func (h *Handler) requiredAuth() fiber.Handler {
	return auth.NewRequiredAuth(&auth.RequiredAuthConfig{
		I18n:   h.i18n,
		MsgKey: "auth_unauthenticated",
	})
}
```

### Access User Data in Handler

You can access the user data in the handler by using the `ParseCurrentUser` function. this function returns `CurrentUser` back.

type:

```go
type CurrentUser struct {
	ID    string `json:"uuid"`
	Email string `json:"email"`
}
```

example:

```go
// api.go
package internal

import(
    "github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/ssi-core/pkg/auth"
)

func (h *Handler) ExtendToken(c *fiber.Ctx) error {
    u := auth.ParseCurrentUser(c) // typeof u is *CurrentUser
    // do something with user

    return c.JSON(fiber.Map{
        "message": "token extended",
        "email": u.Email,
        "id": u.Id,
    })
}
```
