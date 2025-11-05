package activity

import (
	"database/sql"
	// "time"
)


type Repository interface {
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db: db}
}

