package sleep

import (
	"database/sql"
	"time"
)


type Repository interface {
	FindAll(userID int) ([]Sleep, error)
	FindByID(userID, id int) (Sleep, error)
	Save(sleep Sleep) (Sleep, error)
	Update(sleep Sleep) (Sleep, error)
	Delete(id int) error
	GetWeeklyStats(userID int) ([]SleepStat, error) 
	GetMonthlyStats(userID int) ([]SleepStat, error) 
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}


func (r *repository) FindAll(userID int) ([]Sleep, error) {
	var sleeps []Sleep

	rows, err := r.db.Query("SELECT id, user_id, sleep_start, sleep_end, duration_hours, created_at, updated_at FROM sleep_logs WHERE user_id = ?", userID)
	if err != nil {
		return sleeps, err
	}
	defer rows.Close()

	for rows.Next() {
		var sleep Sleep
		err := rows.Scan(&sleep.ID, &sleep.User_id, &sleep.SleepStart, &sleep.SleepEnd, &sleep.Duration, &sleep.CreatedAt, &sleep.UpdatedAt)
		if err != nil {
			return sleeps, err
		}
		sleeps = append(sleeps, sleep)
	}
	return sleeps, nil
}

func (r *repository) FindByID(userID, id int) (Sleep, error) {
	var sleep Sleep
	err := r.db.QueryRow("SELECT id, user_id, sleep_start, sleep_end,duration_hours, created_at, updated_at FROM sleep_logs WHERE id = ? AND user_id = ?", id, userID).
	Scan(&sleep.ID, &sleep.User_id, &sleep.SleepStart, &sleep.SleepEnd, &sleep.Duration, &sleep.CreatedAt, &sleep.UpdatedAt)
	if err != nil {
		return sleep, err
	}
	return sleep, nil
}

func (r *repository) Save(sleep Sleep) (Sleep, error) {

	query := "INSERT INTO sleep_logs (user_id, sleep_start, sleep_end, duration_hours, created_at,  updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	result, err := r.db.Exec(query, sleep.User_id, sleep.SleepStart, sleep.SleepEnd, sleep.Duration, time.Now(), time.Now())
	if err != nil {
		return sleep, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return sleep, err
	}

	return r.FindByID(sleep.User_id, int(lastID))
}

func (r *repository) Update(sleep Sleep) (Sleep, error) {

	query := "UPDATE sleep_logs SET user_id = ?, sleep_start = ?, sleep_end = ?, duration_hours = ?, updated_at = ? WHERE id = ?"
	_, err := r.db.Exec(query, sleep.User_id, sleep.SleepStart, sleep.SleepEnd, sleep.Duration, time.Now(), sleep.ID)
	
	if err != nil {
		return sleep, err
	}
	return r.FindByID(sleep.User_id, sleep.ID)
}

func (r *repository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM sleep_logs WHERE id = ?", id)
	return err
}

func (r *repository) GetWeeklyStats(userID int) ([]SleepStat, error) {
	query := `
	SELECT
	CONCAT(YEAR(sleep_start), '-W', LPAD(WEEK(sleep_start, 1), 2, '0')) AS period,
	SUM(duration_hours) AS total_hours,
	AVG(duration_hours) AS avg_hours
	FROM sleep_logs
	WHERE user_id = ?
	GROUP BY YEAR(sleep_start), WEEK(sleep_start, 1)
	ORDER BY period DESC;
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []SleepStat
	for rows.Next() {
		var s SleepStat
		if err := rows.Scan(&s.Period, &s.TotalHours, &s.AvgHours); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}
	return stats, nil
}

func (r *repository) GetMonthlyStats(userID int) ([]SleepStat, error) {
	query := `
	SELECT
	DATE_FORMAT(sleep_start, '%Y-%m') AS period,
	SUM(duration_hours) AS total_hours,
	AVG(duration_hours) AS avg_hours
	FROM sleep_logs
	WHERE user_id = ?
	GROUP BY YEAR(sleep_start), MONTH(sleep_start)
	ORDER BY period DESC;
	`

	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []SleepStat
	for rows.Next() {
		var s SleepStat
		if err := rows.Scan(&s.Period, &s.TotalHours, &s.AvgHours); err != nil {
			return nil, err
		}
		stats = append(stats, s)
	}
	return stats, nil
}
