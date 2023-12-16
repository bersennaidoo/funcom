package repositories

import (
	"database/sql"
	"fmt"

	"github.com/bersennaidoo/funcom/domain/models"
)

type UsersRepository struct {
	dbclient *sql.DB
}

func NewUsersRepository(dbclient *sql.DB) *UsersRepository {
	return &UsersRepository{
		dbclient: dbclient,
	}
}

func (urp *UsersRepository) UserCreate(newUser models.User) (sql.Result, error) {

	query := "INSERT INTO users set user_nickname='" + newUser.Name +
		"', user_first='" + newUser.First + "', user_last='" + newUser.Last +
		"', user_email='" + newUser.Email + "'"

	q, err := urp.dbclient.Exec(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return q, nil
}
