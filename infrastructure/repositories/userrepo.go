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

func (urp *UsersRepository) UserCreate(newUser models.User) error {

	_, err := urp.dbclient.Exec("INSERT INTO funcom.users(user_name, user_first, user_last, user_email) VALUES(?, ?, ?, ?)",
		newUser.Name, newUser.First, newUser.Last, newUser.Email)
	if err != nil {
		return err
	}

	return nil
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

func (urp *UsersRepository) UsersUpdate(uid, email string) (int, error) {

	var userCount int
	err := urp.dbclient.QueryRow("SELECT count(user_id) from users where user_id=?",
		uid).Scan(&userCount)
	if err != nil {
		return userCount, nil
	}

	_, err = urp.dbclient.Exec("UPDATE users set user_email=? where user_id=?",
		email, uid)
	if err != nil {
		return userCount, err
	}

	return userCount, nil
}
