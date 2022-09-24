## Ssi Core / Nats

This package includes the implementation of [JetStream by nats](https://docs.nats.io/nats-concepts/jetstream) for microservices to communicate with each other.

### Types

#### connection

```go
type Conn struct {
	C  *nats.Conn
	JS nats.JetStreamContext
}
```

Conn is a struct that includes the connection and the JetStream context.

#### functions

```go
func (c *Conn) Close()
func (c *Conn) Publish(subject string, data []byte) (*nats.PubAck, error)
func (c *Conn) Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error)
func (c *Conn) AddStream(name string) (*nats.StreamInfo, error)
func (c *Conn) UpdateStream(name string) (*nats.StreamInfo, error)
```

### Create a Connection

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/nats"
    "fmt"
)

func main() {
    n := nats.New()
    fmt.Println(n.C.ConnectedUrl())
}
```

### Close a Connection

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/nats"
    "fmt"
)

func main() {
    n := nats.New()
    n.Close()
    fmt.Println(n.C.ConnectedUrl())
}
```

### Publish a Message

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/nats"
    "fmt"
)

func main() {
    n := nats.New()
    _, err := n.Publish("test", []byte("Hello, World!"))
    if err != nil {
        fmt.Println(err)
    }
}
```

### Subscribe a Message

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/nats"
	natsGo "github.com/nats-io/nats.go"
    "fmt"
)

func main() {
    n := nats.New()
    _, err := n.Subscribe("test", func(msg *natsGo.Msg) {
        fmt.Println(string(msg.Data))
    })
    if err != nil {
        fmt.Println(err)
    }
}
```

### Add a Stream

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/nats"
    "fmt"
)

func main() {
    n := nats.New()
    _, err := n.AddStream("test")
    if err != nil {
        fmt.Println(err)
    }
}
```

### Update a Stream

```go
package main

import (
    "github.com/ssibrahimbas/ssi-core/pkg/nats"
    "fmt"
)

func main() {
    n := nats.New()
    _, err := n.UpdateStream("test")
    if err != nil {
        fmt.Println(err)
    }
}
```