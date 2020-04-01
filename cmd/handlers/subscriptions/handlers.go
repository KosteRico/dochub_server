package subscriptions

import (
	"checkaem_server/cmd/database/subscription"
	"checkaem_server/cmd/entities/post"
	"checkaem_server/cmd/handlers/util"
	"net/http"
)

var GetSubscribedPosts = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	//limit, err := strconv.Atoi(r.URL.Query()["limit"][0])
	//
	//if err != nil {
	//	util.RespondInvalidRequest(w)
	//	return
	//}
	//
	//offset, err := strconv.Atoi(r.URL.Query()["offset"][0])
	//
	//if err != nil {
	//	util.RespondInvalidRequest(w)
	//	return
	//}

	username := util.GetUsername(r)

	posts, err := subscription.GetPosts(username)

	if err != nil {
		util.RespondUserNotFound(w, username)
		return
	}

	resp := util.Message(true, "Posts was found")

	posts = post.Mergesort(posts, func(a, b *post.Post) bool {
		return a.DateCreated.Time.Unix() <= b.DateCreated.Time.Unix()
	})

	resp["posts"] = posts

	util.Respond(w, resp)
})
