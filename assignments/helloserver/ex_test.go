package helloserver

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//TestHandler is used to test the handlerscrud func
func TestHandler(t *testing.T) {
	testCases := []struct {
		desc           string
		method         string
		input          string
		expectedOutput string
	}{
		{"Url With Name", "GET", "/hello?name=priyanka", "Hello!priyanka"},
		{"No Query", "GET", "/hello", "Hello!"},
		{"Url without name", "GET", "/hello?name=", "Bad Request!"},

		{"Other Methods", "POST", "/hello", "method not allowed"},
	}
	for i := range testCases {
		req := httptest.NewRequest(testCases[i].method, testCases[i].input, nil)
		w := httptest.NewRecorder()
		Handler(w, req)
		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Unexpected Error got %v", err)
		}
		if string(body) != testCases[i].expectedOutput {
			t.Errorf("Test Case Failed Desc :%v Excepeted %v got %v", testCases[i].desc, testCases[i].expectedOutput, string(body))
		}
	}
}

//BenchmarkHandler is used to test the performance
func BenchmarkHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/hello?name=priyanka", nil)
		w := httptest.NewRecorder()
		Handler(w, req)
	}
}
