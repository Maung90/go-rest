package habit

import (
	"net/http"
	"strconv"

	"go-rest/internal/service" 
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service[Habit]
}

// Perbaikan 1: Tipe parameter 'service' harus sesuai dengan struct field
func NewHandler(service *service.Service[Habit]) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetHabits(c *gin.Context) {
	activities, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
}

func (h *Handler) GetHabit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Habit, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}
	c.JSON(http.StatusOK, Habit)
}

func (h *Handler) CreateHabit(c *gin.Context) {
	var input CreateHabitInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	HabitModel := Habit{
		User_id:  input.User_id,
		Title: input.Title,
		Description: input.Description,
	}

	newHabit, err := h.service.Create(HabitModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newHabit)
}

func (h *Handler) UpdateHabit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input UpdateHabitInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	existingHabit, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}

	existingHabit.User_id = input.User_id
	existingHabit.Title = input.Title
	existingHabit.Description = input.Description

	updatedHabit, err := h.service.Update(existingHabit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedHabit)
}

func (h *Handler) DeleteHabit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Habit deleted successfully"})
}