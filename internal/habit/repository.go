package habit

import (
    "database/sql"
    "time"
)


type repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
    return &repository{db: db}
}

func (r *repository) FindAll() ([]Habit, error) {
    var habits []Habit

    rows, err := r.db.Query("SELECT id, user_id, title, description, created_at, updated_at FROM habits")
    if err != nil {
        return habits, err
    }
    defer rows.Close()

    for rows.Next() {
        var habit Habit
        err := rows.Scan(&habit.ID, &habit.User_id, &habit.Title, &habit.Description, &habit.CreatedAt, &habit.UpdatedAt)
        if err != nil {
            return habits, err
        }
        habits = append(habits, habit)
    }
    return habits, nil
}

func (r *repository) FindByID(id int) (Habit, error) {
    var habit Habit
    err := r.db.QueryRow("SELECT id, user_id, title, description, created_at, updated_at FROM habits WHERE id = ?", id).
        Scan(&habit.ID, &habit.User_id, &habit.Title, &habit.Description, &habit.CreatedAt, &habit.UpdatedAt)
    if err != nil {
        return habit, err
    }
    return habit, nil
}

func (r *repository) Save(habit Habit) (Habit, error) {
    
    query := "INSERT INTO habits (user_id, title, description, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
    result, err := r.db.Exec(query, habit.User_id, habit.Title, habit.Description, time.Now(), time.Now())
    if err != nil {
        return habit, err
    }

    lastID, err := result.LastInsertId()
    if err != nil {
        return habit, err
    }

    return r.FindByID(int(lastID))
}

func (r *repository) Update(habit Habit) (Habit, error) {
    query := "UPDATE habits SET user_id = ?, title = ?, description = ?, updated_at = ? WHERE id = ?"
    _, err := r.db.Exec(query, habit.User_id, habit.Title, habit.Description, time.Now(), habit.ID)
    if err != nil {
        return habit, err
    }
    return r.FindByID(habit.ID)
}

func (r *repository) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM habits WHERE id = ?", id)
    return err
}