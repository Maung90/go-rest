package user

import "time"

// User struct merepresentasikan tabel 'users' di database
type User struct {
    ID        int       `json:"id"`
    Name      string    `json:"name"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// Struct untuk menampung data input dari user
type CreateUserInput struct {
    Name     string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserInput struct {
    Name  string `json:"name"`
    Email string `json:"email" binding:"email"`
}