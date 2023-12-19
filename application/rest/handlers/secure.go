package handlers

import (
	"fmt"
	"log"
	"net/http"
)

const SSLprotocol = "https://"
const serverName = "localhost"
const port = ":8443"

func SecureRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You have arrived at port 8443, and now you are now marginally more secure.")
}

func RedirectNonSecure(w http.ResponseWriter, r *http.Request) {
	log.Println("Non-secure request initiated, redirecting.")
	fmt.Println(r.RequestURI)
	http.Redirect(w, r, "https://localhost"+r.RequestURI, http.StatusMovedPermanently)
}
