package helloserver

import (
	"net/http"
)

//Handler is used to display hello message
func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if len(query) == 0 && r.Method == http.MethodGet {
		w.Write([]byte("Hello!"))
		return
	} else if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))

	} else if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad Request!"))

	} else {
		w.Write([]byte("Hello!" + name))
	}
}
