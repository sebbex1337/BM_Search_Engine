package handlers

import (
	"net/http"
)

func RootGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

// Error logic if search query is empty
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	if query == "" {
		http.Error(w, "Missing query parameter", http.StatusBadRequest)
		return
	}
	w.Write([]byte("Searching for"))
	//w.Write([]byte("Searching for " + query))
}