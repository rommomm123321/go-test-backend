package app

type Discipline struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	SportsHallId int    `json:"sports_hall_id"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}
