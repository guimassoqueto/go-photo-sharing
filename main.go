package main

import (
	"fmt"
	"gps/views"
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

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "home.gohtml"), nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "contact.gohtml"), nil)
}

type Faq struct{ Question string; Answer string }
func faqHandler(w http.ResponseWriter, r *http.Request) {
	data := []Faq{
		{"What is your name?", "My name is Go."},
		{"What is your quest?", "To learn Go."},
		{"What is your favorite color?", "Blue."},
		{"What is the airspeed velocity of an unladen swallow?", "African or European?"},
	}
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"), struct{Data []Faq}{Data: data})
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	message := "<h1>Page not found</h1>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, message)
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(notFoundHandler)
	fmt.Println("Starting server on :3000...")
	err := http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
