package i18n

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestI18n_ModuleTesting(t *testing.T) {
	t.Run("Should return a new I18n instance", func(t *testing.T) {
		i := New("en")
		assert.NotNil(t, i)
	})

	t.Run("Should load languages", func(t *testing.T) {
		i := New("en")
		i.LoadLanguages("./locales_test", "en", "tr")
		assert.Equal(t, 2, len(i.b.LanguageTags()))
	})

	t.Run("Should translate", func(t *testing.T) {
		i := New("en")
		i.LoadLanguages("./locales_test", "en", "tr")
		assert.Equal(t, "Hello from Test.", i.Translate("test"))
	})

	t.Run("Should translate with params", func(t *testing.T) {
		i := New("en")
		i.LoadLanguages("./locales_test", "en", "tr")
		assert.Equal(t, "Hello Test.", i.TranslateWithParams("greeting", map[string]string{"Name": "Test"}))
	})
}
