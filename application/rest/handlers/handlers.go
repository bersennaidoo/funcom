package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bersennaidoo/funcom/application/rest/specs"
	"github.com/bersennaidoo/funcom/domain/models"
	"github.com/bersennaidoo/funcom/foundation/errmsg"
	"github.com/bersennaidoo/funcom/foundation/formats"
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

func (h *UsersHandler) UsersInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Allow", "DELETE,GET,HEAD,OPTIONS,POST,PUT")

	UserDoc := []specs.DocMethod{}
	UserDoc = append(UserDoc, specs.UserPOST)
	UserDoc = append(UserDoc, specs.UserOPTIONS)
	output := formats.SetFormat(UserDoc)

	fmt.Fprintln(w, string(output))
}

func (h *UsersHandler) UserCreate(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:9000")
	NewUser := models.User{}
	NewUser.Name = r.FormValue("username")
	NewUser.Email = r.FormValue("email")
	NewUser.First = r.FormValue("first")
	NewUser.Last = r.FormValue("last")

	Response := CreateResponse{}

	err := h.usersRepository.UserCreate(NewUser)
	if err != nil {
		errorMessage, errorCode := errmsg.DbErrorParse(err.Error())
		fmt.Println(errorMessage)
		error, httpCode, msg := errmsg.ErrorMessages(errorCode)
		Response.Error = msg
		Response.ErrorCode = error
		http.Error(w, "Conflict", httpCode)
		createOutput, _ := json.Marshal(Response)
		fmt.Fprintln(w, string(createOutput))
	}

	fmt.Fprintln(w, "Create")
}

func (h *UsersHandler) UsersRetrieve(w http.ResponseWriter, r *http.Request) {
	log.Println("starting retrieval")
	start := 0
	limit := 10

	next := start + limit

	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Link", "<http://localhost:3000/api/users?start="+string(next)+"; rel=\"next\"")

	users, err := h.usersRepository.UsersRetrieve()
	if err != nil {
		fmt.Fprintf(w, "No such user")
	}

	output, _ := json.Marshal(users)
	fmt.Fprintf(w, string(output))
}

func (h *UsersHandler) UsersUpdate(w http.ResponseWriter, r *http.Request) {
	Response := UpdateResponse{}
	params := mux.Vars(r)
	uid := params["id"]
	email := r.FormValue("email")

	userCount, err := h.usersRepository.UsersUpdate(uid, email)
	if userCount == 0 {
		error, httpCode, msg := errmsg.ErrorMessages(404)
		log.Println(error)
		log.Println(w, msg, httpCode)
		Response.Error = msg
		Response.ErrorCode = httpCode
		http.Error(w, msg, httpCode)
		return
	}

	if err != nil {
		_, errorCode := errmsg.DbErrorParse(err.Error())
		_, httpCode, msg := errmsg.ErrorMessages(errorCode)
		Response.Error = msg
		Response.ErrorCode = httpCode
		http.Error(w, msg, httpCode)
		return
	}

	Response.Error = "success"
	Response.ErrorCode = 0
	output := formats.SetFormat(Response)
	fmt.Fprintln(w, string(output))
}
