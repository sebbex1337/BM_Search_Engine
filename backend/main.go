package main

import (
	//Internal imports
	"github.com/UpsDev42069/BM_Search_Engine/backend/handlers"
	"net/http"

	//External imports
	"github.com/gorilla/mux"
)


func main(){
	r := mux.NewRouter()
	r.HandleFunc("/", handlers.RootGet).Methods("GET")
	
	// Route for registering a new user
	r.HandleFunc("/api/register", handlers.RootPost).Methods("POST")
	http.ListenAndServe(":8080", r)
}