package jwt

import (
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJwt_Module(t *testing.T) {
	t.Run("Sign", func(t *testing.T) {
		j := New("secret")
		_, err := j.Sign("test")
		assert.NoError(t, err)
	})
	t.Run("Parse", func(t *testing.T) {
		j := New("secret")
		token, err := j.Sign("test")
		assert.NoError(t, err)
		res, err := j.Parse(token)
		assert.NoError(t, err)
		payload := res.Claims.(jwt.MapClaims)["payload"]
		assert.Equal(t, "test", payload)
	})
	t.Run("Parse function testing with RSA method", func(t *testing.T) {
		j := New("secret")
		token := "eyJhbGciOiJSUzM4NCIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.o1hC1xYbJolSyh0-bOY230w22zEQSk5TiBfc-OCvtpI2JtYlW-23-8B48NpATozzMHn0j3rE0xVUldxShzy0xeJ7vYAccVXu2Gs9rnTVqouc-UZu_wJHkZiKBL67j8_61L6SXswzPAQu4kVDwAefGf5hyYBUM-80vYZwWPEpLI8K4yCBsF6I9N1yQaZAJmkMp_Iw371Menae4Mp4JusvBJS-s6LrmG2QbiZaFaxVJiW8KlUkWyUCns8-qFl5OMeYlgGFsyvvSHvXCzQrsEXqyCdS4tQJd73ayYA4SPtCb9clz76N1zE5WsV4Z0BYrxeb77oA7jJhh994RAPzCG0hmQ"
		_, err := j.Parse(token)
		assert.Error(t, err)
		assert.Equal(t, err.Error(), "unexpected signing method: RS384")
	})
}
