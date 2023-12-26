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

	tpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))


	tpl, err = views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	r.NotFound(notFoundHandler)
	fmt.Println("Starting server on http://localhost:3000")
	err = http.ListenAndServe(":3000", r)
	if err != nil {
		panic(err)
	}
}
