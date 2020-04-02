package post

import (
	dbPost "checkaem_server/cmd/database/post"
	"checkaem_server/cmd/entities/post"
	"checkaem_server/cmd/handlers/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx"
	"net/http"
)

var GetCreated = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username := util.GetUsername(r)

	posts, err := dbPost.GetByCreator(username)

	if err != nil {
		if err == pgx.ErrNoRows {
			util.RespondUserNotFound(w, username)
		}
		return
	}

	posts = post.Mergesort(posts, func(a, b *post.Post) bool {
		return a.DateCreated.Time.Unix() >= b.DateCreated.Time.Unix()
	})

	resp := util.Message(true, fmt.Sprintf(`All created posts of "%s"`, username))
	resp["posts"] = posts

	util.Respond(w, resp)
})

var Create = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	p := post.NewEmpty()

	err := json.NewDecoder(r.Body).Decode(&p)

	p.CreatorUsername = util.GetUsername(r)

	if err != nil {
		util.RespondInvalidRequest(w)
		return
	}

	p, err = dbPost.Insert(p)

	if err != nil {
		util.RespondInternalServerError(w, "Database insertion error")
		return
	}

	resp := util.Message(true, "Post was successfully added")

	resp["post"] = p

	util.Respond(w, resp)
})

var Delete = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	postId := mux.Vars(r)["id"]
	username := util.GetUsername(r)

	if postId == "" {
		w.WriteHeader(http.StatusBadRequest)
		util.RespondMessage(w, false, "Post id is undefined")
		return
	}

	p, err := dbPost.Delete(postId, username)

	if err != nil {
		if err == pgx.ErrNoRows {
			util.RespondNotFound(w, fmt.Sprintf("Post with uuid %q not found", postId))
		} else if err == dbPost.ErrNotAdmin {
			util.RespondNotAdmin(w)
		}
		return
	}

	resp := util.Message(true, "Post was successfully deleted")

	resp["post"] = p
	util.Respond(w, resp)
})

var Modify = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	p := post.NewEmpty()

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		util.RespondInvalidRequest(w)
		return
	}

	p.Id = mux.Vars(r)["id"]
	p.CreatorUsername = util.GetUsername(r)

	p, err = dbPost.Modify(p)

	if err != nil {
		if err == dbPost.ErrNotAdmin {
			util.RespondNotAdmin(w)
		} else {
			util.RespondInternalServerError(w, "Database update error")
		}
		return
	}

	resp := util.Message(true, "Your post was successfully modified")
	resp["post"] = p

	util.Respond(w, resp)
})
