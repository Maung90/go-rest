package activity

import "time"

type Activity struct{
	ID        int       `json:"id"`
	User_id   int       `json:"user_id"`
	Activity  string    `json:"activity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Struct untuk menampung data input dari user
type CreateActivityInput struct {
	User_id int `json:"user_id" binding:"required"`
	Activity string `json:"activity" binding:"required"`
}

type UpdateActivityInput struct {
	User_id int `json:"user_id"`
	Activity string `json:"activity"`
}