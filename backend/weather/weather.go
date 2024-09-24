package weather

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// baseURL is the base URL for the OpenWeatherMap API (corrected to v2.5)
const baseURL = "https://api.openweathermap.org/data/2.5/weather"

// WeatherResponse represents the weather data returned by the API
type WeatherResponse struct {
	Main struct {
		Temp float64 `json:"temp"` // Temperature in Celsius
	} `json:"main"`
	Name string `json:"name"` // Name of the city
}

// GetWeather gets the weather data for a given city
// It returns a WeatherResponse struct and an error if there is one
// Parameters:
// city - the name of the city we want to know the weather for
// apiKey - the API key for the OpenWeatherMap API
// Returns:
// - A pointer to a WeatherResponse struct containing the weather data
func GetWeather(city, apiKey string) (*WeatherResponse, error) {
	endpoint, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("error parsing endpoint: %w", err)
	}

	// Set the query parameters
	params := url.Values{}
	params.Add("q", city)
	params.Add("appid", apiKey)
	params.Add("units", "metric")
	endpoint.RawQuery = params.Encode()

	// Make the request
	resp, err := http.Get(endpoint.String())
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check status code
	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body) // Reading the body for detailed error message
		return nil, fmt.Errorf("unexpected status code: %d, body: %s", resp.StatusCode, string(body))
	}

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Parse the response body as JSON into the WeatherResponse Struct
	var weatherResponse WeatherResponse
	if err := json.Unmarshal(body, &weatherResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response body: %w", err)
	}
	return &weatherResponse, nil
}