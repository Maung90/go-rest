package dailyStory

import (
	"strconv"
	"database/sql"
	"go-rest/pkg/response"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) GetAllStories(c *gin.Context) {
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

func (h *Handler) GetStoriesByDate(c *gin.Context) {
	date := c.Param("date") //ambil param dari url
	userIDContext, exists := c.Get("userID") // ambil dari middleware

	if !exists {
		response.Unauthorized(c, "User tidak ditemukan, pastikan login terlebih dahulu.")
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		response.Unauthorized(c, "User tidak ditemukan di context.")
		return
	}

	stories, err := h.service.FindByDate(userID, date)

	if err != nil {
		response.HandleError(c, err, "Gagal Mengambil cerita")
		return
	}
	response.Created(c, "Cerita berhasil Diambil", stories)
}

func (h *Handler) CreateStories(c *gin.Context) {
	userIDContext, exists := c.Get("userID") 

	if !exists {
		response.Unauthorized(c, "User tidak ditemukan, pastikan login terlebih dahulu.")
		return
	}
	userID, ok := userIDContext.(int)
	if !ok {
		response.Unauthorized(c, "User tidak ditemukan di context.")
		return
	}
	
	var input StoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Cerita gagal disimpan, data yang diinputkan tidak valid", err)
		return
	}

	story, err := h.service.Save(userID, input)

	if err != nil {
		response.HandleError(c, err, "Cerita gagal disimpan")
		return
	}
	response.Created(c, "Cerita berhasil disimpan", story)
}

func (h *Handler) UpdateStories(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		response.ValidationError(c, "Data yang dikirim tidak valid", err.Error())
		return
	}

	userIDContext, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "User ID tidak ditemukan di context.")
		return
	}
	userID := userIDContext.(int)
	
	var input StoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Data yang diinputkan tidak valid", err)
		return
	}

	story, err := h.service.Update(id, userID, input)
	if err != nil {
		if err == sql.ErrNoRows {
			response.NotFound(c, "Data tidak ditemukan atau user tidak mempunyai ijin")
			return
		}
		response.HandleError(c,err, "Data gagal di update")
		return
	}
	response.OK(c, "Data berhasil diupdate", story)
}

func (h *Handler) DeleteStories(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
			response.NotFound(c, "Data tidak ditemukan!")
		return
	}
	response.OK(c, "Data berhasil Dihapus", "")
}