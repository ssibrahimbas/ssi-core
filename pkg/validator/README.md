## Ssi Core / Validator

It has been developed to provide validation infrastructure related to i18n to our microservices. Custom also includes some tags and the microservice we use amy not load it if it wants to.

### Type

```go
type ErrorResponse struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}
```

### Functions

```go
func New(i *i18n.I18n) *Validator

func (v *Validator) ValidateStruct(s interface{}, languages ...string) []*ErrorResponse
func (v *Validator) ConnectCustom()
func (v *Validator) RegisterTagName()
```

### Example

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/i18n"
    "github.com/ssibrahimbas/ssi-core/pkg/validator"
    "fmt"
)

type User struct {
    Name string `validate:"required"`
    Age  int    `validate:"required"`
}

func main() {
    i := i18n.New("tr")
    i.LoadLanguages("./locales", "tr", "en")
    v := validator.New(i)

    u := &User{
        Name: "Sami",
        Age:  20,
    }
    	if errors := v.ValidateStruct(u, "en", "tr"); len(errors) > 0 {
		fmt.Println("An error occurred while validating the user.")
	}
}
```

### Load Custom Validators

Custom validators: [`username`, `password`, `locale`, `object_id`, `slug`]

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/i18n"
    "github.com/ssibrahimbas/ssi-core/pkg/validator"
    "fmt"
)

type User struct {
    Name string `validate:"required"`
    Age  int    `validate:"required"`
}

func main() {
    i := i18n.New("tr")
    i.LoadLanguages("./locales", "tr", "en")
    v := validator.New(i)
    v.ConnectCustom()
    v.RegisterTagName()
}
```