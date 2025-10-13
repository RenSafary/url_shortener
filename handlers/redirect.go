package handlers

import (
	"net/http"

	"URL_Shortener/db"
	_ "github.com/lib/pq"
)

func Redirect(d *db.DBshort) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
