package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

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
	hts.router.HandleFunc("/api/users/{id:[0-9]+}", hts.usersHandler.UsersUpdate).Methods("PUT")
	hts.router.HandleFunc("/api/users", hts.usersHandler.UsersInfo).Methods("OPTIONS")
	http.Handle("/", hts.router)
}

func (hts *HttpServer) Start() {
	wg := sync.WaitGroup{}
	httpaddr := hts.config.GetString("http.http_addr")
	httpsaddr := hts.config.GetString("https.https_addr")
	fmt.Println(httpsaddr)
	fmt.Println(httpaddr)

	wg.Add(1)
	go func() {
		err := http.ListenAndServe(httpaddr, http.HandlerFunc(handlers.RedirectNonSecure))
		log.Fatal(err)
		wg.Done()
	}()
	wg.Add(1)
	go func() {
		err := http.ListenAndServeTLS(httpsaddr, "./documentation/certs/server.crt", "./documentation/certs/server.key", http.HandlerFunc(handlers.SecureRequest))
		log.Fatal(err)
		wg.Done()
	}()

	wg.Wait()

}
