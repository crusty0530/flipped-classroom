package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	url := os.Getenv("DATABASE_URL")

	if url == "" {
		log.Fatal("DATABASE_URL is not set")
	}

	db, err := sql.Open("postgres", url)

	if err != nil {
		log.Fatal("Failed to open database connection: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	log.Println("Connected to database")
	return db
}
