package main

import (

	//Internal imports
	"log"
	"net/http"

	"github.com/UpsDev42069/BM_Search_Engine/backend/db"
	"github.com/UpsDev42069/BM_Search_Engine/backend/handlers"

	//External imports
	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)


func main(){
	database, err := db.ConnectDB(false)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()
	handlers.SetRouter(r)

	// PAGE ROUTES
	r.HandleFunc("/", handlers.RootTemplateHandler(database)).Methods("GET").Name("root")
	r.HandleFunc("/search", handlers.SearchTemplateHandler(database)).Methods("GET").Name("search")
	r.HandleFunc("/about", handlers.AboutTemplateHandler).Methods("GET").Name("about")
	r.HandleFunc("/login", handlers.LoginTemplateHandler).Methods("GET").Name("login")
	r.HandleFunc("/register", handlers.RegisterTemplateHandler).Methods("GET").Name("register")

	// API ROUTES
	r.HandleFunc("/api/search", handlers.SearchHandler(database)).Methods("GET")
	r.HandleFunc("/api/register", handlers.RegisterHandler(database)).Methods("POST")
	r.HandleFunc("api/login", handlers.LoginHandler(database)).Methods("POST")


	log.Println("Server started at :8080")
	log.Println("http://localhost:8080")

	http.ListenAndServe(":8080", r)
}