package main

import (
	"database/sql"
	"fmt"
	"gps/models"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	sslmode  string
}

// String returns a string representation of the PostgresConfig struct.
// The string is formatted as a PostgreSQL connection string.
//
// host: The host of the PostgreSQL server.
// port: The port of the PostgreSQL server.
// user: The user to connect to the PostgreSQL server.
// password: The password to connect to the PostgreSQL server.
// dbname: The name of the database to connect to.
// sslmode: The SSL mode to use when connecting to the server.
//
// Returns a string that can be used to connect to a PostgreSQL server.
func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.sslmode)
}

// Panic checks if the provided error is not nil and panics if it is.
// err: The error to check. If err is not nil, the function will panic.
// It's used to halt the execution of the program when a fatal error is encountered.
func Panic(err error) {
	if err != nil {
		panic(err)
	}
}

type User struct {
	ID   int
	Name string
}

func main() {
	cfg := PostgresConfig{
		Host:     	"localhost",
		Port:     	5432,
		User:    	"postgres",
		Password: 	"postgres",
		Database: 	"postgres",
		sslmode:  	"disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	Panic(err)
	defer db.Close()
	err = db.Ping() // test connection
	Panic(err)
	fmt.Println("Connected!")

	// create a test table
	// _, err = db.Exec(`
	// 	CREATE TABLE IF NOT EXISTS test (
	// 		id VARCHAR(36) PRIMARY KEY, email VARCHAR(255) UNIQUE NOT NULL, password_hash TEXT NOT NULL
	// 	);
	// `)
	// Panic(err)
	// fmt.Println("Table created!")

	nu := models.UserService{
		DB: db,
	}
	user, error := nu.Create(&models.NewUser{
		Email: "gaymen@gmail.com",
		Password: "password123",
	})
	Panic(error)
	fmt.Println(user)

	// name := "Guilherme"
	// row := db.QueryRow(`INSERT INTO test (name) VALUES ($1) RETURNING *;`, name)
	// Panic(row.Err())
	// var user User
	// err = row.Scan(&user.ID, &user.Name)
	// Panic(err)
	// fmt.Println(user.ID, user.Name)

	// id := 2
	// row := db.QueryRow(`SELECT * FROM test WHERE id = $1;`, id)
	// Panic(row.Err())
	// var user User
	// err = row.Scan(&user.ID, &user.Name)
	// Panic(err)
	// fmt.Println(user.ID, user.Name)

	// rows, err := db.Query(`SELECT * FROM test;`)
	// Panic(err)
	// var users []User
	// for rows.Next() {
	// 	var u User
	// 	err = rows.Scan(&u.ID, &u.Name)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		continue
	// 	}
	// 	users = append(users, u)
	// }
	// fmt.Println(users)
}