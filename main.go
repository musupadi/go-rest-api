package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/musupadi/go-rest-api/handlers"
)

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_inventory")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	router.HandleFunc("/login", handlers.Login).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
