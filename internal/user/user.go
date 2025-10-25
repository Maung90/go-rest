package user

import "time"

// User struct merepresentasikan tabel 'users' di database
type User struct {
    ID        int       `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    Password  string    `json:"-"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
}

// Struct untuk menampung data input dari user
type CreateUserInput struct {
    Username string `json:"username" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=6"`
}

type UpdateUserInput struct {
    Username string `json:"username"`
    Email string `json:"email" binding:"email"`
}