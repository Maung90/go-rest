package habit

import "time"

type Habit struct{
	ID        int       `json:"id"`
	User_id   int       `json:"user_id"`
	Title  	 string    `json:"title"`
	Description     string    `json:"description"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateHabitInput struct {
	User_id int `json:"user_id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Description  string    `json:"description"`
}

type UpdateHabitInput struct {
	User_id int `json:"user_id" binding:"required"`
	Title string `json:"title" binding:"required"`
	Description string   `json:"description"`
}