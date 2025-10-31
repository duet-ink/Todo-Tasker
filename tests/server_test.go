package tests

import (
	"Todo-Tasker/server"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestNew tests the main server route creation
func TestNew(t *testing.T) {
	mux := server.New()
	if mux == nil {
		t.Fatal("Expected non-nil ServeMux from New()")
	}

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
		checkRedirect  bool
		redirectTo     string
	}{
		{
			name:           "index page",
			method:         http.MethodGet,
			path:           "/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "404 page",
			method:         http.MethodGet,
			path:           "/404/",
			expectedStatus: http.StatusNotFound,
		},
		{
			name:           "error page",
			method:         http.MethodGet,
			path:           "/error/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "non-existent route redirects to 404",
			method:         http.MethodGet,
			path:           "/nonexistent",
			expectedStatus: http.StatusFound,
			checkRedirect:  true,
			redirectTo:     "/404",
		},
		{
			name:           "components POST endpoint",
			method:         http.MethodPost,
			path:           "/c/navbar",
			expectedStatus: http.StatusFound, // Will redirect to 404 if component doesn't exist
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.checkRedirect {
				location := w.Header().Get("Location")
				if !strings.Contains(location, tt.redirectTo) {
					t.Errorf("Expected redirect to contain %s, got %s", tt.redirectTo, location)
				}
			}
		})
	}
}

// TestNewAdmin tests the admin server route creation
func TestNewAdmin(t *testing.T) {
	mux := server.NewAdmin()
	if mux == nil {
		t.Fatal("Expected non-nil ServeMux from NewAdmin()")
	}

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "admin page",
			method:         http.MethodGet,
			path:           "/admin/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "non-admin route returns 404",
			method:         http.MethodGet,
			path:           "/",
			expectedStatus: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

// TestIndexPage tests the index page handler
func TestIndexPage(t *testing.T) {
	mux := server.New()

	tests := []struct {
		name           string
		path           string
		expectedStatus int
		checkRedirect  bool
	}{
		{
			name:           "root path returns OK",
			path:           "/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "non-root path redirects to 404",
			path:           "/something",
			expectedStatus: http.StatusFound,
			checkRedirect:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.path, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}

			if tt.checkRedirect {
				location := w.Header().Get("Location")
				if location == "" {
					t.Error("Expected redirect but got none")
				}
			}

			if w.Code == http.StatusOK {
				body := w.Body.String()
				if body == "" {
					t.Error("Expected non-empty response body")
				}
			}
		})
	}
}

// TestAdminPage tests the admin page handler
func TestAdminPage(t *testing.T) {
	mux := server.NewAdmin()

	req := httptest.NewRequest(http.MethodGet, "/admin/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	body := w.Body.String()
	if body == "" {
		t.Error("Expected non-empty response body")
	}
}

// TestPageNotFound tests the 404 error handler
func TestPageNotFound(t *testing.T) {
	mux := server.New()

	req := httptest.NewRequest(http.MethodGet, "/404/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d, got %d", http.StatusNotFound, w.Code)
	}

	body := w.Body.String()
	if body == "" {
		t.Error("Expected non-empty response body for 404 page")
	}

	// Check if the response contains error-related content
	if !strings.Contains(body, "404") && !strings.Contains(body, "not found") && !strings.Contains(body, "Not found") {
		t.Log("Warning: 404 page body may not contain expected error message")
	}
}

// TestErrorPage tests the generic error handler
func TestErrorPage(t *testing.T) {
	mux := server.New()

	req := httptest.NewRequest(http.MethodGet, "/error/", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status %d, got %d", http.StatusOK, w.Code)
	}

	body := w.Body.String()
	if body == "" {
		t.Error("Expected non-empty response body for error page")
	}

	// Check if the response contains error-related content
	if !strings.Contains(body, "500") && !strings.Contains(body, "error") && !strings.Contains(body, "Error") {
		t.Log("Warning: Error page body may not contain expected error message")
	}
}

// TestComponentsPage tests the dynamic components endpoint
func TestComponentsPage(t *testing.T) {
	mux := server.New()

	tests := []struct {
		name           string
		componentName  string
		expectedStatus int
		checkRedirect  bool
	}{
		{
			name:           "component request",
			componentName:  "navbar",
			expectedStatus: http.StatusFound, // Will redirect if component doesn't exist
			checkRedirect:  true,
		},
		{
			name:           "another component request",
			componentName:  "footer",
			expectedStatus: http.StatusFound,
			checkRedirect:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/c/"+tt.componentName, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			// Component endpoint will redirect to 404 if component doesn't exist
			// or return OK if it does
			if w.Code != http.StatusOK && w.Code != http.StatusFound {
				t.Errorf("Expected status %d or %d, got %d", http.StatusOK, http.StatusFound, w.Code)
			}
		})
	}
}

// TestStaticAssets tests that the assets handler is properly configured
func TestStaticAssets(t *testing.T) {
	mux := server.New()

	// Test that assets path is handled (even if specific asset doesn't exist)
	req := httptest.NewRequest(http.MethodGet, "/assets/test.css", nil)
	w := httptest.NewRecorder()

	mux.ServeHTTP(w, req)

	// Should either return OK (if asset exists) or 404 (if it doesn't)
	// but should not redirect to our custom 404 page
	if w.Code != http.StatusOK && w.Code != http.StatusNotFound {
		t.Errorf("Expected status %d or %d for assets, got %d", http.StatusOK, http.StatusNotFound, w.Code)
	}
}

// TestHTTPMethods tests that routes respond correctly to different HTTP methods
func TestHTTPMethods(t *testing.T) {
	mux := server.New()

	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "GET index",
			method:         http.MethodGet,
			path:           "/",
			expectedStatus: http.StatusOK,
		},
		{
			name:           "POST to index not allowed",
			method:         http.MethodPost,
			path:           "/",
			expectedStatus: http.StatusMethodNotAllowed,
		},
		{
			name:           "POST to components allowed",
			method:         http.MethodPost,
			path:           "/c/test",
			expectedStatus: http.StatusFound, // or OK if component exists
		},
		{
			name:           "GET to components not allowed",
			method:         http.MethodGet,
			path:           "/c/test",
			expectedStatus: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			if w.Code != tt.expectedStatus {
				t.Errorf("Expected status %d, got %d", tt.expectedStatus, w.Code)
			}
		})
	}
}

// TestContentType tests that responses have appropriate content types
func TestContentType(t *testing.T) {
	mux := server.New()

	tests := []struct {
		name                string
		path                string
		method              string
		expectedContentType string
	}{
		{
			name:                "HTML page returns text/html",
			path:                "/",
			method:              http.MethodGet,
			expectedContentType: "text/html",
		},
		{
			name:                "404 page returns text/html",
			path:                "/404/",
			method:              http.MethodGet,
			expectedContentType: "text/html",
		},
		{
			name:                "Error page returns text/html",
			path:                "/error/",
			method:              http.MethodGet,
			expectedContentType: "text/html",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			w := httptest.NewRecorder()

			mux.ServeHTTP(w, req)

			contentType := w.Header().Get("Content-Type")
			if !strings.Contains(contentType, tt.expectedContentType) {
				t.Errorf("Expected content type to contain %s, got %s", tt.expectedContentType, contentType)
			}
		})
	}
}
