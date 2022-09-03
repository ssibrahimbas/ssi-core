package hook

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestOnClose(t *testing.T) {
	t.Run("should create new OnClose", func(t *testing.T) {
		o := NewOnClose()
		if o == nil {
			t.Fatal("OnClose is nil")
		}
	})

	t.Run("AddHandler should add handler", func(t *testing.T) {
		o := NewOnClose()
		o.AddHandler(func(sig os.Signal) {
			fmt.Println("signal received")
		})
		if len(o.handlers) != 1 {
			t.Fatal("handler is not added")
		}
	})

	t.Run("Listen should listen to signal", func(t *testing.T) {
		o := NewOnClose()
		o.AddHandler(func(sig os.Signal) {
			fmt.Println("signal received")
		})
		go func() {
			time.Sleep(10 * time.Millisecond)
			p, err := os.FindProcess(os.Getpid())
			if err != nil {
				panic(err.Error())
			}
			err = p.Signal(os.Interrupt)
			if err != nil {
				panic(err.Error())
			}
		}()

		o.Listen(os.Interrupt)
	})
}
