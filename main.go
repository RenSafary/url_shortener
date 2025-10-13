package main

import (
	"log"
	"net/http"

	"URL_Shortener/db"
	"URL_Shortener/redis"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	database, err := db.Conn()
	if err != nil {
		log.Println("here")
		log.Fatal(err)
	}
	defer database.DB.Close()

	if err := redis.RDB(database); err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter()

	//r.HandleFunc("/main").Methods("POST")

	log.Println("Server is listening on port :8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}
}
