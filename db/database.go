package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

func dbConnect() (db *sql.DB) {
	const (
		host = "localhost"
		port = 5432
	)
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return &db
}
