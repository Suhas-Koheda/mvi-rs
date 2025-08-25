package oauth

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/suhas-koheda/mvi-rs/pkg/env"
)

var (
	key []byte
	t   *jwt.Token
	s   string
)

func init() {
	key = []byte(loadenv.HashSecret)
	t = jwt.New(jwt.SigningMethodES256)
	s, _ = t.SignedString(key)
}
