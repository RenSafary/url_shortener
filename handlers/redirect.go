package handlers

import (
	"log"
	"net/http"

	"URL_Shortener/db"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func Redirect(d *db.DBshort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		short := vars["link"]

		full, err := d.Get(short)
		if err != nil {
			log.Println("Couldn't get full link:", err)
			http.Redirect(w, r, "/error", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, full, http.StatusMovedPermanently)
	}
}
