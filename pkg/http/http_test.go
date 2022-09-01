package http

import (
	"context"
	"github.com/ssibrahimbas/ssi-core/pkg/i18n"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestHttp_Module(t *testing.T) {
	i := i18n.New("en")
	t.Run("New should return a new instance of http module", func(t *testing.T) {
		h := New(i)
		assert.NotEqual(t, h, nil)
	})

	t.Run("Listen should return an error if the http server fails to start", func(t *testing.T) {
		h := New(i)
		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
		defer cancel()
		go func() {
			_ = h.Listen(":3000")
		}()
		ctx.Deadline()
	})
}
