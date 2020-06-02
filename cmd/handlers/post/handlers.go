package post

import (
	"bytes"
	dbPost "checkaem_server/cmd/database/post"
	"checkaem_server/cmd/entities/post"
	"checkaem_server/cmd/handlers/util"
	"checkaem_server/cmd/tika"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx"
	"io"
	"net/http"
)

var Get = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username, err := util.GetUsername(r)

	uuid := util.GetId(r)

	if err != nil || uuid == "" {
		util.RespondInvalidTokenPayload(w)
		return
	}

	post, err := dbPost.Get(username, uuid)

	if err != nil {
		if err == pgx.ErrNoRows {
			util.RespondNotFound(w, "Post not found")
			return
		}
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "Post was found")
	resp["post"] = post

	util.Respond(w, resp)
})

var GetCreated = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

	posts, err := dbPost.GetByCreator(username)

	if err != nil {
		if err == pgx.ErrNoRows {
			util.RespondUserNotFound(w, username)
		}
		return
	}

	resp := util.Message(true, fmt.Sprintf(`All created posts of "%s"`, username))
	resp["posts"] = posts

	util.Respond(w, resp)
})

var UploadFile = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	uuid := util.GetId(r)

	_ = r.ParseMultipartForm(250 << 20)

	file, h, err := r.FormFile("file")

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}
	defer file.Close()

	err = dbPost.UploadFile(file, uuid, h.Size)

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "File was successfully uploaded")

	util.Respond(w, resp)
})

var DownloadFile = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	uuid := util.GetId(r)

	byteArray, err := dbPost.DownloadFile(uuid)

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	file := bytes.NewReader(byteArray)

	fileType, err := tika.GetType(file)

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	name, err := dbPost.GetName(uuid)

	if err != nil {
		return
	}

	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%q", name))
	w.Header().Set("Content-Length", fmt.Sprintf("%v", len(byteArray)))
	w.Header().Set("Content-Type", fileType)
	w.Header().Set("Filename", name)

	file = bytes.NewReader(byteArray)

	io.Copy(w, file)
})

var Create = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	p := post.NewEmpty()

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		util.RespondInvalidRequest(w)
		return
	}

	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

	p.CreatorUsername = username
	p.Id = util.GetId(r)

	p, err = dbPost.Insert(p)

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "Post was successfully added")

	resp["post"] = p

	util.Respond(w, resp)
})

var Delete = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	postId := util.GetId(r)
	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

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

	p.Id = util.GetId(r)

	username, err := util.GetUsername(r)

	if err != nil {
		util.RespondInvalidTokenPayload(w)
		return
	}

	p.CreatorUsername = username

	p, err = dbPost.Modify(p)

	if err != nil {
		if err == dbPost.ErrNotAdmin {
			util.RespondNotAdmin(w)
		} else {
			util.RespondInternalServerError(w, err)
		}
		return
	}

	resp := util.Message(true, "Your post was successfully modified")
	resp["post"] = p

	util.Respond(w, resp)
})
