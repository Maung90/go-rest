package auth

import (
	"database/sql"
	"errors"
	"go-rest/internal/user"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type Repository interface {
	DeleteRefreshToken(tokenUUID string) error
	FindByEmail(email string) (user.User, error)
	Register(newUser user.User) (user.User, error)
	FetchRefreshToken(tokenID string) (int, error)
	Login(email string, password string) (user.User, error)
	SaveRefreshToken(userID int, tokenID string, expiresAt time.Time) error
	FindByNameAndEmail(name, email string) (user.User, error)
	UpdatePassword(userID int, newPasswordHash string) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

type RefreshToken struct {
    ID         int       `json:"required"`
    UserID     int
    Token      string    `json:"unique"`
    ExpiredAt  time.Time
}

func (r *repository) FetchRefreshToken(tokenID string) (int, error) {
	var userID int
	query := "SELECT user_id FROM refresh_tokens WHERE id = ?"
	err := r.db.QueryRow(query, tokenID).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}

func (r *repository) SaveRefreshToken(userID int, tokenID string, expiresAt time.Time) error {
	query := "INSERT INTO refresh_tokens (id, user_id, expires_at) VALUES (?, ?, ?)"
	_, err := r.db.Exec(query, tokenID, userID, expiresAt)
	return err
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

func (r *repository) DeleteRefreshToken(tokenUUID string) error {
	query := "DELETE FROM refresh_tokens WHERE id = ?"
	_, err := r.db.Exec(query, tokenUUID)
	return err
}

func (r *repository) FindByNameAndEmail(name, email string) (user.User, error) {
	var foundUser user.User
	query := "SELECT id, name, email, password FROM users WHERE name = ? AND email = ?"
	err := r.db.QueryRow(query, name, email).Scan(&foundUser.ID, &foundUser.Name, &foundUser.Email, &foundUser.Password)
	return foundUser, err
}

func (r *repository) UpdatePassword(userID int, newPasswordHash string) error {
	query := "UPDATE users SET password = ? WHERE id = ?"
	_, err := r.db.Exec(query, newPasswordHash, userID)
	return err
}