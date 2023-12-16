package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bersennaidoo/funcom/domain/models"
	"github.com/bersennaidoo/funcom/infrastructure/repositories"
	"github.com/gorilla/mux"
)

type UsersHandler struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersHandler(usersRepository *repositories.UsersRepository) *UsersHandler {
	return &UsersHandler{
		usersRepository: usersRepository,
	}
}

func (usrh *UsersHandler) UserCreate(w http.ResponseWriter, r *http.Request) {

	NewUser := models.User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")
	output, err := json.Marshal(NewUser)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	q, err := usrh.usersRepository.UserCreate(NewUser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}

func (usrh *UsersHandler) UsersRetrieve(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Pragma", "no-cache")

	urlParams := mux.Vars(r)
	id := urlParams["id"]

	user, err := usrh.usersRepository.UsersRetrieve(id)
	if err != nil {
		fmt.Printf("%w", err)
	}

	output, _ := json.Marshal(user)
	fmt.Fprintf(w, string(output))
}
