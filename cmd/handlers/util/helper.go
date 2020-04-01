package util

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func GenerateTokenPair(username string) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	signedStrBytes := []byte(os.Getenv("token_password"))

	c := token.Claims.(jwt.MapClaims)
	c["name"] = username
	c["exp"] = time.Now().Add(20 * time.Minute).Unix()

	t, err := token.SignedString(signedStrBytes)

	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)

	rtClaims := refreshToken.Claims.(jwt.MapClaims)

	rtClaims["name"] = username
	rtClaims["exp"] = time.Now().Add(24 * time.Hour).Unix()

	rt, err := refreshToken.SignedString(signedStrBytes)

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil

}

func GetUsername(r *http.Request) string {
	return mux.Vars(r)["username"]
}
