package habitLog

import (
	"fmt"
	"go-rest/internal/habit"
	"go-rest/internal/service"
	"net/http"
	"strconv"
	"go-rest/pkg/response"
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
		response.ValidationError(c, "Log Habit gagal disimpan, data yang diinputkan tidak valid", err)
		return
	}

	_, err := h.habitService.GetByID(habit_id)
	if err != nil {
		response.NotFound(c, "id habit tidak ditemukan")
		return
	}

	input.Habit_id = habit_id
	input.User_id = input.User_id
	input.LogDate = input.LogDate

	habitlog, err := h.habitLogService.CreateLogs(input)
	if err != nil {
		response.HandleError(c, err, "Log Habit gagal disimpan")
		return
	}
	response.Created(c, "Log habit berhasil disimpan", habitlog)
}

func (h *Handler) GetLogsByDate(c *gin.Context) {
	var input GetHabitLogInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c,"Data yang diinputkan tidak valid", err)
		return
	}
	fmt.Println("DEBUG log_date:", input)

	habitLogs, err := h.habitLogService.FindHabitLogs(input)
	if err != nil {
		response.NotFound(c, "habit log tidak ditemukan pada tanggal tersebut")
		c.JSON(http.StatusNotFound, gin.H{"error": "Habit Logs not found for that date"})
		return
	}
	response.OK(c, "habit log berhasil diambil", habitLogs)
}