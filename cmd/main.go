package main

import (
	"encoding/json"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Book struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func SetJsonHeader(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
}

func main() {
	apiRouter := mux.NewRouter().PathPrefix("/api").Subrouter()

	CSRF := csrf.Protect([]byte("32-byte-long-auth-key"))

	apiRouter.HandleFunc("/test/{id}", TestHandler)

	log.Println("Server started")
	http.Handle("/", CSRF(apiRouter))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	SetJsonHeader(&w)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return
	}
	b := Book{
		Id:    id,
		Title: "War and peace",
	}
	json.NewEncoder(w).Encode(b)
}
