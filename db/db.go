package db

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectWithDB() *sql.DB {
	connection := "user=postgres dbname=postgres password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}
