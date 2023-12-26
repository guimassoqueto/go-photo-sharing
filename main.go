package main

import (
	"fmt"
	"gps/controllers"
	"gps/templates"
	"gps/views"
	"net/http"

	"github.com/go-chi/chi/v5"
)


func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	message := "<h1>Page not found</h1>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, message)
}

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticMinifiedHandler(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "layout-parts.gohtml", "home.gohtml"))))

	r.Get("/contact", controllers.StaticMinifiedHandler(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "layout-parts.gohtml", "contact.gohtml"))))

	r.Get("/faq", controllers.FAC(
		views.Must(views.ParseFS(templates.FS, "layout-page.gohtml", "layout-parts.gohtml", "faq.gohtml"))))

	r.NotFound(notFoundHandler)

	fmt.Println("Starting server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
