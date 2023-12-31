package controllers

import (
	"gps/views"
	"net/http"
)

type Users struct{
	Templates struct{
		New views.Template
	}
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.ExecuteTemplateMinified(w, nil)
}