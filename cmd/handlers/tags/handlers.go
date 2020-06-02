package tags

import (
	"checkaem_server/cmd/database/tag"
	tagStruct "checkaem_server/cmd/entities/tag"
	"checkaem_server/cmd/handlers/util"
	"fmt"
	"net/http"
)

func getTagName(r *http.Request) string {
	tagNameArray, ok := r.URL.Query()["name"]

	if !ok || len(tagNameArray) != 1 {
		return ""
	}

	return tagNameArray[0]
}

var GetPosts = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	tagName := getTagName(r)

	if tagName == "" {
		util.RespondInvalidRequest(w)
		return
	}
	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

	t, err := tag.Get(username, tagName)

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, fmt.Sprintf("Posts for %q was successfully recieved", tagName))

	resp["tag"] = t

	util.Respond(w, resp)
})

var Create = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	tagName := getTagName(r)

	if tagName == "" {
		util.RespondInvalidRequest(w)
		return
	}

	t := tagStruct.New(tagName)

	t, err := tag.Insert(t)

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "Tag was successfully created")
	resp["tag"] = t

	util.Respond(w, resp)
})

var GetAll = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	tags, err := tag.GetAll()

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "All tags was successfully retrieved")
	resp["tags"] = tags

	util.Respond(w, resp)
})

var SearchTags = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	query, err := util.GetFromQuery(r, "q")

	if err != nil {
		util.RespondInvalidRequest(w)
		return
	}

	tags, err := SearchTagsFunc(query, 4)

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "Tags was found")

	resp["tags"] = tags

	util.Respond(w, resp)
})
