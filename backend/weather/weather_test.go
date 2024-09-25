package weather

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetWeather(t *testing.T) {
	// Load environment variables from the absolute path to .env file
	envPath, err := filepath.Abs("../.env")
	if err != nil {
		t.Fatalf("Error resolving absolute path for .env file: %v", err)
	}

	err = godotenv.Load(envPath)
	if err != nil {
		t.Fatalf("Error loading .env file: %v", err)
	}

	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		t.Fatal("API_KEY is not set in .env file")
	}

	t.Logf("Using API_KEY: %s", apiKey) // Verifying that the API Key is loaded correctly

	// Perform the actual API call
	weatherResponse, err := GetWeather("Copenhagen", apiKey)
	if err != nil {
		t.Fatalf("Error fetching weather: %v", err)
	}

	// Logging the full response for debugging purposes
	t.Logf("Full weather response: %+v", weatherResponse)

	// Assertions to ensure we got valid data
	if weatherResponse.Name == "" {
		t.Error("Expected city name, got empty string")
	}
	if weatherResponse.Main.Temp == 0 {
		t.Error("Expected non-zero temperature, got zero")
	}
}
