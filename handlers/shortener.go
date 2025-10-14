package handlers

import (
	"URL_Shortener/utils"
	"encoding/json"
	"log"
	"net/http"
	"text/template"
)

func Make_Short(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == http.MethodGet {
		w.Header().Set("Content-Type", "text/html")
		tmpl, err := template.ParseFiles("./templates/shortener_form.html")
		if err != nil {
			log.Println("Template error:", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		full := r.FormValue("full")
		if full == "" {
			http.Error(w, "No URL provided", http.StatusBadRequest)
			return
		}

		short, err := utils.Generate_Short(full)
		if err != nil {
			log.Println(err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"link": short,
		})
	}
}
