package activity

import (
	// "strconv"
	// "database/sql"
	"go-rest/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllActivities(c *gin.Context) {
	userIDContext, exists := c.Get("userID") 

	if !exists {
		response.Unauthorized(c, "Silahkan login terlebih dahulu!")
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		response.Unauthorized(c, "Tipe user id tidak sesuai")
		return
	}

	story, err := h.service.FindAll(userID)

	if err != nil {
		response.HandleError(c, err, "Gagal mengambil cerita")
		// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response.OK(c, "Cerita berhasil diambil", story)
	// c.JSON(http.StatusOK, story)
}