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

	r.Get("/", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "home.html", "tailwind.html"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.html",  "tailwind.html"))))

	r.Get("/faq", controllers.FAC(
		views.Must(views.ParseFS(templates.FS, "faq.html", "tailwind.html"))))

	usersCtrl := controllers.Users{}
	usersCtrl.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.html", "tailwind.html"))
	r.Get("/signup", usersCtrl.New)
	
	r.NotFound(notFoundHandler)

	fmt.Println("Starting server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
