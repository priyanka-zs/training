package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockHandler struct{}

// ServeHTTP is a mock serveHttp function
func (m mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// TestMiddleware checks if given api-key is valid
func TestMiddleware(t *testing.T) {
	testCases := []struct {
		desc       string
		value      string
		statusCode int
	}{
		{"Authentication successful", "123", http.StatusOK},
		{"Authentication failed", "315", http.StatusUnauthorized},
	}

	for i, v := range testCases {
		req := httptest.NewRequest(http.MethodPost, "https://car", nil)
		req.Header.Add("api-key", v.value)

		w := httptest.NewRecorder()

		a := Middleware(mockHandler{})
		a.ServeHTTP(w, req)

		res := w.Result()

		assert.Equal(t, v.statusCode, res.StatusCode, "Test case %d Failed: %s", i, v.desc)
		err := res.Body.Close()

		if err != nil {
			return
		}
	}
}
