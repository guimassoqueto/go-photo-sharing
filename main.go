package main

import (
	"fmt"
	"gps/views"
	"gps/controllers"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)


func executeTemplate(w http.ResponseWriter, tmplPath string, data interface{}) {
	t, err := views.Parse(tmplPath)
	if err != nil {
		log.Printf("parsing template failed: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	t.ExecuteTemplateMinified(w, data)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	message := "<h1>Page not found</h1>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, message)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(notFoundHandler)

	fmt.Println("Starting server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
