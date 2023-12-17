package server

import (
	"log"
	"net/http"

	"github.com/bersennaidoo/funcom/application/rest/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config       *viper.Viper
	router       *mux.Router
	usersHandler *handlers.UsersHandler
}

func New(config *viper.Viper, usersHandler *handlers.UsersHandler) *HttpServer {
	return &HttpServer{
		config:       config,
		usersHandler: usersHandler,
	}
}

func (hts *HttpServer) InitRouter() {

	hts.router = mux.NewRouter()
	hts.router.HandleFunc("/api/users", hts.usersHandler.UserCreate).Methods("POST")
	hts.router.HandleFunc("/api/users", hts.usersHandler.UsersRetrieve).Methods("GET")
	http.Handle("/", hts.router)
}

func (hts *HttpServer) Start() {
	srvaddr := hts.config.GetString("http.server_address")

	log.Println("Server started on", srvaddr)
	err := http.ListenAndServe(srvaddr, nil)
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
