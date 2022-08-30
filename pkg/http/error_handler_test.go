package http

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/ssibrahimbas/ssi-core/pkg/result"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
	"time"
)

func TestHttp_ErrorHandler(t *testing.T) {
	h := New()
	assert.NotEqual(t, h, nil)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(time.Millisecond*80))
	defer cancel()

	h.App.Get("/return-result", func(c *fiber.Ctx) error {
		return result.Error("Something went wrong", fiber.StatusBadRequest)
	})

	h.App.Get("/return-data-result", func(c *fiber.Ctx) error {
		return result.SuccessData("Operation Success", map[string]string{"env": "test"}, fiber.StatusOK)
	})

	h.App.Get("/return-error", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusBadRequest, "Something went wrong")
	})

	h.App.Get("/any-error", func(c *fiber.Ctx) error {
		panic("Something went wrong")
		return nil
	})

	go func() {
		_ = h.Listen(":3050")
	}()

	t.Run("errorHandler should return an error if the http server fails to start", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/return-result", nil)
		assert.NoError(t, err)
		res, err := h.App.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	t.Run("errorHandler should return an error if the http server fails to start", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/return-data-result", nil)
		assert.NoError(t, err)
		res, err := h.App.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusOK, res.StatusCode)
	})

	t.Run("errorHandler should return an error if the http server fails to start", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/return-error", nil)
		assert.NoError(t, err)
		res, err := h.App.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusBadRequest, res.StatusCode)
	})

	t.Run("errorHandler should return an error if the http server fails to start", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/any-error", nil)
		assert.NoError(t, err)
		res, err := h.App.Test(req)
		assert.NoError(t, err)
		assert.Equal(t, fiber.StatusInternalServerError, res.StatusCode)
	})

	ctx.Deadline()
}
