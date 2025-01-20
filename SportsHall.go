package app

type SportsHall struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" db:"name" binding:"required"`
	OwnerId     int    `json:"owner_id" db:"owner_id"`
	Description string `json:"description" db:"description" binding:"required"`
	Image_url   string `json:"image_url" db:"image_url"`
	CreatedAt   string `json:"created_at" db:"created_at"`
	UpdatedAt   string `json:"updated_at" db:"updated_at"`
}
