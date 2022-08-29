package nats

import (
	nats "github.com/nats-io/nats.go"
)

func connectNats() *nats.Conn {
	nc, _ := nats.Connect(nats.DefaultURL)
	return nc
}

func connectJS(nc *nats.Conn) nats.JetStreamContext {
	js, _ := nc.JetStream()
	return js
}

func initDefault() (*nats.Conn, nats.JetStreamContext) {
	nc := connectNats()
	js := connectJS(nc)
	return nc, js
}
