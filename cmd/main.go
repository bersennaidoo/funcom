package main

import (
	"log"

	"github.com/bersennaidoo/funcom/physical/config"
	"github.com/bersennaidoo/funcom/physical/dbconn"
)

func main() {
	log.Println("Initializing configuration")
	config := config.InitConfig(config.GetConfigFileName())

	log.Println("Initializing database")
	_ = dbconn.InitDatabase(config)
	log.Println("Database Initialized")
}
