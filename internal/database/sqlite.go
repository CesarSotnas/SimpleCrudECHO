package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var DB *sql.DB

func InitSQLite() {
	var err error
	DB, err = sql.Open("sqlite3", "./internal/database/app.db")
	if err != nil {
		log.Fatalf("Error while connecting to SQLITE: %v", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatalf("Error while pinging database: %v", err)
	}

	log.Println("Connection Successful")
}
