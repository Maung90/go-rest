package habit

import (
    "database/sql"
    "time"
    sqlbuilder "go-rest/pkg/sql"
)


type repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
    return &repository{db: db}
}

func (r *repository) FindAll() ([]Habit, error) {
    var habits []Habit
    builder := sqlbuilder.NewSQLBuilder(r.db, "habits").
    Select("user_id", "title", "description")
    
    rows, err := builder.Get()
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
    builder := sqlbuilder.NewSQLBuilder(r.db, "habits").
    Select("user_id", "title", "description").
    Where("id = ?", id)

    rows, err := builder.Get()
    if err != nil {
        return habit, err
    }
    defer rows.Close()

    if rows.Next(){ 
        err := rows.
        Scan(&habit.ID, &habit.User_id, &habit.Title, &habit.Description, &habit.CreatedAt, &habit.UpdatedAt)
        if err != nil {
            return habit, err
        }
    }

    return habit, nil
}

func (r *repository) Save(habit Habit) (Habit, error) {

    data := map[string]interface{}{
        "user_id"       : habit.User_id, 
        "title"         : habit.Title, 
        "description"   : habit.Description,
        "created_at"    : time.Now(),
        "updated_at"    : time.Now(),
    }

    result, err := sqlbuilder.Insert(r.db, "habits", data)
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

    data := map[string]interface{}{
        "user_id"       : habit.User_id, 
        "title"         : habit.Title, 
        "description"   : habit.Description,
        "updated_at"    : time.Now(),
    }

    _, err := sqlbuilder.Update(
        r.db,
        "habits",
        data,
        "id = ?",
        habit.ID,
    )
    if err != nil {
        return habit, err
    }
    return r.FindByID(habit.ID)
}

func (r *repository) Delete(id int) error {
    _, err := sqlbuilder.Delete(r.db, "habits", "id = ?", id)
    return err
}