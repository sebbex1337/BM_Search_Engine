package handlers

// Imports for the handlers
import (
	"database/sql"
	"encoding/json"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"github.com/UpsDev42069/BM_Search_Engine/backend/db"
	"github.com/UpsDev42069/BM_Search_Engine/backend/security"
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

		// Hash password before saving it to the database
		hashedPassword, err := security.HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		
		_, err = database.Exec(insertUserSQL, user.Username, hashedPassword, user.Email)
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

/////////////////////////////////////////
//	 	Login logic for useres         //
/////////////////////////////////////////

// LoginHandler handles user login
func LoginHandler(database *sql.DB) http.HandlerFunc {
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

        if user.Username == "" || user.Password == "" {
            http.Error(w, "Missing required fields", http.StatusBadRequest)
            return
        }

        // Query the database for the user
        var dbUser User
        err = database.QueryRow("SELECT username, password, email FROM users WHERE username = ?", user.Username).Scan(&dbUser.Username, &dbUser.Password, &dbUser.Email)
        if err != nil {
            if err == sql.ErrNoRows {
                http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            } else {
                http.Error(w, "Internal server error", http.StatusInternalServerError)
            }
            return
        }

        // Validate the password
        if !CheckPasswordHash(user.Password, dbUser.Password) {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

        // Return success response
        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"message":"Login successful"}`))
    }
}

// CheckPasswordHash compares a plain text password with a hashed password
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}