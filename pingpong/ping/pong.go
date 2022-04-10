package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("pong"))

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ping", Handler).Methods(http.MethodGet)
	http.Handle("/", r)
	log.Println("Listen at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
