## Ssi Core / JWT

This package provides a simple JWT implementation.

### Usage

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/jwt"
    "github.com/ssibrahimbas/ssi-core/pkg/auth"
    "github.com/ssibrahimbas/ssi-core/pkg/helper"
    "fmt"
)

func main() {
    // Create a new JWT instance with the secret key
    j := jwt.New("secret")
    u := &auth.CurrentUser{
        Id: "123",
        Email: "info@ssibrahimbas.com"
    }
    t, err := j.Sign(u)
    helper.CheckErr(err)
    fmt.Println(t)
}
```