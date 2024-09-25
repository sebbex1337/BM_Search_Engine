package security

import (
	"net/http"

	"github.com/gorilla/sessions"
)

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
func GetSession(r *http.Request) (string, error) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		return "", err
	}

	userID, ok := session.Values["user_id"].(string)
	if !ok {
		return "", nil // No session found
	}

	return userID, nil
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
