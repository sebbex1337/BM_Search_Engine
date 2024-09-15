package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSearchHandler(t *testing.T) {
	tests := []struct {
		name         string
		query        string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "valid query",
			query:        "q=test",
			expectedCode: http.StatusOK,
			expectedBody: "Searching for test",
		},
		{
			name:         "missing query parameter",
			query:        "",
			expectedCode: http.StatusBadRequest,
			expectedBody: "Missing query parameter\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/api/search?"+tt.query, nil)
			if err != nil {
				t.Fatalf("could not create request: %v", err)
			}

			// Create a ResponseRecorder to record the response.
			rr := httptest.NewRecorder()

			// Call the handler function
			SearchHandler(rr, req)

			// Check the status code
			if got := rr.Code; got != tt.expectedCode {
				t.Errorf("expected status %v; got %v", tt.expectedCode, got)
			}

			// Check the response body
			if got := rr.Body.String(); got != tt.expectedBody {
				t.Errorf("expected body %q; got %q", tt.expectedBody, got)
			}
		})
	}
}