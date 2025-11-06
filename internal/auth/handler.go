 package auth

import (
	"go-rest/internal/service"
	"go-rest/internal/user"
	"net/http"
	"go-rest/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	authService AuthService
	userService *service.Service[user.User]
}

func NewHandler(authService AuthService, userService *service.Service[user.User]) *Handler {
	return &Handler{
		authService: authService,
		userService: userService,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Data yang diinputkan tidak valid", err)
		return
	}

	user, err := h.authService.Register(input)
	if err != nil {
			response.HandleError(c, err, "Registrasi gagal")
		return
	}
	response.Created(c, "Registrasi Berhasil", user)
}

func (h *Handler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.authService.FindByEmail(email)
	if err != nil {
			response.NotFound(c, "Data tidak ditemukan")
		return
	}
	response.OK(c, "Data berhasil diambil", user)
}

func (h *Handler) Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Data yang diinputkan tidak valid", err)
		return
	}

	tokens, err := h.authService.Login(input)
	if err != nil {
		response.Unauthorized(c, "Login gagal email atau password anda salah!")
		return
	}

	response.OK(c, "Login berhasil!", tokens)
}

func (h *Handler) Refresh(c *gin.Context) {
	type RefreshInput struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	var input RefreshInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Data yang diinputkan tidak valid", err)
		return
	}

	newAccessToken, err := h.authService.Refresh(input.RefreshToken)
	if err != nil {
		response.Unauthorized(c, "Token tidak valid")
		return
	}

	response.OK(c, "Token berhasil direfresh", gin.H{"access_token": newAccessToken})
}

func (h *Handler) Logout(c *gin.Context) {

	var input LogoutInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Data yang dikirim tidak valid", err.Error())
		return
	}

	err := h.authService.Logout(input.RefreshToken)
	if err != nil {
		response.HandleError(c, err, "Logout gagal, silahkan coba lagi")
		return
	}
	response.OK(c, "Logout berhasil", "")
}

func (h *Handler) GetProfile(c *gin.Context) {
	userIDContext, exists := c.Get("userID") // ambil user id yg di kirim dari middleware
	if !exists {
		response.Unauthorized(c, "User tidak ditemukan, atau autentikasi gagal")
		return
	}

	userID, ok := userIDContext.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tipe user id tidak sesuai"})
		return
	}

	userProfile, err := h.userService.GetByID(userID)
	if err != nil {
		response.HandleError(c, err, "Profile user tidak ditemukan")
		return
	}
	response.OK(c, "Profile berhasil didapatkan", userProfile)
}

func (h *Handler) UpdateProfile(c *gin.Context) {
	userIDContext, exists := c.Get("userID") // ambil user id yg di kirim dari middleware
	if !exists {
		response.Unauthorized(c, "User tidak ditemukan atau  atauautentikasi gagal")
		return
	}

	userID, ok := userIDContext.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Tipe user id tidak sesuai"})
		return
	}

	var input UpdateProfileInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Data yang dikirim tidak valid", err.Error())
		return
	}

	existingUser, err := h.userService.GetByID(userID)
	if err != nil {
			response.NotFound(c, "user id tidak ditemukan")
		return
	}

	existingUser.Name = input.Name
	existingUser.Email = input.Email

	updatedUser, err := h.userService.Update(existingUser)
	if err != nil {
		response.HandleError(c,err, "Data gagal di update")
		return
	}
	response.OK(c, "Profile berhasil diupdate", updatedUser) 
}

func (h *Handler) ForgotPassword(c *gin.Context) {
	var input ForgotPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resetToken, err := h.authService.ForgotPassword(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
// kirim email seharusnya 
	c.JSON(http.StatusOK, gin.H{"reset_token": resetToken})
}

func (h *Handler) ResetPassword(c *gin.Context) {
	var input ResetPasswordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Data yang diinputkan tidak valid", err)
		return
	}

	err := h.authService.ResetPassword(input)
	if err != nil {
		response.HandleError(c,err, "Password gagal di update")
		return
	}
	response.OK(c, "Password berhasil diupdate", "")
}