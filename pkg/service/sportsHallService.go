package service

import (
	app "github.com/rommomm123321/go-test-backend"
	"github.com/rommomm123321/go-test-backend/pkg/repository"
)

type SportsHallService struct {
	repo repository.SportsHall
}

func NewSportsHallService(repo repository.SportsHall) *SportsHallService {
	return &SportsHallService{repo: repo}
}

func (s *SportsHallService) Create(userId int, data app.SportsHall) (app.SportsHall, error) {
	return s.repo.Create(userId, data)
}

func (s *SportsHallService) GetAll(userId int) ([]app.SportsHall, error) {
	return s.repo.GetAll(userId)
}
