package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/UpsDev42069/BM_Search_Engine/backend/db"
)

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