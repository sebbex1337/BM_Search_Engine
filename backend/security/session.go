package security

import (
	"net/http"
	"database/sql"

	"github.com/gorilla/sessions"
)

type User struct {
	ID int
	Username string
	Email string
}

var store = sessions.NewCookieStore([]byte("your-secret-key"))

// CreateSession creates a new session with userID
func CreateSession(w http.ResponseWriter, r *http.Request, userID string) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return err
	}

	session.Values["user_id"] = userID
	return session.Save(r, w)
}

// GetSession retrieves the session and returns the userID
func GetSession(r *http.Request, db *sql.DB) (*User, error) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return nil, err
	}

	userID, ok := session.Values["user_id"].(string)
	if !ok {
		return nil, nil // No session found
	}

	user, err := fetchUserByID(db, userID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// DestroySession destroys the current session
func DestroySession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return err
	}

	session.Options.MaxAge = -1
	return session.Save(r, w)
}

// GetUser retrieves the user from the database using the userID
func fetchUserByID(db *sql.DB, userID string) (*User, error) {
	var user User
	query := "SELECT id, username, email FROM users WHERE id = ?"
	err := db.QueryRow(query, userID).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}