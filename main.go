package main

import (
	"fmt"
	"net/http"
)

func handleFunc(w http.ResponseWriter, r *http.Request) {
	message := "<h1>Welcome to my awesome website!</h1>"
	fmt.Fprint(w, message)
}

func main() {
	http.HandleFunc("/", handleFunc)
	fmt.Println("Starting server on :3000...")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		panic(err)
	}
}
