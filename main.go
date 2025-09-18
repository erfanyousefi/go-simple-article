package main

import (
	"log"

	"github.com/erfanyousefi/simple-article/config"
	"github.com/erfanyousefi/simple-article/database"
)

func main() {
	cfg := config.LoadConfig()

	database.ConnectGORM(cfg)

	log.Println("all connections are true")
}
