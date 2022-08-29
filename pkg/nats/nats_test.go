package nats

import (
	nats "github.com/nats-io/nats.go"
	"testing"
)

func TestConnectNats(t *testing.T) {
	n := New()
	if !n.C.IsConnected() {
		t.Fatal("nats is not connected")
	}
	n.Close()
}

func TestConnectNatsShouldNats(t *testing.T) {
	n := connectNats()
	if !n.IsConnected() {
		t.Fatal("nats is not connected")
	}
	n.Close()
}

func TestConnectJSShouldJetStream(t *testing.T) {
	n := connectNats()
	js := connectJS(n)
	if js == nil {
		t.Fatal("jetStream is not connected")
	}
	n.Close()
}

func TestInitDefaultShouldNatsAndJetStream(t *testing.T) {
	nc, js := initDefault()
	if !nc.IsConnected() {
		t.Fatal("nats is not connected")
	}
	if js == nil {
		t.Fatal("jetStream is not connected")
	}
	nc.Close()
}

func TestConn_AddStream(t *testing.T) {
	n := New()
	_, err := n.AddStream("TEST")
	if err != nil {
		t.Fatal(err)
	}
}

func TestConn_UpdateStream(t *testing.T) {
	n := New()
	n.AddStream("test")
	defer n.Close()
	_, err := n.UpdateStream("test")
	if err != nil {
		t.Fatal(err)
	}
}

func TestConn_Publish(t *testing.T) {
	n := New()
	n.AddStream("TEST")
	defer n.Close()
	_, err := n.Publish("TEST.1", []byte("any"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestConn_Subscribe(t *testing.T) {
	n := New()
	n.AddStream("Test")
	defer n.Close()
	_, err := n.Subscribe("Test.1", func(msg *nats.Msg) {
		t.Log(msg)
	})
	if err != nil {
		t.Fatal(err)
	}
	_, err = n.Publish("Test.1", []byte("any"))
	if err != nil {
		t.Fatal(err)
	}
}
