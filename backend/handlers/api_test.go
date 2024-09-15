package handlers_test

import (
    "bytes"
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/UpsDev42069/BM_Search_Engine/backend/handlers"
)

func TestHandleRegister(t *testing.T) {
    tests := []struct {
        name           string
        method         string
        body           interface{}
        expectedStatus int
        expectedBody   string
    }{
        {
            name:           "Valid request",
            method:         http.MethodPost,
            body:           handlers.User{Username: "testuser", Password: "password"},
            expectedStatus: http.StatusCreated,
            expectedBody:   "User registered successfully",
        },
        {
            name:           "Invalid method",
            method:         http.MethodGet,
            body:           nil,
            expectedStatus: http.StatusMethodNotAllowed,
            expectedBody:   "Method not allowed\n",
        },
        {
            name:           "Invalid payload",
            method:         http.MethodPost,
            body:           "invalid",
            expectedStatus: http.StatusBadRequest,
            expectedBody:   "Invalid request payload\n",
        },
        {
            name:           "Missing fields",
            method:         http.MethodPost,
            body:           handlers.User{Username: "", Password: ""},
            expectedStatus: http.StatusBadRequest,
            expectedBody:   "Missing required fields\n",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            var body []byte
            if tt.body != nil {
                body, _ = json.Marshal(tt.body)
            }

            req := httptest.NewRequest(tt.method, "/register", bytes.NewBuffer(body))
            w := httptest.NewRecorder()

            handlers.HandleRegister(w, req)

            resp := w.Result()
            defer resp.Body.Close()

            if resp.StatusCode != tt.expectedStatus {
                t.Errorf("expected status %d, got %d", tt.expectedStatus, resp.StatusCode)
            }

            respBody := new(bytes.Buffer)
            respBody.ReadFrom(resp.Body)
            if respBody.String() != tt.expectedBody {
                t.Errorf("expected body %q, got %q", tt.expectedBody, respBody.String())
            }
        })
    }
}