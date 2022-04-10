package handlers

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_HandlerServes(t *testing.T) {
	testcases := []struct {
		desc     string
		exOutput string
	}{
		{"success", "Hello"},
	}

	for _, tc := range testcases {
		r := httptest.NewRequest(http.MethodGet, "https://hello", nil)
		w := httptest.NewRecorder()
		HelloServer(w, r)
		assert.Equal(t, tc.exOutput, w.Body.String())
	}
}
