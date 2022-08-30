package cipher

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCipher_Module(t *testing.T) {
	c := New()
	t.Run("Encrypt", func(t *testing.T) {
		t.Run("should return encrypted text", func(t *testing.T) {
			t.Parallel()
			text := "test"
			encrypted, err := c.Encrypt(text)
			assert.NoError(t, err)
			assert.NotEqual(t, len(encrypted), 0)
		})
	})
	t.Run("Compare", func(t *testing.T) {
		t.Run("should return true if text and hash are equal", func(t *testing.T) {
			t.Parallel()
			text := "test"
			encrypted, err := c.Encrypt(text)
			assert.NoError(t, err)
			assert.NotEqual(t, len(encrypted), 0)
			equal, err := c.Compare(text, encrypted)
			assert.NoError(t, err)
			assert.Equal(t, equal, true)
		})
		t.Run("should return false if text and hash are not equal", func(t *testing.T) {
			t.Parallel()
			text := "test"
			encrypted, err := c.Encrypt(text)
			assert.NoError(t, err)
			assert.NotEqual(t, len(encrypted), 0)
			equal, err := c.Compare("test1", encrypted)
			assert.Error(t, err)
			assert.Equal(t, equal, false)
		})
	})
}
