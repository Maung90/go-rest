package activity

import "time"

type Activity struct{
	ID         					int       `json:"id"`
	User_id    					int       `json:"user_id"`
	ActivityDate 			time.Time `json:"activity_date"`
	Title  									string 			`json:"title"`
	DurationMinutes int 						`json:"duration_minutes"`
	Notes										 string 			`json:"notes"`
	CreatedAt  					time.Time `json:"created_at"`
	UpdatedAt  					time.Time `json:"updated_at"`
}

type ActivityInput struct{
	User_id    					int       `json:"user_id" binding:"required,numeric"`
	ActivityDate 			string 		 `json:"activity_date" binding:"required,datetime=2006-01-02"`
	Title  									string 			`json:"story_text" binding:"required"`
	DurationMinutes int 						`json:"duration_minutes" binding:"required,numeric"`
	Notes										 string 			`json:"notes"`
}