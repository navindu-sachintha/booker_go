package config

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
)

var (
	db *sql.DB
)

func init() {
	connect()
}

func connect() {
	godotenv.Load()

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL not found in .env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Can't connect to database", err)
	}

	db = conn
}

func GetDB() *sql.DB {
	return db
}
