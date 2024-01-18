package db

import (
	"database/sql"
)

// UserService is a service for managing users
type UserService struct {
	DB *sql.DB
}

// // User represents a user in the system
// type UserService2 interface {
// 	Create(name, email, password string) (*models.User, error)
// }


func (us *UserService) Create(name, email, password string) (*User, error) {
}