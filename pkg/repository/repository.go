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
	Create(userId int, data app.SportsHall) (app.SportsHall, error)
	GetAll(userId int) ([]app.SportsHall, error)
}

type Repository struct {
	Authorization
	SportsHall
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthRepository(db),
		SportsHall:    NewSportsHallRepository(db),
	}
}
