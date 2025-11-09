package activity

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

	activity, err := h.service.FindAll(userID)

	if err != nil {
		response.HandleError(c, err, "Gagal mengambil aktivitas")
		return
	}
	response.OK(c, "Aktivitas berhasil diambil", activity)
}

func (h *Handler) GetActvitiesById(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id")) //ambil param dari url
	if err != nil {
		response.ValidationError(c, "Id tidak valid", err.Error())
		return 
	}

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

	activities, err := h.service.FindById(userID, id)

	if err != nil {
		response.HandleError(c, err, "Gagal Mengambil aktivitas")
		return
	}
	response.Created(c, "Aktivitas berhasil Diambil", activities)
}

func (h *Handler) SaveActivities(c *gin.Context){
	userIDContext, exists := c.Get("userID")
	if !exists {
		response.Unauthorized(c, "User tidak ditemukan, pastikan login terlebih dahulu")
		return 
	}
	userID, ok := userIDContext.(int)
	if !ok{
		response.Unauthorized(c, "Tipe user id tidak valid")
		return 
	}

	var input ActivityInput
	if err := c.ShouldBindJSON(&input); err != nil{
		response.ValidationError(c, "Activities gagal disimpan, data yang diinputkan tidak valid", err)
		return 
	}

	activity, err := h.service.Save(userID, input)
	if err != nil {
		response.HandleError(c, err, "Gagal menyimpan aktivitas")
		return 
	}
	response.Created(c, "Berhasil menyimpan aktivitas", activity)
}

func (h *Handler) UpdateActivity(c *gin.Context) {
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
	
	var input ActivityInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.ValidationError(c, "Data yang diinputkan tidak valid", err)
		return
	}

	activity, err := h.service.Update(id, userID, input)
	if err != nil {
		if err == sql.ErrNoRows {
			response.NotFound(c, "Data tidak ditemukan atau user tidak mempunyai ijin")
			return
		}
		response.HandleError(c,err, "Data gagal di update")
		return
	}
	response.OK(c, "Data berhasil diupdate", activity)
}

func (h *Handler) DeleteActivity(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.service.Delete(id)
	if err != nil {
		response.NotFound(c, "Data tidak ditemukan!")
		return
	}
	response.OK(c, "Data berhasil Dihapus", "")
}