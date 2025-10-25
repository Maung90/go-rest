package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go-rest/internal/service"
)

type Handler struct {
	service *service.Service[User]
}

func NewHandler(service *service.Service[User]) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (h *Handler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbaikan 2: Mapping dari DTO (input) ke Model (User)
	// Service kita mengharapkan model User, bukan CreateUserInput
	userModel := User{
		Username:  input.Username,
		Email: input.Email,
		Password: input.Password,
	}
	newUser, err := h.service.Create(userModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newUser)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbaikan 3: Proses Update yang Benar
	// 1. Ambil dulu data yang ada di database
	existingUser, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 2. Ubah field-nya dengan data dari input
	existingUser.Username = input.Username
	existingUser.Email = input.Email

	// 3. Kirim model yang sudah lengkap ke service
	updatedUser, err := h.service.Update(existingUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}