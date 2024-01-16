package main

import (
	"database/sql"
	"fmt"
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

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.sslmode)
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
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping() // test connection
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")
}