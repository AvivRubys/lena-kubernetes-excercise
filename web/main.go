package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	tmpl, err := template.ParseGlob("pages/*.html")
	if err != nil {
		log.Fatalln(err)
	}

	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.ExecuteTemplate(w, "home.html", nil)
	})
	r.Get("/welcome", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "arya"
		}

		tmpl.ExecuteTemplate(w, "welcome.html", struct{ Name string }{name})
	})

	log.Printf("Listening...")
	log.Fatalln(http.ListenAndServe(":8080", r))
}
