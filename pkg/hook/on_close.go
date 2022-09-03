package hook

import (
	"os"
	"os/signal"
	"sync"
)

type OnCloseHandler func(sig os.Signal)

type OnClose struct {
	handlers []OnCloseHandler
}

func NewOnClose() *OnClose {
	return &OnClose{
		handlers: []OnCloseHandler{},
	}
}

func (o *OnClose) AddHandler(h OnCloseHandler) {
	o.handlers = append(o.handlers, h)
}

func (o *OnClose) Listen(signals ...os.Signal) {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, signals...)
	sig := <-ch
	var wg sync.WaitGroup
	for _, h := range o.handlers {
		wg.Add(1)
		go func(h OnCloseHandler, s os.Signal) {
			defer wg.Done()
			h(s)
		}(h, sig)
	}
	wg.Wait()
}
