package main

import (
	"log"
	"net/http"
	"os"

	"github.com/UpsDev42069/BM_Search_Engine/backend/db"
	"github.com/UpsDev42069/BM_Search_Engine/backend/handlers"

	_ "github.com/UpsDev42069/BM_Search_Engine/backend/docs"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title BM Search Engine API
// @version 0.1.0
// @description This is a sample server for a BM Search Engine.
// @host localhost:8080
// @BasePath /

func main() {

	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API_KEY is not set in .env file")
	}

	// Connecting to the database
	database, err := db.ConnectDB(false)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	if err := db.InitDB(); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	// Existing routes
	r.HandleFunc("/", handlers.RootGet).Methods("GET")
	r.HandleFunc("/api/search", handlers.SearchHandler(database)).Methods("GET")
	r.HandleFunc("/api/register", handlers.RegisterHandler(database)).Methods("POST")
	r.HandleFunc("/api/login", handlers.LoginHandler(database)).Methods("POST")
	r.HandleFunc("/api/weather", handlers.WeatherHandler).Methods("GET")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Println("Server started at :8080")
	log.Println("http://localhost:8080")

	http.ListenAndServe(":8080", r)
}
