package sleep

import "time"

type Sleep struct{
	ID         int       `json:"id"`
	User_id    int       `json:"user_id"`
	SleepStart time.Time `json:"sleep_start"`
	SleepEnd   time.Time `json:"sleep_end"`
	Duration   float64 `json:"duration_hours"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type SleepInput struct {
	User_id 			int 						`json:"user_id" binding:"required"`
	SleepStart string 			`json:"sleep_start" binding:"required"`
	SleepEnd   string 			`json:"sleep_end" binding:"required"`
}
