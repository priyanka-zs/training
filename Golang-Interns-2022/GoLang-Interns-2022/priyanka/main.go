package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/zopsmart/GoLang-Interns-2022/handlers"
	"github.com/zopsmart/GoLang-Interns-2022/middleware"
	"github.com/zopsmart/GoLang-Interns-2022/services"
	"github.com/zopsmart/GoLang-Interns-2022/store/car"
	"github.com/zopsmart/GoLang-Interns-2022/store/engine"
)

func DBConn() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "priyanka"
	dbPass := "Hani@2001"
	dbName := "cardealer"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)"+"/"+dbName)

	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	r := mux.NewRouter()

	db, err := DBConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	carStore := car.New(db)
	engineStore := engine.New(db)
	service := services.New(carStore, engineStore)
	h := handlers.New(service)
	r.HandleFunc("/car", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/car/{id}", h.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/car/{id}", h.Delete).Methods(http.MethodDelete)
	r.HandleFunc("/car", h.GetByBrand).Methods(http.MethodGet)
	r.HandleFunc("/car/{id}", h.Update).Methods(http.MethodPut)

	r.Use(middleware.Middleware)
	http.Handle("/", r)
	log.Println("Listen at 8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}
