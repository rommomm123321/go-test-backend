package app

type MembershipType struct {
	Id           int     `json:"id"`
	Name         string  `json:"name"`
	DisciplineId int     `json:"discipline_id"`
	Price        float64 `json:"price"`
	Duration     int     `json:"duration"`
	Limited      bool    `json:"limited"`
	CreatedAt    string  `json:"created_at"`
	UpdatedAt    string  `json:"updated_at"`
}
