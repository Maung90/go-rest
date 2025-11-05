package activity

// import (
// 	"net/http"
// 	"strconv"
// 	"database/sql"
// 	"github.com/gin-gonic/gin"
// )

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}