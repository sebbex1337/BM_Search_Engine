package main

import (
	"log"
	"net/http"
	"os"

	"github.com/UpsDev42069/BM_Search_Engine/backend/db"
	"github.com/UpsDev42069/BM_Search_Engine/backend/handlers"
	"github.com/UpsDev42069/BM_Search_Engine/backend/weather"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

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

	// Fetching weather information
	weatherResponse, err := weather.GetWeather("Copenhagen", apiKey)
	if err != nil {
		log.Fatalf("Error fetching weather: %v", err)
	}
	log.Printf("Weather data for %s: Temperature %.2f°C", weatherResponse.Name, weatherResponse.Main.Temp)

	r := mux.NewRouter()

	// Existing routes
	r.HandleFunc("/", handlers.RootGet).Methods("GET")
	r.HandleFunc("/api/search", handlers.SearchHandler(database)).Methods("GET")
	r.HandleFunc("/api/register", handlers.RegisterHandler(database)).Methods("POST")
	r.HandleFunc("/api/login", handlers.LoginHandler(database)).Methods("POST")
	r.HandleFunc("/api/weather", handlers.WeatherHandler).Methods("GET")

	log.Println("Server started at :8080")
	log.Println("http://localhost:8080")

	http.ListenAndServe(":8080", r)
}
