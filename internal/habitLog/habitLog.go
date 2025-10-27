package habitLog

import "time"

type HabitLog struct{
	ID         int       `json:"id"`
	Habit_id   int       `json:"habit_id"`
	User_id    int       `json:"user_id"`
	Status 				string 		 `json:"status"`
	Log_date   time.Time `json:"log_date"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

type CreateHabitLogInput struct {
	Habit_id int       `json:"habit_id" binding:"required"`
	User_id  int       `json:"user_id"  binding:"required"`
	LogDate  string    `json:"log_date" binding:"required"`
}
 