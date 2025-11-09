package activity

import (
	"database/sql"
	"time"

	sqlbuilder "go-rest/pkg/sql"
)

type Repository interface {
	FindAll(userId int) ([]Activity, error)
	FindByDate(date string, userId int) ([]Activity, error)
	FindById(id, userId int) (Activity, error)
	Save(activity Activity) (Activity, error)
	Update(activity Activity) (Activity, error)
	Delete(id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(userId int) ([]Activity, error) {
	builder := sqlbuilder.NewSQLBuilder(r.db, "activities").
	Select("user_id", "activity_date", "title", "duration_minutes", "notes").
	Where("user_id = ?", userId).
	OrderBy("activity_date DESC")

	rows, err := builder.Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activities []Activity
	for rows.Next() {
		var s Activity
		err := rows.Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt, &s.User_id, &s.ActivityDate, &s.Title, &s.DurationMinutes, &s.Notes)
		if err != nil {
			return nil, err
		}
		activities = append(activities, s)
	}
	return activities, nil
}

func (r *repository) FindByDate(date string, userId int) ([]Activity, error) {
	var activities []Activity

	builder := sqlbuilder.NewSQLBuilder(r.db, "activities").
	Select("user_id", "activity_date", "title", "duration_minutes", "notes").
	Where("user_id = ? AND activity_date = ?", userId, date).
	OrderBy("activity_date DESC")

	rows, err := builder.Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next(){
		var activity Activity
		err := rows.Scan(
			&activity.ID,
			&activity.CreatedAt,
			&activity.UpdatedAt,
			&activity.User_id,
			&activity.ActivityDate,
			&activity.Title,
			&activity.DurationMinutes,
			&activity.Notes,
		)
		if err != nil {
			return nil, err
		}
		activities = append(activities,activity)
	}
	return activities, nil
}

func (r *repository) FindById(id int, userId int) (Activity, error) {
	var activity Activity
	builder := sqlbuilder.NewSQLBuilder(r.db, "activities").
	Select("user_id", "activity_date", "title", "duration_minutes", "notes").
	Where("id = ? AND user_id = ?", id, userId).OrderBy("story_date DESC")

	rows,err := builder.Get()
	if err != nil {
		return activity, err
	}
	defer rows.Close()

	for rows.Next(){
		err := rows.Scan(
			&activity.ID,
			&activity.CreatedAt,
			&activity.UpdatedAt,
			&activity.User_id,
			&activity.ActivityDate,
			&activity.Title,
			&activity.DurationMinutes,
			&activity.Notes,
		)
		if err != nil {
			return activity, err
		}
	}
	return activity, err
}

func (r *repository) Save(activity Activity) (Activity, error) {
	data := map[string]interface{}{
		"user_id"									:	activity.User_id,
		"activity_date"			: activity.ActivityDate,
		"title"											: activity.Title,
		"duration_minutes": activity.DurationMinutes,
		"notes"											: activity.Notes,
		"created_at"						: time.Now(),
		"updated_at"						: time.Now(),
	}
	result, err := sqlbuilder.Insert(r.db, "activities", data)
	if err != nil {
		return activity, err
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		return activity, err
	}

	return r.FindById(int(lastId), activity.User_id)
}

func (r *repository) Update(activity Activity) (Activity, error) {
	data := map[string]interface{}{
		"activity_date"			: activity.ActivityDate,
		"title"											: activity.Title,
		"duration_minutes": activity.DurationMinutes,
		"notes"											: activity.Notes,
		"updated_at"						: time.Now(),
	}
	_,err := sqlbuilder.Update(
		r.db,
		"activities",
		data,
		"id = ? AND user_id = ?",
		activity.ID, 
		activity.User_id,
	)
	if err != nil {
		return Activity{}, err
	}

	return r.FindById(activity.ID, activity.User_id)
}

func (r *repository) Delete(id int) error {
	_, err := sqlbuilder.Delete(r.db, "activities", "id = ?", id)
	return err
}


