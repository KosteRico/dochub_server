package bookmarks

import (
	"checkaem_server/cmd/database/bookmark"
	"checkaem_server/cmd/handlers/util"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var Modify = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

	postId := mux.Vars(r)["id"]

	status, err := util.GetStatusFromQuery(r)

	if err != nil {
		log.Println(err.Error())
		util.RespondInvalidRequest(w)
		return
	}

	if status {
		err = bookmark.Insert(username, postId)
	} else {
		err = bookmark.Delete(username, postId)
	}

	if err != nil {
		log.Println(err.Error())
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "Bookmark were successfully changed")
	util.Respond(w, resp)
})

var GetPosts = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

	posts, err := bookmark.GetPosts(username)

	if err != nil {
		log.Println(err.Error())
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "Posts was found")

	resp["posts"] = posts

	util.Respond(w, resp)
})
