package auth

import (
	userDb "checkaem_server/cmd/database/user"
	"checkaem_server/cmd/entities/user"
	"checkaem_server/cmd/handlers/util"
	"encoding/json"
	"net/http"
)

var CreateAccount = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	u := user.NewEmpty()

	err := json.NewDecoder(r.Body).Decode(u)

	if err != nil {
		util.RespondInvalidRequest(w)
		return
	}

	resp, err := u.Create()

	if err != nil {
		util.RespondInternalServerError(w, "Creation error occurred")
		return
	}

	err = userDb.Insert(u)

	if err != nil {
		util.RespondInternalServerError(w, "Database error occurred")
		return
	}

	util.Respond(w, resp)
})

var Login = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	u := user.NewEmpty()

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		util.RespondInvalidRequest(w)
		return
	}

	resp := login(u.Username, u.Password)
	util.Respond(w, resp)
})

var RefreshToken = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	username := util.GetUsername(r)

	_, err := userDb.Get(username)

	if err != nil {
		util.RespondUserNotFound(w, username)
		return
	}

	respStr, err := util.GenerateTokenPair(username)

	resp := util.Message(true, "Access and refresh tokens were successfully generated")

	resp["token"] = respStr

	util.Respond(w, resp)
})
