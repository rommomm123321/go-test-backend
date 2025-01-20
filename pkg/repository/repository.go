package repository

import (
	"github.com/jmoiron/sqlx"
	app "github.com/rommomm123321/go-test-backend"
)

type Authorization interface {
	CreateUser(user app.User) (app.User, error)
	GetUser(email, password string) (app.User, error)
}

type SportsHall interface {
}

type Repository struct {
	Authorization
	SportsHall
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
	}
}
