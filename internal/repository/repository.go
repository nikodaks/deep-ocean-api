package repository

import "github.com/jmoiron/sqlx"

type User interface {
	Create(username, email, password string) error
}

type Repositories struct {
	User User
}

func NewRepositories(db *sqlx.DB) *Repositories {

	userRepository := &UserRepository{
		db: db,
	}

	repositories := &Repositories{
		User: userRepository,
	}

	return repositories
}
