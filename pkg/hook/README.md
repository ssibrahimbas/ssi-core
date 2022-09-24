
## Ssi Core / Hook

This package contains some hooks that listen for events sent by the machine.

### OnClose

This hook is called when the program is closed.

type:

```go
func NewOnClose() *OnClose
func (o *OnClose) AddHandler(h OnCloseHandler)
func (o *OnClose) Listen(signals ...os.Signal)
```

example:

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/hook"
    "fmt"
    "os"
)

func main() {
    o := hook.NewOnClose()
    o.AddHandler(func(sig os.Signal) {
        fmt.Println("closing")
    })
    o.Listen(os.Interrupt)
}
```