package activity

import "time"

type Activity struct{
	ID         					int       `json:"id"`
	User_id    					int       `json:"user_id"`
	ActivityDate 			time.Time `json:"activity_date"`
	Title  									string 			`json:"story_text"`
	DurationMinutes string 			`json:"mood"`
	CreatedAt  					time.Time `json:"created_at"`
	UpdatedAt  					time.Time `json:"updated_at"`
}

