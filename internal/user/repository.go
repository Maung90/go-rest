package user

import (
    "database/sql"
    "time"

    customsql "go-rest/pkg/sql"
)

type Repository interface {
    FindAll() ([]User, error)
    FindByID(id int) (User, error)
    Save(user User) (User, error)
    Update(user User) (User, error)
    Delete(id int) error
}

type repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
    return &repository{db: db}
}

func (r *repository) FindAll() ([]User, error) {
    builder := customsql.NewSQLBuilder(r.db, "users").
        Select("name", "email").
        OrderBy("created_at DESC")

    rows, err := builder.Get()
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var users []User
    for rows.Next() {
        var u User
        err := rows.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Name, &u.Email)
        if err != nil {
            return nil, err
        }
        users = append(users, u)
    }
    return users, nil
}

func (r *repository) FindByID(id int) (User, error) {
    builder := customsql.NewSQLBuilder(r.db, "users").
        Select("name", "email").
        Where("id = ?", id)

    rows, err := builder.Get()
    if err != nil {
        return User{}, err
    }
    defer rows.Close()

    var u User
    if rows.Next() {
        err := rows.Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt, &u.Name, &u.Email)
        if err != nil {
            return User{}, err
        }
    }
    return u, nil
}

func (r *repository) Save(u User) (User, error) {
    data := map[string]interface{}{
        "name":       u.Name,
        "email":      u.Email,
        "password":   u.Password,
        "created_at": time.Now(),
        "updated_at": time.Now(),
    }

    result, err := customsql.Insert(r.db, "users", data)
    if err != nil {
        return User{}, err
    }

    lastID, err := result.LastInsertId()
    if err != nil {
        return User{}, err
    }

    return r.FindByID(int(lastID))
}

func (r *repository) Update(u User) (User, error) {
    data := map[string]interface{}{
        "name":       u.Name,
        "email":      u.Email,
        "updated_at": time.Now(),
    }

    _, err := customsql.Update(r.db, "users", data, "id = ?", u.ID)
    if err != nil {
        return User{}, err
    }

    return r.FindByID(u.ID)
}

func (r *repository) Delete(id int) error {
    _, err := customsql.Delete(r.db, "users", "id = ?", id)
    return err
}
