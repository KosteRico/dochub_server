package main

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/database/user"
	"checkaem_server/cmd/handlers/auth"
	"checkaem_server/cmd/handlers/middleware"
	"checkaem_server/cmd/handlers/post"
	"checkaem_server/cmd/handlers/subscriptions"
	"checkaem_server/cmd/handlers/util"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Book struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func init() {
	err := database.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database initializer: %v (PID)", database.Connection.PID())
}

func main() {
	defer database.Close()

	testHandler := http.HandlerFunc(TestFunc)

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()

	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userRouter.Handle("/new", auth.CreateAccount).Methods("POST")
	userRouter.Handle("/login", auth.Login).Methods("GET")

	authenticatedUserRouter := userRouter.PathPrefix("/{username}").Subrouter()
	authenticatedUserRouter.Use(middleware.IsAuthenticated)

	authenticatedUserRouter.Handle("/refresh", auth.RefreshToken).Methods("GET")
	authenticatedUserRouter.Handle("/test/{msg}", testHandler)
	authenticatedUserRouter.Handle("/posts", post.GetCreated).Methods("GET")
	authenticatedUserRouter.Handle("/posts", post.Create).Methods("POST")

	postsRouter := authenticatedUserRouter.PathPrefix("/posts/{id}").Subrouter()

	postsRouter.Handle("", post.Delete).Methods("DELETE")
	postsRouter.Handle("", post.Modify).Methods("PATCH")

	subscriptionsRouter := authenticatedUserRouter.PathPrefix("/subscriptions").Subrouter()
	subscriptionsRouter.Handle("/posts", subscriptions.GetSubscribedPosts).Methods("GET")
	subscriptionsRouter.Handle("/tags", subscriptions.GetSubscribedTags).Methods("GET")

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func TestFunc(w http.ResponseWriter, r *http.Request) {

	username := util.GetUsername(r)

	u, err := user.Get(username)

	if err != nil {
		util.RespondMessage(w, false, fmt.Sprintf("User with name %q not found", username))
		return
	}

	args := mux.Vars(r)

	resp := util.Message(true, fmt.Sprintf("Test message: %s", args["msg"]))

	resp["account"] = u

	util.Respond(w, resp)

}
