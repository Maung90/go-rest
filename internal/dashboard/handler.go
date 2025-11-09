package dashboard

import (

	"go-rest/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetDashboard(c *gin.Context) {
	userIDContext, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "Silahkan login terlebih dahulu!")
		return
	}

	userID, ok := userIDContext.(int)
	if !ok {
		response.Unauthorized(c, "Tipe user ID tidak sesuai.")
		return
	}

	dateStr := c.Query("date")
	if dateStr == "" {
		response.ValidationError(c, "Parameter 'date' wajib diisi, format: YYYY-MM-DD", nil)
		return
	}

	// date, err := parser.ParseDateString(dateStr)
	// if err != nil {
	// 	response.ValidationError(c, "Format tanggal tidak valid. Gunakan format YYYY-MM-DD", err.Error())
	// 	return
	// }

	dashboardData, err := h.service.GetDashboard(userID, dateStr)
	if err != nil {
		response.HandleError(c, err, "Gagal mengambil data dashboard")
		return
	}

	response.OK(c, "Dashboard berhasil diambil", dashboardData)
}
