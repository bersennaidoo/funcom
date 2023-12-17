package repositories

import (
	"database/sql"

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

func (urp *UsersRepository) UserCreate(newUser models.User, fileString string) (sql.Result, error) {

	query := "INSERT INTO users set user_image='" + fileString + "', user_nickname='" + newUser.Name +
		"', user_first='" + newUser.First + "', user_last='" + newUser.Last +
		"', user_email='" + newUser.Email + "'"

	q, err := urp.dbclient.Exec(query)
	if err != nil {
		return nil, err
	}

	return q, nil
}

func (urp *UsersRepository) UsersRetrieve() (*models.Users, error) {

	rows, err := urp.dbclient.Query("select * from users LIMIT 10")

	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	}

	Response := models.Users{}

	for rows.Next() {
		user := models.User{}
		rows.Scan(&user.ID, &user.Name, &user.First, &user.Last, &user.Email)
		Response.Users = append(Response.Users, user)
	}

	return &Response, nil
}
