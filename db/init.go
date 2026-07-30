package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=postgres password=newpassword dbname=go_sqlc sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
