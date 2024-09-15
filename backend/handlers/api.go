package handlers

// Imports for the handlers
import (

	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/UpsDev42069/BM_Search_Engine/backend/db"

)

// RootGet handles the root GET request
func RootGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}


func SearchHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		language := r.URL.Query().Get("language")
		if language == "" {
			language = "en"
		}

		var searchResults []map[string]interface{}
		if q != "" {
			query := "SELECT * FROM pages WHERE language = ? AND content LIKE ?"
			args := []interface{}{language, "%"+q+"%"}
			results, err := db.QueryDB(database, query, args...)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			searchResults = results
		}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(searchResults); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// RootPost handles the root POST request
func RootPost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}


/////////////////////////////////////////
// Funtions for registering a new user //
/////////////////////////////////////////

// User represents the user registration data
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// HandleRegister handles the user registration
func HandleRegister(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the request body
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Validate the user data
	if user.Username == "" || user.Password == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}
	
	// TODO save your user to a database

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User registered successfully"))
}

// Error logic if search query is empty
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}
	
	w.Write([]byte("Searching for " + query))
}

