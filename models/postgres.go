package models

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

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
		sslmode:  "disable",
	}
}

// Open opens a connection to a PostgreSQL server.
// Whenever this function is called, it will open a new connection to the server,
// so it's recommended to call it only once and store the returned pointer in a variable,
// closing it with defer db.Close().
func Open(cfg PostgresConfig) (*sql.DB, error) {
	return sql.Open("pgx", cfg.String())
}

// String returns a string representation of the PostgresConfig struct.
func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.sslmode)
}
