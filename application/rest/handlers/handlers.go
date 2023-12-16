package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bersennaidoo/funcom/domain/models"
	"github.com/bersennaidoo/funcom/infrastructure/repositories"
)

type UsersHandler struct {
	usersRepository *repositories.UsersRepository
}

func NewUsersHandler(usersRepository *repositories.UsersRepository) *UsersHandler {
	return &UsersHandler{
		usersRepository: usersRepository,
	}
}

func (usr *UsersHandler) UserCreate(w http.ResponseWriter, r *http.Request) {

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

	q, err := usr.usersRepository.UserCreate(NewUser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(q)
}
