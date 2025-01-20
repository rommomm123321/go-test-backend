package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	app "github.com/rommomm123321/go-test-backend"
)

type SportsHallRepository struct {
	db *sqlx.DB
}

func NewSportsHallRepository(db *sqlx.DB) *SportsHallRepository {
	return &SportsHallRepository{db: db}
}

func (r *SportsHallRepository) Create(userId int, data app.SportsHall) (app.SportsHall, error) {
	var sportsHall app.SportsHall

	createSportHall := fmt.Sprintf("INSERT INTO %s (name, description, image_url, owner_id) values ($1, $2, $3, $4) RETURNING id, name, description, owner_id, image_url, created_at, updated_at", sportsHallsTable)
	row := r.db.QueryRow(createSportHall, data.Name, data.Description, data.Image_url, userId)
	err := row.Scan(
		&sportsHall.Id,
		&sportsHall.Name,
		&sportsHall.Description,
		&sportsHall.OwnerId,
		&sportsHall.Image_url,
		&sportsHall.CreatedAt,
		&sportsHall.UpdatedAt,
	)
	if err != nil {
		return app.SportsHall{}, err
	}
	return sportsHall, nil
}

func (r *SportsHallRepository) GetAll(userId int) ([]app.SportsHall, error) {
	var sportsHallData []app.SportsHall
	query := fmt.Sprintf("SELECT id, name, description, image_url, owner_id, created_at, updated_at FROM %s WHERE owner_id = $1", sportsHallsTable)

	err := r.db.Select(&sportsHallData, query, userId)
	if err != nil {
		return nil, err
	}
	return sportsHallData, nil

}
