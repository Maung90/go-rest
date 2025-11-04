package dailyStory

import "time"

type DailyStory struct{
	ID         int       `json:"id"`
	User_id    int       `json:"user_id"`
	StoryDate 	time.Time `json:"story_date"`
	StoryText  string 			`json:"story_text"`
	Mood   				string 			`json:"mood"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type StoryInput struct{
	StoryDate 	string `json:"story_date" binding:"required,datetime=2006-01-02"`
	StoryText  string 			`json:"story_text" binding:"required,min=15"`
	Mood   				string 			`json:"mood" binding:"required,oneof=happy angry excited neutral sad"`
}

