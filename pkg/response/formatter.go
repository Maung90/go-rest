package response

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)


// SuksesResponse adalah struktur untuk respons 2xx
type SuksesResponse struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	Data      any    `json:"data,omitempty"`
	Timestamp string `json:"timestamp"`
}

// ErrorResponse adalah struktur untuk respons 4xx & 5xx
type ErrorResponse struct {
	Status    bool   `json:"status"`
	Message   string `json:"message"`
	Errors    any    `json:"errors,omitempty"`
	Timestamp string `json:"timestamp"`
}

// PaginatedResponse adalah struktur untuk data dengan paginasi
type PaginatedResponse struct {
	Status     bool   `json:"status"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Pagination any    `json:"pagination"`
	Timestamp  string `json:"timestamp"`
}

// --- FUNGSI INTERNAL ---

// sukses adalah basis untuk respons 2xx
func sukses(c *gin.Context, message string, data any, code int) {
	c.JSON(code, SuksesResponse{
		Status:    true,
		Message:   message,
		Data:      data,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

// error adalah basis untuk respons 4xx/5xx.
// Menggunakan AbortWithStatusJSON untuk menghentikan handler chain.
func sendError(c *gin.Context, message string, errors any, code int) {
	c.AbortWithStatusJSON(code, ErrorResponse{
		Status:    false,
		Message:   message,
		Errors:    errors,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	})
}

// --- Fungsi Publik untuk Sukses ---

// OK mengirim respons 200 OK
func OK(c *gin.Context, message string, data any) {
	sukses(c, message, data, http.StatusOK)
}

// Created mengirim respons 201 Created
func Created(c *gin.Context, message string, data any) {
	if message == "" {
		message = "Data berhasil dibuat."
	}
	sukses(c, message, data, http.StatusCreated)
}

// Paginated mengirim respons 200 OK dengan data paginasi
func Paginated(c *gin.Context, message string, data any, pagination any) {
	c.JSON(http.StatusOK, PaginatedResponse{
		Status:     true,
		Message:    message,
		Data:       data,
		Pagination: pagination,
		Timestamp:  time.Now().UTC().Format(time.RFC3339),
	})
}

// --- Fungsi Publik untuk Error ---

// ValidationError mengirim respons 422 Unprocessable Entity
func ValidationError(c *gin.Context, message string, errors any) {
	if message == "" {
		message = "Validasi gagal"
	}
	sendError(c, message, errors, http.StatusUnprocessableEntity)
}

// NotFound mengirim respons 404 Not Found
func NotFound(c *gin.Context, message string) {
	if message == "" {
		message = "Sumber daya tidak ditemukan"
	}
	sendError(c, message, nil, http.StatusNotFound)
}

// Unauthorized mengirim respons 401 Unauthorized
func Unauthorized(c *gin.Context, message string) {
	if message == "" {
		message = "Akses tidak sah"
	}
	sendError(c, message, nil, http.StatusUnauthorized)
}

// Forbidden mengirim respons 403 Forbidden
func Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = "Akses ditolak"
	}
	sendError(c, message, nil, http.StatusForbidden)
}

// HandleError mengirim respons 500 Internal Server Error
// Ini setara dengan fungsi handleError(Exception $e) Anda
func HandleError(c *gin.Context, err error, defaultMessage string) {
	// Log error untuk debugging
	log.Printf("API Error: %v\n", err)

	isDebug := os.Getenv("APP_DEBUG") == "true"
	var errorDetails any
	if isDebug {
		errorDetails = err.Error() // Hanya tampilkan pesan error jika debug
	}

	if defaultMessage == "" {
		defaultMessage = "Terjadi kesalahan internal. Silakan coba lagi nanti."
	}
	sendError(c, defaultMessage, errorDetails, http.StatusInternalServerError)
}