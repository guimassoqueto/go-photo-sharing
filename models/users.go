package models

import (
	"database/sql"
	"fmt"
	"strings"

	"gps/utils"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system (password here is hashed)
type User struct {
	ID	   			string  `json:"id"`
	Email   		string `json:"email"`
	PasswordHash  	string `json:"password_hash"`
}

// NewUser represents a new user signing up (password here is plaintext)
type NewUser struct {
	Email  		string `json:"email"`
	Password 	string `json:"password"`
}

type UserService struct {
	DB *sql.DB
}

func (us *UserService) Create(nu *NewUser) (*User, error) {
	nu.Email = strings.ToLower(nu.Email)
	nu.Email = strings.TrimSpace(nu.Email)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("error hashing password: %w", err)
	}
	nu.Password = string(hashedPassword)

	uuidv4, err := utils.UUIDV4()
	if err != nil {
		return nil, fmt.Errorf("error generating uuidv4 for new user: %w", err)
	}

	row := us.DB.QueryRow("INSERT INTO users (id, email, password_hash) VALUES ($1, $2, $3) RETURNING *", uuidv4, nu.Email, nu.Password)
	if row.Err() != nil {
		return nil, fmt.Errorf("error inserting new user: %w", row.Err())
	}

	user := User{}
	err = row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("error scanning new user: %w", err)
	}

	return &user, nil
}