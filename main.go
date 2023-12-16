package main

import (
	"fmt"
	"net/http"
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

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		notFoundHandler(w, r)
	}
}


func main() {
	fmt.Println("Starting server on :3000...")
	err := http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
	if err != nil {
		panic(err)
	}

}
