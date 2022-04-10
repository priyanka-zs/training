package handlers

import (
	"fmt"
	"net/http"
)

func HelloServer(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, "Hello")
	if err != nil {
		return
	}
}
