package middleware

import (
	"checkaem_server/cmd/handlers/util"
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"net/http"
)

var jwtMid = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return util.GetSignedStrBytes(), nil
	},
})

func IsAuthenticated(h http.Handler) http.Handler {
	return jwtMid.Handler(h)
}
