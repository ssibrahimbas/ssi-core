## Ssi Core / Config

This package allows you to use the values from the `.env` file by collation of the env values of the machine.

abstract:

```go
LoadConfig(p string, c interface{}) interface{}
```

example:

```go
package main

import(
    "github.com/ssibrahimbas/ssi-core/pkg/config"
)

func main() {
    type Config struct {
        Port string `env:"PORT"`
    }
    c := &Config{}
    config.LoadConfig(".", &c)
    fmt.Println(c.Port)
}
```