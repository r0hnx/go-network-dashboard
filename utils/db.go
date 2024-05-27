package utils

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func setupDatabase() (*sql.DB, error) {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil
}
