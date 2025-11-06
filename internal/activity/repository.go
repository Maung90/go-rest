package activity

import (
	"database/sql"
	// "time"

	sqlbuilder "go-rest/pkg/sql"
)

type Repository interface {
	FindAll(userid int) ([]Activity, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}


func (r *repository) FindAll(userid int) ([]Activity, error) {
	builder := sqlbuilder.NewSQLBuilder(r.db, "activities").
		Select("user_id", "activity_date", "title", "duration_minutes", "notes").
		Where("user_id = ?", userid).
		OrderBy("activity_date DESC")

	rows, err := builder.Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var activites []Activity
	for rows.Next() {
		var s Activity
		err := rows.Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt, &s.User_id, &s.ActivityDate, &s.Title, &s.DurationMinutes, &s.Notes)
		if err != nil {
			return nil, err
		}
		activites = append(activites, s)
	}
	return activites, nil
}
