package auth

import (
	"database/sql"
	"errors"
	"go-rest/internal/user"
	"golang.org/x/crypto/bcrypt"
	"time"
	sqlbuilder "go-rest/pkg/sql"
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
	query, err := sqlbuilder.RawQuery(r.db, `SELECT user_id FROM refresh_tokens WHERE id = ?`, tokenID)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	if query.Next(){
		err := query.Scan(&userID)
		if err != nil {
			return 0, err
		}
	}
	return userID, nil
}

func (r *repository) SaveRefreshToken(userID int, tokenID string, expiresAt time.Time) error {
	data := map[string]interface{}{
		"id" 		 : tokenID, 
		"user_id" 	 : userID, 
		"expires_at" : expiresAt,
	}
	_, err := sqlbuilder.Insert(r.db, "refresh_tokens", data)
	return err
}

func (r *repository) Register(newUser user.User) (user.User, error) {
	data := map[string]interface{}{
		"name"							:    newUser.Name,
		"email"						: newUser.Email,
		"password"			: newUser.Password,
		"created_at" : time.Now(),
		"updated_at" : time.Now(),

	}

	result, err := sqlbuilder.Insert(r.db, "users", data)
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

	builder := sqlbuilder.NewSQLBuilder(r.db, "users").
	Select("name", "email", "password").
	Where("email = ?", email)
	rows, err := builder.Get()
	if err != nil {
		return user.User{}, err
	}
	defer rows.Close()

	if rows.Next(){
		err := rows.Scan(
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
	}
	return foundUser, nil
}

func (r *repository) FindByID(id int) (user.User, error) {
	var foundUser user.User
	builder := sqlbuilder.NewSQLBuilder(r.db, "users").
	Select("name", "email", "password").
	Where("id = ?", id)

	rows, err := builder.Get()
	if err != nil {
		return user.User{}, err
	}
	defer rows.Close()

	if rows.Next(){
		err := rows.Scan(
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
	_, err := sqlbuilder.Delete(r.db, "refresh_tokens", "id = ?", tokenUUID)
	return err
}

func (r *repository) FindByNameAndEmail(name, email string) (user.User, error) {
	var foundUser user.User
	rows, err := sqlbuilder.RawQuery(r.db, `
		SELECT id, name, email, password FROM users WHERE name = ? AND email = ?
		`, name, email)
		if err != nil {
			return user.User{}, err
		}
		defer rows.Close()

		if rows.Next() {
			err := rows.Scan(&foundUser.ID, &foundUser.Name, &foundUser.Email, &foundUser.Password)
			if err != nil {
				return user.User{}, err
			}
		}
		return foundUser, err
	}

	func (r *repository) UpdatePassword(userID int, newPasswordHash string) error {
		data := map[string]interface{}{
			"password": newPasswordHash,
		}
		_, err := sqlbuilder.Update(
			r.db,
			"daily_stories",
			data,
			"user_id = ?",
			userID,
		)
		return err
	}