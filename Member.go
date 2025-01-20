package app

type Member struct {
	Id               int    `json:"id"`
	Name             string `json:"name"`
	PhoneNumber      string `json:"phone_number"`
	SportsHallId     int    `json:"sports_hall_id"`
	DisciplineId     int    `json:"discipline_id"`
	MembershipTypeId int    `json:"membership_type_id"`
	StartDate        string `json:"start_date"`
	EndDate          string `json:"end_date"`
	Description      string `json:"description"`
	Image_url        string `json:"image_url"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}
