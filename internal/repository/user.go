package repository

import (
	"deep-ocean-api/pkg/verification"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

const maxVerificationCodeLength = 6

func (ur *UserRepository) Create(username, email, password string) error {
	authCode, err := verification.GenerateVerificationCode(maxVerificationCodeLength)

	if err != nil {
		return err
	}

	_, err = ur.db.Exec(
		"INSERT INTO users (username, email, password, auth_code) VALUES ($1, $2, $3, $4)",
		username, email, password, authCode)
	if err != nil {
		fmt.Printf("ERROR %v", err)
		return err
	}

	return nil
}
