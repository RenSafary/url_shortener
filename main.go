package main

import (
	"log"
	"net/http"

	"URL_Shortener/db"
	"URL_Shortener/handlers"
	"URL_Shortener/redis"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	database, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer database.DB.Close()

	rdb, err := redis.RDB(database)
	if err != nil {
		log.Fatal(err)
	}

	rdb.GetData()

	r := mux.NewRouter()

	r.HandleFunc("/make_short", handlers.Make_Short).Methods("GET", "POST")
	r.HandleFunc("/{link}", handlers.Redirect(database)).Methods("GET")

	log.Println("Server is listening on port :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
