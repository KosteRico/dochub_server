package main

import (
	"checkaem_server/cmd/database"
	"checkaem_server/cmd/handlers/auth"
	"checkaem_server/cmd/handlers/bookmarks"
	"checkaem_server/cmd/handlers/middleware"
	"checkaem_server/cmd/handlers/post"
	"checkaem_server/cmd/handlers/search"
	"checkaem_server/cmd/handlers/subscriptions"
	"checkaem_server/cmd/handlers/tags"
	"checkaem_server/cmd/searchEngine/redisDb"
	"checkaem_server/cmd/tika"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func init() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Println("File \".env\" wasn't initialized")
	}

	err = database.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database initialized")

	err = tika.Init()

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Apache Tike initialized")

	pongs, err := redisDb.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Redis initialized: %#v\n", pongs)
}

func main() {
	defer database.Close()
	defer tika.Close()

	r := mux.NewRouter()

	apiRouter := r.PathPrefix("/api").Subrouter()

	userRouter := apiRouter.PathPrefix("/users").Subrouter()
	userRouter.Handle("", auth.GetAllUsernames).Methods("GET")
	userRouter.Handle("/exists", auth.CheckUsernameExists).Methods("GET")
	userRouter.Handle("/new", auth.CreateAccount).Methods("POST")
	userRouter.Handle("/login", auth.Login).Methods("POST")

	authenticatedUserRouter := apiRouter.PathPrefix("").Subrouter()
	authenticatedUserRouter.Use(middleware.IsAuthenticated)

	authenticatedUserRouter.Handle("/refresh", auth.RefreshToken).Methods("GET")

	postsRouter := authenticatedUserRouter.PathPrefix("/posts").Subrouter()
	postsRouter.Handle("", post.GetCreated).Methods("GET")

	postIdRouter := postsRouter.PathPrefix("/{id}").Subrouter()
	postIdRouter.Handle("", post.Create).Methods("POST")
	postIdRouter.Handle("", post.Get).Methods("GET")
	postIdRouter.Handle("", post.Delete).Methods("DELETE")
	postIdRouter.Handle("", post.Modify).Methods("PATCH")
	postIdRouter.Handle("/file", post.UploadFile).Methods("POST")
	postIdRouter.Handle("/file", post.DownloadFile).Methods("GET")

	subscriptionsRouter := authenticatedUserRouter.PathPrefix("/subscriptions").Subrouter()
	subscriptionsRouter.Handle("", subscriptions.GetPosts).Methods("GET")
	subscriptionsRouter.Handle("/tags", subscriptions.GetTags).Methods("GET")
	subscriptionsRouter.Handle("/tags", subscriptions.Modify).Methods("POST")

	bookmarksRouter := authenticatedUserRouter.PathPrefix("/bookmarks").Subrouter()
	bookmarksRouter.Handle("", bookmarks.GetPosts).Methods("GET")
	bookmarksRouter.Handle("/{id}", bookmarks.Modify).Methods("POST")

	tagsRouter := authenticatedUserRouter.PathPrefix("/tags").Subrouter()
	tagsRouter.Handle("", tags.GetPosts).Methods("GET")
	tagsRouter.Handle("", tags.Create).Methods("POST")
	tagsRouter.Handle("/all", tags.GetAll).Methods("GET")
	tagsRouter.Handle("/search", tags.SearchTags).Methods("GET")

	searchRouter := authenticatedUserRouter.PathPrefix("/search").Subrouter()
	searchRouter.Handle("", search.Search).Methods("GET")

	c := corsConfig("http://localhost:8080")

	log.Println("Server started")
	log.Fatal(http.ListenAndServe(":8000", c.Handler(r)))
}
