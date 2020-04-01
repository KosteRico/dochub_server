package auth

import (
	"checkaem_server/cmd/database/user"
	"checkaem_server/cmd/handlers/util"
	"github.com/jackc/pgx"
)

func login(username, password string) map[string]interface{} {
	u, err := user.Get(username)

	if err != nil {
		if err == pgx.ErrNoRows {
			return util.Message(false, "Username not found")
		}
		return util.Message(false, "Connection error. Please, retry later")
	}

	if !u.ComparePassword(password) {
		return util.Message(false, "Invalid password")
	}

	u.Password = ""

	tokenString, err := util.GenerateTokenPair(username)

	if err != nil {
		return util.Message(false, "JWT error occurred")
	}

	u.Token = tokenString

	resp := util.Message(true, "Logged In")
	resp["account"] = u

	return resp
}
