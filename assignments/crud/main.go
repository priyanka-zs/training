package main

import (
	"assignments/crud/handlerscrud"
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DbConn() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "priyanka"
	dbPass := "Hani@2001"
	dbName := "cardealer"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)"+"/"+dbName)
	if err != nil {
		//log.Fatal(fmt.Errorf("unexpected error %v", err.Error()))
		return nil, err
	}
	return db, nil
}

func main() {
	r := mux.NewRouter()
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := handlerscrud.New(db)
	r.HandleFunc("/car", d.Post).Methods(http.MethodPost)
	r.HandleFunc("/car/{Id}", d.GetById).Methods(http.MethodGet)
	r.HandleFunc("/car/{Id}", d.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/car", d.Get).Methods(http.MethodGet)
	r.HandleFunc("/car/{Id}", d.Update).Methods(http.MethodPut)
	http.Handle("/", r)
	log.Println("Listen at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
