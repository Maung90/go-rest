package habit

import (
	"net/http"
	"strconv"
	"go-rest/pkg/response"
	"go-rest/internal/service" 
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service[Habit]
}

func NewHandler(service *service.Service[Habit]) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetHabits(c *gin.Context) {
	habits, err := h.service.GetAll()
	if err != nil {
		response.HandleError(c, err, "Habit gagal diambil")
		return
	}
	response.OK(c, "Habit berhasil Diambil", habits)
}

func (h *Handler) GetHabit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Habit, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Habit tidak ditemukan")
		return
	}
	response.OK(c, "Habit berhasil Diambil", Habit)
}

func (h *Handler) CreateHabit(c *gin.Context) {
	var input CreateHabitInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Habit gagal disimpan, data yang diinputkan tidak valid", err)
		return
	}

	HabitModel := Habit{
		User_id:  input.User_id,
		Title: input.Title,
		Description: input.Description,
	}

	newHabit, err := h.service.Create(HabitModel)
	if err != nil {
		response.HandleError(c, err, "Habit gagal disimpan")
		return
	} 
	response.Created(c, "Habit berhasil disimpan", newHabit)
}

func (h *Handler) UpdateHabit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input UpdateHabitInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Habit gagal disimpan, data yang diinputkan tidak valid", err)
		return
	} 

	existingHabit, err := h.service.GetByID(id)
	if err != nil {
		response.NotFound(c, "Id habit tidak ditemukan")
		return
	}

	existingHabit.User_id = input.User_id
	existingHabit.Title = input.Title
	existingHabit.Description = input.Description

	updatedHabit, err := h.service.Update(existingHabit)
	if err != nil {
		response.HandleError(c, err, "Habit gagal disimpan")
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response.Created(c, "Habit berhasil disimpan", updatedHabit)
}

func (h *Handler) DeleteHabit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		response.HandleError(c, err, "Habit gagal dihapus")
		return
	}
	response.OK(c, "Habit berhasil dihapus", "")
}