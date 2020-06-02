package search

import (
	"checkaem_server/cmd/handlers/util"
	"checkaem_server/cmd/searchEngine/search"
	"net/http"
)

var Search = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	query, err := util.GetFromQuery(r, "q")

	if err != nil {
		util.RespondInvalidRequest(w)
		return
	}

	ids, err := search.SearchRanking(query)

	if err != nil {
		util.RespondInternalServerError(w, err)
		return
	}

	resp := util.Message(true, "found")

	resp["ids"] = ids

	util.Respond(w, resp)
})
