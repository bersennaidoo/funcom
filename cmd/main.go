package main

import (
	"log"

	"github.com/bersennaidoo/funcom/application/rest/handlers"
	"github.com/bersennaidoo/funcom/application/rest/server"
	"github.com/bersennaidoo/funcom/infrastructure/repositories"
	"github.com/bersennaidoo/funcom/physical/config"
	"github.com/bersennaidoo/funcom/physical/dbconn"
)

func main() {
	log.Println("Initializing configuration")
	config := config.InitConfig(config.GetConfigFileName())

	log.Println("Initializing database")
	dbclient := dbconn.InitDatabase(config)

	urepo := repositories.NewUsersRepository(dbclient)
	uhandler := handlers.NewUsersHandler(urepo)

	srv := server.New(config, uhandler)

	srv.InitRouter()
	srv.Start()
}
