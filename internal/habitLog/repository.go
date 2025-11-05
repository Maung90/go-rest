package habitLog

import (
	"database/sql"
	"time"

	customsql "go-rest/pkg/sql"
)

type Repository interface {
	CreateLogs(newHabitLog HabitLog) (HabitLog, error)
	FindHabitLogs(date time.Time) ([]HabitLog, error)
	FindByID(id int) (HabitLog, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindByID(id int) (HabitLog, error) {
	builder := customsql.NewSQLBuilder(r.db, "habit_logs").
		Select("habit_id", "user_id", "log_date", "status").
		Where("id = ?", id)

	rows, err := builder.Get()
	if err != nil {
		return HabitLog{}, err
	}
	defer rows.Close()

	var h HabitLog
	if rows.Next() {
		err := rows.Scan(&h.ID, &h.Created_at, &h.Updated_at, &h.Habit_id, &h.User_id, &h.Log_date, &h.Status)
		if err != nil {
			return HabitLog{}, err
		}
	}
	return h, nil
}

func (r *repository) CreateLogs(newHabitLog HabitLog) (HabitLog, error) {
	data := map[string]interface{}{
		"habit_id":   newHabitLog.Habit_id,
		"user_id":    newHabitLog.User_id,
		"log_date":   newHabitLog.Log_date.Format("2006-01-02"),
		"status":     "done",
		"created_at": time.Now(),
		"updated_at": time.Now(),
	}

	result, err := customsql.Insert(r.db, "habit_logs", data)
	if err != nil {
		return HabitLog{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return HabitLog{}, err
	}

	return r.FindByID(int(lastID))
}

func (r *repository) FindHabitLogs(date time.Time) ([]HabitLog, error) {
	formattedDate := date.Format("2006-01-02")

	builder := customsql.NewSQLBuilder(r.db, "habit_logs").
		Select("habit_id", "user_id", "log_date", "status").
		Where("log_date = ?", formattedDate).
		OrderBy("created_at DESC")

	rows, err := builder.Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []HabitLog
	for rows.Next() {
		var h HabitLog
		err := rows.Scan(&h.ID, &h.Created_at, &h.Updated_at, &h.Habit_id, &h.User_id, &h.Log_date, &h.Status)
		if err != nil {
			return nil, err
		}
		logs = append(logs, h)
	}
	return logs, nil
}
