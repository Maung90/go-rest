package activity

import (
    "database/sql"
    "time"
)

// Tidak perlu lagi interface lokal jika kita menggunakan generic interface.
// Namun, jika Anda tetap ingin mendefinisikannya, pastikan NewRepository mengembalikan tipe konkrit.

type repository struct {
    db *sql.DB
}

// Perbaikan 1: NewRepository harus mengembalikan tipe konkrit (*repository)
// agar bisa digunakan oleh generic service.
func NewRepository(db *sql.DB) *repository {
    return &repository{db: db}
}

func (r *repository) FindAll() ([]Activity, error) {
    var activities []Activity // Penamaan jamak yang lebih umum

    rows, err := r.db.Query("SELECT id, user_id, activity, created_at, updated_at FROM user_activity")
    if err != nil {
        return activities, err
    }
    defer rows.Close()

    for rows.Next() {
        var activity Activity
        err := rows.Scan(&activity.ID, &activity.User_id, &activity.Activity, &activity.CreatedAt, &activity.UpdatedAt)
        if err != nil {
            return activities, err
        }
        activities = append(activities, activity)
    }
    return activities, nil
}

func (r *repository) FindByID(id int) (Activity, error) {
    var activity Activity
    err := r.db.QueryRow("SELECT id, user_id, activity, created_at, updated_at FROM user_activity WHERE id = ?", id).
        Scan(&activity.ID, &activity.User_id, &activity.Activity, &activity.CreatedAt, &activity.UpdatedAt)
    if err != nil {
        return activity, err
    }
    return activity, nil
}

func (r *repository) Save(activity Activity) (Activity, error) {
    // Perbaikan 2: Query INSERT memiliki 4 kolom tapi 5 placeholder (?).
    // Placeholder terakhir dihapus.
    query := "INSERT INTO user_activity (user_id, activity, created_at, updated_at) VALUES (?, ?, ?, ?)"
    result, err := r.db.Exec(query, activity.User_id, activity.Activity, time.Now(), time.Now())
    if err != nil {
        return activity, err
    }

    lastID, err := result.LastInsertId()
    if err != nil {
        return activity, err
    }

    // Perbaikan 3: Panggil FindByID untuk mendapatkan data lengkap
    // termasuk timestamp yang benar dari database.
    return r.FindByID(int(lastID))
}

func (r *repository) Update(activity Activity) (Activity, error) {
    query := "UPDATE user_activity SET user_id = ?, activity = ?, updated_at = ? WHERE id = ?"
    _, err := r.db.Exec(query, activity.User_id, activity.Activity, time.Now(), activity.ID)
    if err != nil {
        return activity, err
    }

    // Perbaikan 4: Sama seperti Save, panggil FindByID untuk mendapatkan data
    // terbaru setelah update.
    return r.FindByID(activity.ID)
}

func (r *repository) Delete(id int) error {
    _, err := r.db.Exec("DELETE FROM user_activity WHERE id = ?", id)
    return err
}