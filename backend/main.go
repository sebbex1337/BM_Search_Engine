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
	r.HandleFunc("/", handlers.RootGet).Methods("GET")
	r.HandleFunc("/api/search", handlers.SearchHandler(database)).Methods("GET")

	log.Println("Server started at :8080")
	log.Println("http://localhost:8080")
	http.ListenAndServe(":8080", r)
}