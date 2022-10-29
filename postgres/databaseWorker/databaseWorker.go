package databaseworker

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "group"
)

func OpenDbConnection(dbName string) *sql.DB {
	connectionString := fmt.Sprintf("user=%s password=%s host=%s port=%d sslmode=disable dbname=%s", user, password, host, port, dbName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseDbConnection(db *sql.DB) {
	err := db.Close()

	if err != nil {
		log.Fatal(err)
	}
}
