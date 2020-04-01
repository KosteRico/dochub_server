package middleware

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

var jwtMid = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("token_password")), nil
	},
})

func IsAuthenticated(h http.Handler) http.Handler {
	return jwtMid.Handler(h)
}
