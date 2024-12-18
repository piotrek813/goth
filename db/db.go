package db

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func GetDB() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db, err := sql.Open("sqlite", os.Getenv("DB_URL"))

	if err != nil {
		panic(err)
	}

	return db
}
