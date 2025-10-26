package auth

import (
	"golang.org/x/crypto/bcrypt"
	"go-rest/internal/user"
	"database/sql"
	"errors"
	"time"

)

type Repository interface {
	Register(newUser user.User) (user.User, error)
	FindByEmail(email string) (user.User, error)
	Login(email string, password string)(user.User, error) 
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Register(newUser user.User) (user.User, error) {
	query := "INSERT INTO users (name, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?)"
	result, err := r.db.Exec(query, newUser.Name, newUser.Email, newUser.Password, time.Now(), time.Now())
	if err != nil {
		return user.User{}, err
	}

	lastID, err := result.LastInsertId()
	if err != nil {
		return user.User{}, err
	}
	return r.FindByID(int(lastID))
}

func (r *repository) FindByEmail(email string) (user.User, error) {
	var foundUser user.User
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"
	err := r.db.QueryRow(query, email).Scan(
		&foundUser.ID,
		&foundUser.Name,
		&foundUser.Email,
		&foundUser.Password,
		&foundUser.CreatedAt,
		&foundUser.UpdatedAt,
	)
	if err != nil {
		return user.User{}, err
	}
	return foundUser, nil
}

func (r *repository) FindByID(id int) (user.User, error) {
	var foundUser user.User
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(
		&foundUser.ID,
		&foundUser.Name,
		&foundUser.Email,
		&foundUser.Password,
		&foundUser.CreatedAt,
		&foundUser.UpdatedAt,
	)
	if err != nil {
		return user.User{}, err
	}
	return foundUser, nil
}

func (r *repository) Login(email string, password string) (user.User, error) {
	foundUser, err := r.FindByEmail(email)
	if err != nil {
		return user.User{}, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(password))
	if err != nil {
		return user.User{}, errors.New("invalid email or password")
	}
	return foundUser, nil
}