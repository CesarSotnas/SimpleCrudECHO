package main

import (
	"GinEchoCrud/internal/database"
	"GinEchoCrud/internal/server"
)

func main() {
	database.InitSQLite()
	server.InitNewServer()
}
