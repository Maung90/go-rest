package activity

import (
	"net/http"
	"strconv"

	"go-rest/internal/service" // Pastikan import ini benar
	"github.com/gin-gonic/gin"
)

type Handler struct {
	// Tipe data service sudah benar
	service *service.Service[Activity]
}

// Perbaikan 1: Tipe parameter 'service' harus sesuai dengan struct field
func NewHandler(service *service.Service[Activity]) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetActivitys(c *gin.Context) {
	activities, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, activities)
}

func (h *Handler) GetActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	activity, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}
	c.JSON(http.StatusOK, activity)
}

func (h *Handler) CreateActivity(c *gin.Context) {
	var input CreateActivityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbaikan 2: Mapping dari DTO (input) ke Model (Activity)
	// Service kita mengharapkan model Activity, bukan CreateActivityInput
	activityModel := Activity{
		User_id:  input.User_id,
		Activity: input.Activity,
	}

	newActivity, err := h.service.Create(activityModel)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newActivity)
}

func (h *Handler) UpdateActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input UpdateActivityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Perbaikan 3: Proses Update yang Benar
	// 1. Ambil dulu data yang ada di database
	existingActivity, err := h.service.GetByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}

	// 2. Ubah field-nya dengan data dari input
	existingActivity.User_id = input.User_id
	existingActivity.Activity = input.Activity

	// 3. Kirim model yang sudah lengkap ke service
	updatedActivity, err := h.service.Update(existingActivity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedActivity)
}

func (h *Handler) DeleteActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Activity deleted successfully"})
}