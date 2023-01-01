package repository

import (
	"math/rand"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

const authCodeLimit = 100_000

func (ur *UserRepository) Create(username, email string) error {
	authCode := rand.Intn(authCodeLimit)
	_, err := ur.db.Exec("INSERT INTO users (username, email, auth_code) VALUES ($1, $2, $3)", username, email, authCode)
	if err != nil {
		return err
	}

	return nil
}
