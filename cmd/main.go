package main

import (
	"GinEchoCrud/internal/database"
	"GinEchoCrud/internal/server"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	database.InitSQLite()
	server.InitNewServer()
}
