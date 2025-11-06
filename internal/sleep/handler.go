package sleep

import (
	"strconv"
	"time"
	"go-rest/pkg/response"
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
		response.Unauthorized(c, "User ID tidak ditemukan di context.")
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		response.Unauthorized(c, "Tipe user id tidak sesuai")
		return
	}
	sleeps, err := h.service.GetByUserID(userID)
	if err != nil {
		response.HandleError(c, err, "Gagal mengambil data tidur")
		return
	}
	response.OK(c, "Data tidur berhasil diambil", sleeps)
}
func (h *Handler) GetSleep(c *gin.Context) {
	userIDContext, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "User ID tidak ditemukan atau autentikasi gagal")
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		response.Unauthorized(c, "tipe user ID tidak sesuai")
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	sleep, err := h.service.GetByID(userID, id)
	if err != nil {
		response.NotFound(c, "Data tidur gagal diambil")
		return
	}
	response.OK(c, "Data tidur berhasil diambil", sleep)
}

func (h *Handler) CreateSleep(c *gin.Context) {
	var input SleepInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Input tidak sesuai ", err)
		return
	}

	layout := "2006-01-02 15:04:05"

	startTime, err := time.Parse(layout, input.SleepStart)
	if err != nil {
		response.HandleError(c, err, "format sleep_start salah")
		return
	}

	endTime, err := time.Parse(layout, input.SleepEnd)
	if err != nil {
		response.HandleError(c, err, "format sleep_end salah")
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
		response.HandleError(c, err, "Gagal menyimpan waktu tidur")
		return
	}
	response.Created(c, "Berhasil menyimpan data tidur", newSleep)
}

func (h *Handler) UpdateSleep(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	var input SleepInput
	if err := c.ShouldBindJSON(&input); err != nil {
	response.ValidationError(c,  "Data yang diinputkan tidak valid", err)
		return
	} 

	existingSleep, err := h.service.GetByID(input.User_id, id)
	if err != nil {
		response.NotFound(c, "User id tidak ditemukan")
		return
	}

	layout := "2006-01-02 15:04:05"

	startTime, err := time.Parse(layout, input.SleepStart)
	if err != nil {
		response.HandleError(c, err, "format sleep_start salah")
		return
	}

	endTime, err := time.Parse(layout, input.SleepEnd)
	if err != nil {
		response.HandleError(c, err, "format sleep_end salah")
		return
	}
	duration := endTime.Sub(startTime).Hours()

	existingSleep.User_id = input.User_id
	existingSleep.SleepStart = startTime
	existingSleep.SleepEnd = endTime
	existingSleep.Duration = duration

	updatedSleep, err := h.service.Update(existingSleep)
	if err != nil {
	response.HandleError(c, err, "Gagal update data tidur")
		return
	}
	response.OK(c, "Data tidur berhasil diupdate", updatedSleep)
}

func (h *Handler) DeleteSleep(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		response.NotFound(c, "Data tidur tidak ditemukan atau error saat menghapus")
		return
	}
	response.OK(c, "Data tidur berhasil di hapus", "")
}

func (h *Handler) GetWeeklyStats(c *gin.Context) {
	userIDContext, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "user id tidak ditemukan atau terjadi kesalahaan saat autentikasi")
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		response.Unauthorized(c, "Tipe user id tidak sesuai")
		return
	}
	stats, err := h.service.GetWeeklyStats(userID)
	if err != nil {
		response.HandleError(c, err, "Data statistik gagal di ambil")
		return
	}
	response.OK(c, "Data statistik berhasil di ambil", stats)
}

func (h *Handler) GetMonthlyStats(c *gin.Context) {
	userIDContext, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "user id tidak ditemukan atau terjadi kesalahaan saat autentikasi")
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		response.Unauthorized(c, "Tipe user id tidak sesuai")
		return
	}
	stats, err := h.service.GetMonthlyStats(userID)
	if err != nil {
		response.HandleError(c, err, "Data statistik gagal di ambil")
		return
	}
	response.OK(c, "Data statistik berhasil di ambil", stats)
}
