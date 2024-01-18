package models

import "database/sql"

type User struct {
	ID	   uint64  `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type UserService struct {
	DB *sql.DB
}