## Ssi Core / Cipher

This package has been developed for password encrypt/decrypt operations in microservices.

### Encrypt

Encrypts the given password with the given salt.

example:

```go
package main

import(
	"github.com/ssibrahimbas/ssi-core/pkg/cipher"
)

func main() {
    c := cipher.New()
    encrypted, err := c.Encrypt("test")
    if err != nil {
        panic(err)
    }
    fmt.Println(encrypted)
}
```

### Compare

Compares the given password with the given encrypted password.

example:

```go
package main

import(
	"github.com/ssibrahimbas/ssi-core/pkg/cipher"
)

func main() {
    c := cipher.New()
    encrypted, err := c.Encrypt("test")
    if err != nil {
        panic(err)
    }
    equal, err := c.Compare("test1", encrypted)
    if err != nil {
        panic(err)
    }
    fmt.Println(equal)
}
```