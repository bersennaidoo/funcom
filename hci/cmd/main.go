package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "hci/userform.html")
	})
	log.Printf("Started on :9000")
	http.ListenAndServe(":9000", nil)

}
