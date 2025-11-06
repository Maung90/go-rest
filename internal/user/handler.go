package user

import (
	"strconv"
	"go-rest/pkg/response"
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
		response.HandleError(c, err, "Gagal mengambil data users")
		return
	}
	response.OK(c, "Data user berhasil diambil", users)
}

func (h *Handler) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Gagal mengambil data users")
		return
	}
	response.OK(c, "Data user berhasil diambil", user)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Data user tidak sesuai validasi", err)
		return
	}

	userModel := User{
		Name:  input.Name,
		Email: input.Email,
		Password: input.Password,
	}
	newUser, err := h.service.Create(userModel)
	if err != nil {
		response.HandleError(c, err, "Gagal membuat user")
		return
	}
	response.Created(c, "Data user berhasil disimpan", newUser)
}

func (h *Handler) UpdateUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input UpdateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Input tidak sesuai validasi", err)
		return
	}

	existingUser, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "id user tidak dapat ditemukan")
		return
	}

	existingUser.Name = input.Name
	existingUser.Email = input.Email

	updatedUser, err := h.service.Update(existingUser)
	if err != nil {
		response.HandleError(c, err, "Data user gagal diupdate")
		return
	}
	response.OK(c, "Data user berhasil diupdate", updatedUser)
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		response.NotFound(c, "Gagal dihapus, user tidak ditemukan")
		return
	}
	response.OK(c, "Data user berhasil di hapus", "")
}