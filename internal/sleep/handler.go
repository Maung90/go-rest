package sleep

import (
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetSleeps(c *gin.Context) {

	userIDContext, exists := c.Get("userID") // ambil user id yg di kirim dari middleware
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context. Authorization may have failed."})
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID in context is not of expected type"})
		return
	}
	activities, err := h.service.GetByUserID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
}
func (h *Handler) GetSleep(c *gin.Context) {
	userIDContext, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context. Authorization may have failed."})
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID in context is not of expected type"})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	Sleep, err := h.service.GetByID(userID, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sleep not found"})
		return
	}
	c.JSON(http.StatusOK, Sleep)
}

func (h *Handler) CreateSleep(c *gin.Context) {
	var input SleepInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	layout := "2006-01-02 15:04:05"

	startTime, err := time.Parse(layout, input.SleepStart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sleep_start format"})
		return
	}

	endTime, err := time.Parse(layout, input.SleepEnd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sleep_end format"})
		return
	}

	duration := endTime.Sub(startTime).Hours()
	SleepModel := Sleep{
		User_id:  input.User_id,
		SleepStart: startTime,
		SleepEnd: endTime,
		Duration: duration,
	}

	newSleep, err := h.service.Create(SleepModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newSleep)
}

func (h *Handler) UpdateSleep(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input SleepInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} 

	existingSleep, err := h.service.GetByID(input.User_id, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sleep not found"})
		return
	}

	layout := "2006-01-02 15:04:05"

	startTime, err := time.Parse(layout, input.SleepStart)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sleep_start format"})
		return
	}

	endTime, err := time.Parse(layout, input.SleepEnd)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sleep_end format"})
		return
	}
	duration := endTime.Sub(startTime).Hours()

	existingSleep.User_id = input.User_id
	existingSleep.SleepStart = startTime
	existingSleep.SleepEnd = endTime
	existingSleep.Duration = duration

	updatedSleep, err := h.service.Update(existingSleep)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedSleep)
}

func (h *Handler) DeleteSleep(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sleep not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Sleep deleted successfully"})
}

func (h *Handler) GetWeeklyStats(c *gin.Context) {
	userIDContext, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context. Authorization may have failed."})
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID in context is not of expected type"})
		return
	}
	stats, err := h.service.GetWeeklyStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}

func (h *Handler) GetMonthlyStats(c *gin.Context) {
	userIDContext, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found in context. Authorization may have failed."})
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID in context is not of expected type"})
		return
	}
	stats, err := h.service.GetMonthlyStats(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, stats)
}
