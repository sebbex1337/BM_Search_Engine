package handlers

// Imports for the handlers
import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/UpsDev42069/BM_Search_Engine/backend/db"
	"github.com/UpsDev42069/BM_Search_Engine/backend/security"
)

type AuthResponse struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

type SearchResponse struct {
	Data []map[string]interface{} `json:"data"`
}

type StandardResponse struct {
	Data map[string]interface{} `json:"data"`
}

type RequestValidationError struct {
	StatusCode int `json:"statusCode"`
	Message string `json:"message"`
}

type HTTPValidationError struct {
	Detail []ValidationError `json:"detail"`
}

type ValidationError struct {
	Loc []interface{} `json:"loc"`
	Msg string `json:"msg"`
	Type string `json:"type"`
}


// RootGet handles the root GET request
func RootGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}


func SearchHandler(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("q")
		if q == "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(RequestValidationError{
				StatusCode: 422,
				Message: "Query parameter is required",
			})
			return
		}

		language := r.URL.Query().Get("language")
		if language == "" {
			language = "en"
		}

		var searchResults []map[string]interface{}
		query := "SELECT * FROM pages WHERE language = ? AND content LIKE ?"
		args := []interface{}{language, "%"+q+"%"}
		results, err := db.QueryDB(database, query, args...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		searchResults = results

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SearchResponse{
			Data: searchResults,
		})
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
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(HTTPValidationError{
				Detail: []ValidationError{
					{
						Loc: []interface{}{"body", "username"},
						Msg: "field required",
						Type: "value_error.missing",
					},
					{
						Loc: []interface{}{"body", "password"},
						Msg: "field required",
						Type: "value_error.missing",
					},
					{
						Loc: []interface{}{"body", "email"},
						Msg: "field required",
						Type: "value_error.missing",
					},
				},
			})
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

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(AuthResponse{StatusCode: http.StatusOK, Message: "User registered successfully"})
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
            w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusUnprocessableEntity)
			json.NewEncoder(w).Encode(HTTPValidationError{
				Detail: []ValidationError{
					{
						Loc: []interface{}{"body", "username"},
						Msg: "field required",
						Type: "value_error.missing",
					},
					{
						Loc: []interface{}{"body", "password"},
						Msg: "field required",
						Type: "value_error.missing",
					},
				},
			})
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
        if !security.CheckPasswordHash(dbUser.Password, user.Password) {
            http.Error(w, "Invalid credentials", http.StatusUnauthorized)
            return
        }

		//Create a new session
		err = security.CreateSession(w, r, dbUser.Username)
		if err != nil {
			http.Error(w, "Failed to create session", http.StatusInternalServerError)
			return
		}

        // Return success response
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(AuthResponse{StatusCode: http.StatusOK, Message: "Login successful"})
    }
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	err := security.DestroySession(w, r)
	if err != nil {
		http.Error(w, "Failed to destroy session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(AuthResponse{StatusCode: http.StatusOK, Message: "Logout successful"})
}