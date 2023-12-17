package handlers

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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

func (usrh *UsersHandler) UserCreate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
	NewUser := models.User{}
	NewUser.Name = r.FormValue("user")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")

	f, _, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err.Error())
	}
	fileData, _ := ioutil.ReadAll(f)
	fileString := base64.StdEncoding.EncodeToString(fileData)

	output, err := json.Marshal(NewUser)
	fmt.Println(string(output))
	if err != nil {
		fmt.Println("Something went wrong!")
	}

	Response := CreateResponse{}

	q, err := usrh.usersRepository.UserCreate(NewUser, fileString)
	if err != nil {
		errorMessage, errorCode := dbErrorParse(err.Error())
		fmt.Println(errorMessage)
		error, httpCode, msg := ErrorMessages(errorCode)
		Response.Error = msg
		Response.ErrorCode = error
		http.Error(w, "Conflict", httpCode)
	}
	fmt.Println(q)
	createOutput, _ := json.Marshal(Response)
	fmt.Fprintln(w, string(createOutput))
}

func (usrh *UsersHandler) UsersRetrieve(w http.ResponseWriter, r *http.Request) {
	log.Println("starting retrieval")
	start := 0
	limit := 10

	next := start + limit

	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Link", "<http://localhost:3000/api/users?start="+string(next)+"; rel=\"next\"")

	users, err := usrh.usersRepository.UsersRetrieve()
	if err != nil {
		fmt.Fprintf(w, "No such user")
	}

	output, _ := json.Marshal(users)
	fmt.Fprintf(w, string(output))
}
