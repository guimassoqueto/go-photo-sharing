package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	message := "<h1>Welcome to my awesome website!</h1>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, message)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	message := "<h1>Contact Page</h1>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, message)
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	message := "<h1>Page not found</h1>"
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, message)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	message := `
<h1>FAQ</h1>
<h2>Is there a free version?</h2>
<p>Yes, now go fuck yourself slowly.</h2>
<h2>How do I pay?</h2>
<p>With money, bitch.</h2>
`
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
