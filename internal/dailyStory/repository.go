package dailyStory

import (
	"database/sql"
	"time"

	sqlbuilder "go-rest/pkg/sql"
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
	builder := sqlbuilder.NewSQLBuilder(r.db, "daily_stories").
		Select("user_id", "story_date", "story_text", "mood").
		Where("user_id = ?", userID).
		OrderBy("story_date DESC")

	rows, err := builder.Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []DailyStory
	for rows.Next() {
		var s DailyStory
		err := rows.Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt, &s.User_id, &s.StoryDate, &s.StoryText, &s.Mood)
		if err != nil {
			return nil, err
		}
		stories = append(stories, s)
	}
	return stories, nil
}

func (r *repository) FindByID(userID, id int) (DailyStory, error) {
	var s DailyStory

	rows, err := sqlbuilder.RawQuery(r.db, `
		SELECT id, user_id, story_date, story_text, mood, created_at, updated_at
		FROM daily_stories 
		WHERE id = ? AND user_id = ?
	`, id, userID)
	if err != nil {
		return s, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&s.ID, &s.User_id, &s.StoryDate, &s.StoryText, &s.Mood, &s.CreatedAt, &s.UpdatedAt)
		if err != nil {
			return s, err
		}
	}
	return s, nil
}

func (r *repository) FindByDate(userID int, date string) ([]DailyStory, error) {
	builder := sqlbuilder.NewSQLBuilder(r.db, "daily_stories").
		Select("user_id", "story_date", "story_text", "mood").
		Where("user_id = ?", userID).
		Where("story_date = ?", date).
		OrderBy("created_at DESC")

	rows, err := builder.Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stories []DailyStory
	for rows.Next() {
		var s DailyStory
		err := rows.Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt, &s.User_id, &s.StoryDate, &s.StoryText, &s.Mood)
		if err != nil {
			return nil, err
		}
		stories = append(stories, s)
	}
	return stories, nil
}

func (r *repository) Save(story DailyStory) (DailyStory, error) {
	data := map[string]interface{}{
		"user_id":    story.User_id,
		"story_date": story.StoryDate,
		"story_text": story.StoryText,
		"mood":       story.Mood,
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	result, err := sqlbuilder.Insert(r.db, "daily_stories", data)
	if err != nil {
		return story, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return story, err
	}

	return r.FindByID(story.User_id, int(lastID))
}

func (r *repository) Update(story DailyStory) (DailyStory, error) {
	data := map[string]interface{}{
		"story_date": story.StoryDate.Format("2006-01-02"),
		"story_text": story.StoryText,
		"mood":       story.Mood,
		"updated_at": time.Now(),
	}

	_, err := sqlbuilder.Update(
		r.db,
		"daily_stories",
		data,
		"id = ? AND user_id = ?",
		story.ID, story.User_id,
	)
	if err != nil {
		return DailyStory{}, err
	}

	return r.FindByID(story.User_id, story.ID)
}

func (r *repository) Delete(id int) error {
	_, err := sqlbuilder.Delete(r.db, "daily_stories", "id = ?", id)
	return err
}
