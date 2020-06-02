package subscriptions

import (
	"checkaem_server/cmd/database/subscription"
	"checkaem_server/cmd/handlers/util"
	"log"
	"net/http"
)

var GetPosts = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

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

	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

	posts, err := subscription.GetPosts(username)

	if err != nil {
		util.RespondUserNotFound(w, username)
		return
	}

	resp := util.Message(true, "Posts was found")

	resp["posts"] = posts

	util.Respond(w, resp)
})

var GetTags = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

	tags, err := subscription.GetTagNames(username)

	if err != nil {
		log.Println(err)
		util.RespondUserNotFound(w, username)
		return
	}

	resp := util.Message(true, "Tags was found")
	resp["tags"] = tags

	util.Respond(w, resp)

})

var Modify = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	status, err := util.GetStatusFromQuery(r)

	if err != nil {
		log.Printf(err.Error())
		util.RespondInvalidRequest(w)
		return
	}

	tagName, err := util.GetFromQuery(r, "name")

	if err != nil {
		log.Println(err.Error())
		util.RespondInvalidRequest(w)
		return
	}

	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

	if status {
		err = subscription.Insert(username, tagName)
	} else {
		err = subscription.Delete(username, tagName)
	}

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "Subscription was successfully changed")
	util.Respond(w, resp)
})
