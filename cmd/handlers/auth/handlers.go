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
		util.RespondInternalServerError(w, err)
		return
	}

	err = userDb.Insert(u)

	if err != nil {
		util.RespondInternalServerError(w, err)
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

	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

	_, err = userDb.Get(username)

	if err != nil {
		util.RespondUserNotFound(w, username)
		return
	}

	respStr, err := util.GenerateTokenPair(username)

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "Access and refresh tokens were successfully generated")

	resp["token"] = respStr

	util.Respond(w, resp)
})

var GetAllUsernames = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	usernames, err := userDb.GetAllNames()

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "Usernames was successfully retrieved")
	resp["usernames"] = usernames

	util.Respond(w, resp)
})

var CheckUsernameExists = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username, err := util.GetFromQuery(r, "username")

	if err != nil {
		util.RespondInvalidRequest(w)
		return
	}

	exists, err := userDb.Exists(username)

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := make(map[string]interface{})

	resp["exists"] = exists

	util.Respond(w, resp)
})
