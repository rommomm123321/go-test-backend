package service

import (
	app "github.com/rommomm123321/go-test-backend"
	"github.com/rommomm123321/go-test-backend/pkg/repository"
)

type Authorization interface {
	CreateUser(user app.User) (app.User, error)
	GenerateToken(email, password string) (string, error)
	ParseToken(token string) (int, error)
}

type SportsHall interface {
}

type Service struct {
	Authorization
	SportsHall
}

func NewService(r *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(r.Authorization),
	}
}
