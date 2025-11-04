package dailyStory

import (
	"database/sql"
	"time"
)


type Repository interface {
	FindAll(userID int) ([]DailyStory, error)
	FindByID(userID, id int) (DailyStory, error)
	FindByDate(userID int, date string) ([]DailyStory, error)
	Save(dailyStory DailyStory) (DailyStory, error)
	Update(dailyStory DailyStory) (DailyStory, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}


func (r *repository) FindAll(userID int) ([]DailyStory, error) {
	var dailyStorys []DailyStory

	rows, err := r.db.Query("SELECT id, user_id, story_date, story_text, mood, created_at, updated_at FROM daily_stories WHERE user_id = ?", userID)
	if err != nil {
		return dailyStorys, err
	}
	defer rows.Close()

	for rows.Next() {
		var dailyStory DailyStory
		err := rows.Scan(&dailyStory.ID, &dailyStory.User_id, &dailyStory.StoryDate, &dailyStory.StoryText, &dailyStory.Mood, &dailyStory.CreatedAt, &dailyStory.UpdatedAt)
		if err != nil {
			return dailyStorys, err
		}
		dailyStorys = append(dailyStorys, dailyStory)
	}
	return dailyStorys, nil
}

func (r *repository) FindByID(userID, id int) (DailyStory, error) {
	var dailyStory DailyStory
	err := r.db.QueryRow("SELECT id, user_id, story_date, story_text, mood, created_at, updated_at FROM daily_stories WHERE id = ? AND user_id = ?", id, userID).
	Scan(&dailyStory.ID, &dailyStory.User_id, &dailyStory.StoryDate, &dailyStory.StoryText, &dailyStory.Mood, &dailyStory.CreatedAt, &dailyStory.UpdatedAt)
	if err != nil {
		return dailyStory, err
	}
	return dailyStory, nil
}

func(r * repository) FindByDate(userID int, date string) ([]DailyStory, error){
	var dailyStorys []DailyStory

	rows, err := r.db.Query("SELECT id, user_id, story_date, story_text, mood, created_at, updated_at FROM daily_stories WHERE story_date = ? AND user_id = ?", date, userID)
	if err != nil {
		return dailyStorys, err
	}
	defer rows.Close()

	for rows.Next() {
		var dailyStory DailyStory
		err := rows.Scan(&dailyStory.ID, &dailyStory.User_id, &dailyStory.StoryDate, &dailyStory.StoryText, &dailyStory.Mood, &dailyStory.CreatedAt, &dailyStory.UpdatedAt)
		if err != nil {
			return dailyStorys, err
		}
		dailyStorys = append(dailyStorys, dailyStory)
	}
	return dailyStorys, nil
}


func (r *repository) Save(dailyStory DailyStory) (DailyStory, error) {

	query := "INSERT INTO daily_stories (user_id, story_date, story_text, mood, created_at,  updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(query, dailyStory.User_id, dailyStory.StoryDate, dailyStory.StoryText, dailyStory.Mood, time.Now(), time.Now())
	if err != nil {
		return dailyStory, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return dailyStory, err
	}

	return r.FindByID(dailyStory.User_id, int(lastID))
}


func (r *repository) Update(dailyStory DailyStory) (DailyStory, error) {
	query := `
		UPDATE daily_stories 
		SET story_date = ?, story_text = ?, mood = ?, updated_at = ? 
		WHERE id = ? AND user_id = ?`

	storyDateString := dailyStory.StoryDate.Format("2006-01-02")

	result, err := r.db.Exec(query,
		storyDateString,
		dailyStory.StoryText,
		dailyStory.Mood,
		time.Now(),
		dailyStory.ID,
		dailyStory.User_id,
	)
	if err != nil {
		return DailyStory{}, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return DailyStory{}, err
	}

	if rowsAffected == 0 {
		return DailyStory{}, sql.ErrNoRows
	}

	return r.FindByID( dailyStory.User_id, dailyStory.ID)
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM daily_stories WHERE id = ?", id)
	return err
}
