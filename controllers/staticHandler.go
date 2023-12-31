package controllers

import (
	"net/http"
)

func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAC(tpl Template) http.HandlerFunc {
	data := []struct{
		Question string
		Answer string
	} {
		{Question: "What is the difference between the free and paid versions?", Answer: "The free version is limited to 10,000 requests per month. The paid version has no limit."},
		{Question: "How do I get an API key?", Answer: "You can get an API key by signing up for a free account. You will be able to use the API immediately."},
		{Question: "How do I use the API?", Answer: "You can use the API by sending HTTP requests to the API endpoint. You can find more information in the documentation."},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, data)
	}
}