package nats

import nats "github.com/nats-io/nats.go"

type Conn struct {
	C  *nats.Conn
	JS nats.JetStreamContext
}
