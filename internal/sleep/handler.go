package sleep

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service SleepService
}

func NewHandler(service SleepService) *Handler {
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
	activities, err := h.service.GetSleepsByUserId(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
}
func (h *Handler) GetSleep(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	Sleep, err := h.service.GetByID(id)
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

	SleepModel := Sleep{
		User_id:  input.User_id,
		SleepStart: input.SleepStart,
		SleepEnd: input.SleepEnd,
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

	existingSleep, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Sleep not found"})
		return
	}

	existingSleep.User_id = input.User_id
	existingSleep.SleepStart = input.SleepStart
	existingSleep.SleepEnd = input.SleepEnd

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