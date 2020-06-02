package util

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func GetSignedStrBytes() []byte {
	return []byte(os.Getenv("token_password"))
}

func GenerateTokenPair(username string) (map[string]string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	signedStrBytes := GetSignedStrBytes()

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
	rtClaims["exp"] = time.Now().Add(48 * time.Hour).Unix()

	rt, err := refreshToken.SignedString(signedStrBytes)

	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil

}

func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	signedStrBytes := GetSignedStrBytes()

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// check token signing method etc
		return signedStrBytes, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid JWT token")
}
