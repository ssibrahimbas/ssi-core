package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type Jwt struct {
	sk []byte
}

func New(s string) *Jwt {
	return &Jwt{
		sk: []byte(s),
	}
}

func (j *Jwt) Sign(p interface{}) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorize"] = true
	claims["payload"] = p
	claims["expire"] = time.Now().Add(time.Minute * 30).Unix()
	return token.SignedString(j.sk)
}

func (j *Jwt) Parse(t string) (*jwt.Token, error) {
	return jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.sk, nil
	})
}
