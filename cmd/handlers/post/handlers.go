package post

import (
	dbPost "checkaem_server/cmd/database/post"
	"checkaem_server/cmd/entities/post"
	"checkaem_server/cmd/handlers/util"
	"fmt"
	"net/http"
)

var GetCreatedPosts = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username := util.GetUsername(r)

	posts, err := dbPost.GetByCreator(username)

	if err != nil {
		util.RespondUserNotFound(w, username)
		return
	}

	posts = post.Mergesort(posts, func(a, b *post.Post) bool {
		return a.DateCreated.Time.Unix() >= b.DateCreated.Time.Unix()
	})

	resp := util.Message(true, fmt.Sprintf(`All created posts of "%s"`, username))
	resp["posts"] = posts

	util.Respond(w, resp)
})
