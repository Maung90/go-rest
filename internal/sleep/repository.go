package sleep

import (
	"database/sql"
	"time"

	customsql "go-rest/pkg/sql"
)

type Repository interface {
	FindAll(userID int) ([]Sleep, error)
	FindByID(userID, id int) (Sleep, error)
	Save(s Sleep) (Sleep, error)
	Update(s Sleep) (Sleep, error)
	Delete(id int) error
	GetWeeklyStats(userID int) ([]SleepStat, error)
	GetMonthlyStats(userID int) ([]SleepStat, error)
	FindByDate(userID int, date string) ([]Sleep, error) 
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) FindAll(userID int) ([]Sleep, error) {
	builder := customsql.NewSQLBuilder(r.db, "sleep_logs").
	Select("user_id", "sleep_start", "sleep_end", "duration_hours").
	Where("user_id = ?", userID).
	OrderBy("created_at DESC")

	rows, err := builder.Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sleeps []Sleep
	for rows.Next() {
		var s Sleep
		err := rows.Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt, &s.User_id, &s.SleepStart, &s.SleepEnd, &s.Duration)
		if err != nil {
			return nil, err
		}
		sleeps = append(sleeps, s)
	}
	return sleeps, nil
}

func (r *repository) FindByID(userID, id int) (Sleep, error) {
	builder := customsql.NewSQLBuilder(r.db, "sleep_logs").
	Select("user_id", "sleep_start", "sleep_end", "duration_hours").
	Where("id = ?", id).
	Where("user_id = ?", userID)

	rows, err := builder.Get()
	if err != nil {
		return Sleep{}, err
	}
	defer rows.Close()

	var s Sleep
	if rows.Next() {
		err := rows.Scan(&s.ID, &s.CreatedAt, &s.UpdatedAt, &s.User_id, &s.SleepStart, &s.SleepEnd, &s.Duration)
		if err != nil {
			return Sleep{}, err
		}
	}
	return s, nil
}

func (r *repository) Save(s Sleep) (Sleep, error) {
	data := map[string]interface{}{
		"user_id":         s.User_id,
		"sleep_start":     s.SleepStart,
		"sleep_end":       s.SleepEnd,
		"duration_hours":  s.Duration,
		"created_at":      time.Now(),
		"updated_at":      time.Now(),
	}

	result, err := customsql.Insert(r.db, "sleep_logs", data)
	if err != nil {
		return Sleep{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return Sleep{}, err
	}

	return r.FindByID(s.User_id, int(lastID))
}

func (r *repository) Update(s Sleep) (Sleep, error) {
	data := map[string]interface{}{
		"sleep_start":     s.SleepStart,
		"sleep_end":       s.SleepEnd,
		"duration_hours":  s.Duration,
		"updated_at":      time.Now(),
	}

	_, err := customsql.Update(r.db, "sleep_logs", data, "id = ? AND user_id = ?", s.ID, s.User_id)
	if err != nil {
		return Sleep{}, err
	}

	return r.FindByID(s.User_id, s.ID)
}

func (r *repository) Delete(id int) error {
	_, err := customsql.Delete(r.db, "sleep_logs", "id = ?", id)
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

	rows, err := customsql.RawQuery(r.db, query, userID)
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

	rows, err := customsql.RawQuery(r.db, query, userID)
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

func (r *repository) FindByDate(userID int, date string) ([]Sleep, error) {

	builder := customsql.NewSQLBuilder(r.db, "sleep_logs").
	Select("sleep_start", "sleep_end", "duration_hours").
	Where("sleep_end = ? AND user_id = ? ", date, userID).
	OrderBy("created_at DESC")

	rows, err := builder.Get()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var logs []Sleep
	for rows.Next() {
		var h Sleep
		err := rows.Scan(&h.ID, &h.CreatedAt, &h.UpdatedAt, &h.SleepStart, &h.SleepEnd, &h.Duration)
		if err != nil {
			return nil, err
		}
		logs = append(logs, h)
	}
	return logs, nil
}