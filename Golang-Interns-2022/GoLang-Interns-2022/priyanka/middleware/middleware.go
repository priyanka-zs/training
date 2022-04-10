package middleware

import "net/http"

func Middleware(handleFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("api-key") != "123" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		handleFunc.ServeHTTP(w, r)
	})
}
