package app

type SportsHall struct {
	Id          int    `json:"id"`
	Name        string `json:"name" binding:"required"`
	OwnerId     int    `json:"owner_id"`
	Description string `json:"description"`
	Image_url   string `json:"image_url"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
