package handlers

import (
	"database/sql"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/UpsDev42069/BM_Search_Engine/backend/security"

	"github.com/gorilla/mux"
)

var router *mux.Router

func SetRouter(r *mux.Router) {
	router = r
}

func urlFunc(name string, params ...string) string {
	url, err := router.Get(name).URL(params...)
	if err != nil {
		log.Println("Error generating URL:", err)
		return ""
	}

	return url.String()
}

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
    funcMap := template.FuncMap{
        "url": urlFunc,
    }

    tmplPath := filepath.Join("templates", tmpl)
    t, err := template.New(filepath.Base(tmplPath)).Funcs(funcMap).ParseFiles(tmplPath)
    if err != nil {
        http.Error(w, "Unable to load template", http.StatusInternalServerError)
        log.Println("Error loading template:", err)
        return
    }

    err = t.Execute(w, data)
    if err != nil {
        http.Error(w, "Unable to render template", http.StatusInternalServerError)
        log.Println("Error rendering template:", err)
    }
}

type SearchResult struct {
	content []string
	language string
	last_updated string
	title string
	url string
}

func RootTemplateHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("q")
		searchResults := []SearchResult{}

		// Make an internal request to the /api/search endpoint
		apiURL := "http://localhost:8080/api/search?q=" + query
		resp, err := http.Get(apiURL)
		if err != nil {
			http.Error(w, "Failed to fetch search results", http.StatusInternalServerError)
			log.Println("Error fetching search results:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			err = json.NewDecoder(resp.Body).Decode(&searchResults)
			if err != nil {
				http.Error(w, "Failed to decode search results", http.StatusInternalServerError)
				log.Println("Error decoding search results:", err)
				return
			}
		}

		user, err := security.GetSession(r, db)
		if err != nil {
			http.Error(w, "Failed to get user session", http.StatusInternalServerError)
			log.Println("Error getting user session:", err)
			return
		}

		data := struct {
			Query         string
			SearchResults []SearchResult
			User          *security.User
		}{
			Query:         query,
			SearchResults: searchResults,
			User:          user,
		}

		RenderTemplate(w, "search.html", data)
	}
}

func AboutTemplateHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "about.html", nil)
}

func LoginTemplateHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "login.html", nil)
}

func RegisterTemplateHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "register.html", nil)
}

func SearchTemplateHandler(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        query := r.URL.Query().Get("q")
        searchResults := []SearchResult{}

        // Make an internal request to the /api/search endpoint
        apiURL := "http://localhost:8080/api/search?q=" + query
        resp, err := http.Get(apiURL)
        if err != nil {
            http.Error(w, "Failed to fetch search results", http.StatusInternalServerError)
            log.Println("Error fetching search results:", err)
            return
        }
        defer resp.Body.Close()

        if resp.StatusCode == http.StatusOK {
            err = json.NewDecoder(resp.Body).Decode(&searchResults)
            if err != nil {
                http.Error(w, "Failed to decode search results", http.StatusInternalServerError)
                log.Println("Error decoding search results:", err)
                return
            }
        }

        user, err := security.GetSession(r, db)
        if err != nil {
            http.Error(w, "Failed to get user session", http.StatusInternalServerError)
            log.Println("Error getting user session:", err)
            return
        }

        data := struct {
            Query         string
            SearchResults []SearchResult
            User          *security.User
        }{
            Query:         query,
            SearchResults: searchResults,
            User:          user,
        }

        RenderTemplate(w, "search.html", data)
    }
}