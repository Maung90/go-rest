package auth

type RegisterInput struct {
    Name string `json:"name" binding:"required"`
    Email    string `json:"email" binding:"required,email"`
    Password string `json:"password" binding:"required,min=8"`
}

type ForgotPassword struct {
    Name string `json:"name"`
    Email string `json:"email" binding:"email"`
}

type LoginInput struct {
    Email string `json:"email"`
    Password string `json:"password" binding:"required,min=8"`
}

type LoginResponse struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
}

type LogoutInput struct {
    RefreshToken string `json:"refresh_token" binding:"required"`
}

type UpdateProfileInput struct {
    Name  string `json:"name"`
    Email string `json:"email" binding:"email"`
}

type ForgotPasswordInput struct {
    Name string `json:"name" binding:"required"`
    Email    string `json:"email"    binding:"required,email"`
}

type ResetPasswordInput struct {
    Token       string `json:"token"        binding:"required"`
    NewPassword string `json:"new_password" binding:"required,min=8"`
}