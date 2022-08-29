package nats

import nats "github.com/nats-io/nats.go"

func New() *Conn {
	nc, js := initDefault()
	return &Conn{
		C:  nc,
		JS: js,
	}
}

func (c *Conn) Close() {
	c.C.Close()
}

func (c *Conn) Publish(subject string, data []byte) (*nats.PubAck, error) {
	return c.JS.Publish(subject, data)
}

func (c *Conn) Subscribe(subject string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return c.JS.Subscribe(subject, cb)
}

func (c *Conn) AddStream(name string) (*nats.StreamInfo, error) {
	return c.JS.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: []string{name + ".*"},
	})
}

func (c *Conn) UpdateStream(name string) (*nats.StreamInfo, error) {
	return c.JS.UpdateStream(&nats.StreamConfig{
		Name:     name,
		MaxBytes: 8,
	})
}
