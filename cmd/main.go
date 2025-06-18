package main

import (
	"GinEchoCrud/internal/database"
	"GinEchoCrud/internal/helpers"
	"GinEchoCrud/internal/server"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}
	port := os.Getenv("PORT")

	helpers.LoadJWTSecret()
	database.InitSQLite()
	server.InitNewServer(port)
}
