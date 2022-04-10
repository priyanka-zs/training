package main

import (
	"database/sql"
	"net/http"
	"petstore/handlers"
	"petstore/services"
	"petstore/stores"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func DBConn() *sql.DB {
	db, err := sql.Open("mysql", "priyanka"+":"+"Hani@2001"+"@tcp(localhost:3306)"+"/"+"petstore")
	if err != nil {
		return nil
	}

	return db
}

func main() { //nolint
	r := mux.NewRouter()
	db := DBConn()
	store := stores.New(db)
	service := services.New(store)
	h := handlers.New(service)
	r.HandleFunc("/pet", h.Post).Methods(http.MethodPost)
	r.HandleFunc("/pet/{id}", h.GetByID).Methods(http.MethodGet)
	http.Handle("/", r)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		return
	}
}
