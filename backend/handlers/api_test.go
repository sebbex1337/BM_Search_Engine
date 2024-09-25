package handlers_test

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"golang.org/x/crypto/bcrypt" // Import bcrypt for password hashing
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/UpsDev42069/BM_Search_Engine/backend/handlers"
	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

// Mock database connection
var mockDB *sql.DB

// User struct (assuming it's defined somewhere in your code)
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// CheckPasswordHash compares a plain password with a hashed password
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func setupMockDB() {
	var err error
	mockDB, err = sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}

	// Create the users table
	createTableSQL := `
    CREATE TABLE users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        password TEXT NOT NULL,
        email TEXT NOT NULL
    );`
	_, err = mockDB.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
}

func TestLoginHandler(t *testing.T) {
	setupMockDB() // Initialize the mock database

	// Insert a test user into the mock database
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("testpassword"), bcrypt.DefaultCost)
	_, err := mockDB.Exec("INSERT INTO users (username, password, email) VALUES (?, ?, ?)", "testuser", string(hashedPassword), "test@example.com")
	if err != nil {
		t.Fatalf("Failed to insert test user: %v", err)
	}

	tests := []struct {
		name         string
		user         User
		expectedCode int
		expectedBody string
	}{
		{
			name: "valid login",
			user: User{
				Username: "testuser",
				Password: "testpassword",
			},
			expectedCode: http.StatusOK,
			expectedBody: `{"message":"Login successful"}`,
		},
		{
			name: "invalid login",
			user: User{
				Username: "wronguser",
				Password: "wrongpassword",
			},
			expectedCode: http.StatusUnauthorized,
			expectedBody: "Invalid credentials\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			payload, err := json.Marshal(tt.user)
			if err != nil {
				t.Fatalf("Failed to marshal user: %v", err)
			}

			req, err := http.NewRequest(http.MethodPost, "/api/login", bytes.NewBuffer(payload))
			if err != nil {
				t.Fatalf("Failed to create request: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := handlers.LoginHandler(mockDB)
			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.expectedCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedCode)
			}

			respBody := rr.Body.String()
			if respBody != tt.expectedBody {
				t.Errorf("expected body %q, got %q", tt.expectedBody, respBody)
			}
		})
	}
}
