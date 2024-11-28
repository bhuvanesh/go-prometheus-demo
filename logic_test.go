package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// func TestDouble(t *testing.T) {
// 	req := httptest.NewRequest(http.MethodGet, "/double/5", nil)
// 	w := httptest.NewRecorder()

// 	doubleHandler(w, req)
// 	res := w.Result()
// 	if res.StatusCode != http.StatusOK {
// 		t.Errorf("Expected status code %d, got %d", http.StatusOK, res.StatusCode)
// 	}
// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		t.Errorf("Error reading response body: %v", err)
// 	}
// 	if string(body) != "5 doubled is 10" {
// 		t.Errorf("Expected response body '5 doubled is 10', got %s", string(body))
// 	}

// }

func TestDouble(t *testing.T) {
	// Create test cases
	tests := []struct {
		name         string
		path         string
		expectedCode int
		expectedBody string
	}{
		{
			name:         "valid number",
			path:         "/double/5",
			expectedCode: http.StatusOK,
			expectedBody: "5 doubled is 10",
		},
		{
			name:         "invalid number",
			path:         "/double/abc",
			expectedCode: http.StatusBadRequest,
			expectedBody: "strconv.Atoi: parsing \"abc\": invalid syntax",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tc.path, nil)
			w := httptest.NewRecorder()

			doubleHandler(w, req)

			res := w.Result()
			defer res.Body.Close()

			if res.StatusCode != tc.expectedCode {
				t.Errorf("Expected status code %d, got %d", tc.expectedCode, res.StatusCode)
			}

			body, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("Error reading response body: %v", err)
			}

			if strings.TrimSpace(string(body)) != strings.TrimSpace(tc.expectedBody) {
				t.Errorf("Expected response body '%s', got '%s'", tc.expectedBody, string(body))
			}
		})
	}
}
