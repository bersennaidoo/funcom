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

func (urp *UsersRepository) UserCreate(newUser models.User) (sql.Result, error) {

	query := "INSERT INTO users set user_nickname='" + newUser.Name +
		"', user_first='" + newUser.First + "', user_last='" + newUser.Last +
		"', user_email='" + newUser.Email + "'"

	q, err := urp.dbclient.Exec(query)
	if err != nil {
		return nil, err
	}

	return q, nil
}

func (urp *UsersRepository) UsersRetrieve(id string) (*models.User, error) {

	ReadUser := models.User{}
	err := urp.dbclient.QueryRow("select * from users where user_id=?", id).Scan(&ReadUser.ID,
		&ReadUser.Name, &ReadUser.First, &ReadUser.Last, &ReadUser.Email)
	switch {
	case err == sql.ErrNoRows:
		return nil, err
	case err != nil:
		return nil, err
	}

	return &ReadUser, nil
}
