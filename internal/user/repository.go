package user

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

func (r *repository) FindAll() ([]User, error) {
    var users []User
    // Menggunakan placeholder ? untuk MySQL
    rows, err := r.db.Query("SELECT id, username, email, created_at, updated_at FROM users")
    if err != nil {
        return users, err
    }
    defer rows.Close()

    for rows.Next() {
        var user User
        err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
        if err != nil {
            return users, err
        }
        users = append(users, user)
    }
    return users, nil
}

func (r *repository) FindByID(id int) (User, error) {
    var user User
    // Menggunakan placeholder ?
    err := r.db.QueryRow("SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?", id).
        Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
    if err != nil {
        return user, err
    }
    return user, nil
}

func (r *repository) Save(user User) (User, error) {
    // Query INSERT untuk MySQL
    query := "INSERT INTO users (username, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
    result, err := r.db.Exec(query, user.Username, user.Email, user.Password, time.Now(), time.Now())
    if err != nil {
        return user, err
    }

    // Ambil ID terakhir yang di-generate oleh MySQL
    lastID, err := result.LastInsertId()
    if err != nil {
        return user, err
    }
    user.ID = int(lastID)

    return user, nil
}

func (r *repository) Update(user User) (User, error) {
    // Query UPDATE untuk MySQL
    query := "UPDATE users SET username = ?, email = ?, updated_at = ? WHERE id = ?"
    _, err := r.db.Exec(query, user.Username, user.Email, time.Now(), user.ID)
    if err != nil {
        return user, err
    }
    return user, nil
}

func (r *repository) Delete(id int) error {
    // Query DELETE untuk MySQL
    _, err := r.db.Exec("DELETE FROM users WHERE id = ?", id)
    return err
}