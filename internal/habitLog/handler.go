package habitLog

import (
	"fmt"
	"go-rest/internal/habit"
	"go-rest/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	habitLogService HabitLogService                       
	habitService    *service.Service[habit.Habit]
}

func NewHandler(habitLogService HabitLogService, habitService *service.Service[habit.Habit]) *Handler {
	return &Handler{
		habitLogService: habitLogService,
		habitService:    habitService,
	}
}

func (h *Handler) CreateLogs(c *gin.Context) {
	habit_id, _ := strconv.Atoi(c.Param("id"))

	var input CreateHabitLogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.habitService.GetByID(habit_id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit not found"})
		return
	}

	input.Habit_id = habit_id
	input.User_id = input.User_id
	input.LogDate = input.LogDate

	habitlog, err := h.habitLogService.CreateLogs(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, habitlog)
}

func (h *Handler) GetLogsByDate(c *gin.Context) {
	var input GetHabitLogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("DEBUG log_date:", input)

	habitLogs, err := h.habitLogService.FindHabitLogs(input)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit Logs not found for that date"})
		return
	}
	c.JSON(http.StatusOK, habitLogs)
}