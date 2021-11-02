package repository

import (
	"database/sql"
	"github.com/bdemirpolat/integration-test/models"
)

type UserRepository interface {
	Create(user models.User) error
}

type UserRepo struct {
	DB *sql.DB
}

func (u UserRepo) Create(user models.User) error {
	stmt, err := u.DB.Prepare("INSERT INTO users (username) VALUES (?);")
	if err != nil {
		return err
	}
	_,err = stmt.Exec(user.Username)
	return err
}

