package controllers

import (
	"fmt"
	"gps/models"
	"net/http"
)

type Users struct {
	Templates struct {
		New Template
	}
	UserService *models.UserService
}

func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	passwordConfirm := r.FormValue("password-confirm")
	if password != passwordConfirm {
		http.Error(w, "Password and Password confirm do not match", http.StatusBadRequest)
		return
	}
	newUser, err := u.UserService.Create(&models.NewUser{
		Email:    email,
		Password: password,
	})
	if err != nil {
		http.Error(w, fmt.Sprintln("Something wrong has happened"), http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "New User created: %+v", newUser)
}
