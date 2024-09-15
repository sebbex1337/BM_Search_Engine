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
	Email    string `json:"email"`
}

// RegisterHandler handles user registration
func RegisterHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var user User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		if user.Username == "" || user.Password == "" || user.Email == "" {
			http.Error(w, "Missing required fields", http.StatusBadRequest)
			return
		}

		// Save the user to the database
		insertUserSQL := `INSERT INTO users (username, password, email) VALUES (?, ?, ?)`
		_, err = database.Exec(insertUserSQL, user.Username, user.Password, user.Email)
		if err != nil {
			http.Error(w, "Failed to register user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User registered successfully"))
	}
}


// Error logic if search query is empty
func SearchHandlerLucas(w http.ResponseWriter, r *http.Request) {
	// function body
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}
	
	w.Write([]byte("Searching for " + query))
}

