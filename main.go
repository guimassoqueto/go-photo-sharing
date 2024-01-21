package main

import (
	"fmt"
	"gps/controllers"
	"gps/models"
	"gps/templates"
	"gps/utils"
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
		views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAC(
		views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	pgConfig := models.DefaultPostgresConfig()
	db, err := models.Open(pgConfig)
	utils.Panic(err)
	defer db.Close()
	userService := models.UserService{
		DB: db,
	}
	usersCtrl := controllers.Users{
		UserService: &userService, // TODO: Inject UserService here
	}
	usersCtrl.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersCtrl.New)
	r.Post("/users", usersCtrl.Create)

	r.NotFound(notFoundHandler)

	fmt.Println("Starting server on http://localhost:3000")
	http.ListenAndServe(":3000", r)
}
