package main

import (
	"GinEchoCrud/internal/database"
	"GinEchoCrud/internal/server"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte("DEFAULT_PASSWORD"), bcrypt.DefaultCost)
	fmt.Println(string(hash))

	database.InitSQLite()
	server.InitNewServer()
}
