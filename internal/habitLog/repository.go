package habitLog

import (
	"database/sql"
	"time"
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
	var foundLog HabitLog
	query := "SELECT id, habit_id, user_id, log_date, status, created_at, updated_at FROM habit_logs WHERE id = ?"
	
	err := r.db.QueryRow(query, id).Scan(
		&foundLog.ID,
		&foundLog.Habit_id,
		&foundLog.User_id,
		&foundLog.Log_date,
		&foundLog.Status,
		&foundLog.Created_at,
		&foundLog.Updated_at,
	)
	if err != nil {
		return HabitLog{}, err
	}
	return foundLog, nil
}


func (r *repository) CreateLogs(newHabitLog HabitLog) (HabitLog, error) {
	query := "INSERT INTO habit_logs (habit_id, user_id, log_date, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	logDateString := newHabitLog.Log_date.Format("2006-01-02")
	result, err := r.db.Exec(query, newHabitLog.Habit_id, newHabitLog.User_id, logDateString, "done", time.Now(), time.Now())
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
	var foundLogs []HabitLog

	parseDate := date.Format("2006-01-02")
	query := "SELECT id, habit_id, user_id, log_date, status, created_at, updated_at FROM habit_logs WHERE log_date = ?"
	rows, err := r.db.Query(query, parseDate)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var log HabitLog
		err := rows.Scan(
			&log.ID,
			&log.Habit_id,
			&log.User_id,
			&log.Log_date,
			&log.Status,
			&log.Created_at,
			&log.Updated_at,
		)
		if err != nil {
			return nil, err
		}
		foundLogs = append(foundLogs, log)
	}

	return foundLogs, nil
}