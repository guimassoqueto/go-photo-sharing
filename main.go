package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/html"
)

func executeTemplateMinified(w http.ResponseWriter, tmplPath string, data interface{}) {
    path := filepath.Join(tmplPath)
    tmpl, err := template.ParseFiles(path)
    if err != nil {
        log.Printf("parsing template failed: %v", err)
        http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
        return
    }

    m := minify.New()
    m.AddFunc("text/html", html.Minify)

    var buf bytes.Buffer
    err = tmpl.Execute(&buf, data)
    if err != nil {
        log.Printf("executing template failed: %v", err)
        http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    minifiedHTML, err := m.String("text/html", buf.String())
    if err != nil {
        log.Printf("minifying template failed: %v", err)
        http.Error(w, "There was an error minifying the template", http.StatusInternalServerError)
        return
    }
    w.Write([]byte(minifiedHTML))
}

func executeTemplate(w http.ResponseWriter, tmplPath string, data interface{}) {
	path := filepath.Join(tmplPath)
	tmpl, err := template.ParseFiles(path)
	if err != nil {
		log.Printf("parsing template failed: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template failed: %v", err)
		http.Error(w, "There was an error executing the temmplate", http.StatusInternalServerError)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplateMinified(w, filepath.Join("templates", "home.gohtml"), nil)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplateMinified(w, filepath.Join("templates", "contact.gohtml"), nil)
}

type Faq struct{ Question string; Answer string }
func faqHandler(w http.ResponseWriter, r *http.Request) {
	data := []Faq{
		{"What is your name?", "My name is Go."},
		{"What is your quest?", "To learn Go."},
		{"What is your favorite color?", "Blue."},
		{"What is the airspeed velocity of an unladen swallow?", "African or European?"},
	}
	executeTemplateMinified(w, filepath.Join("templates", "faq.gohtml"), struct{Data []Faq}{Data: data})
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
