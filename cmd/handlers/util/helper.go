package util

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"strings"
)

func Message(status bool, message string) map[string]interface{} {
	return map[string]interface{}{"status": status, "message": message}
}

func GetUsername(r *http.Request) (string, error) {
	token := strings.Split(r.Header.Get("Authorization"), " ")[1]

	claims, err := ExtractClaims(token)

	if err != nil {
		return "", err
	}

	username := fmt.Sprintf("%v", claims["name"])

	return username, nil
}

func GetFromQuery(r *http.Request, key string) (string, error) {

	queries := r.URL.Query()

	if statusStr, ok := queries[key]; ok {
		if len(statusStr) > 0 {
			return statusStr[0], nil
		}

		return "", errors.New(fmt.Sprintf("empty %q parameter", key))
	}

	return "", errors.New(fmt.Sprintf("param %q doesn't exist", key))

}

func GetId(r *http.Request) string {
	return mux.Vars(r)["id"]
}

func GetStatusFromQuery(r *http.Request) (bool, error) {
	res, err := GetFromQuery(r, "status")

	if err != nil {
		return false, err
	}

	return strconv.ParseBool(res)
}
