package router

import (
	"go-rest/internal/activity"
	"go-rest/internal/user"

	"github.com/gin-gonic/gin"
)

// Perbaikan 1: Tambahkan activityHandler sebagai parameter
func SetupRouter(userHandler *user.Handler, activityHandler *activity.Handler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	// Rute untuk User
	userRoutes := api.Group("/users")
	{
		userRoutes.GET("/", userHandler.GetUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	// Perbaikan 2 & 3: Ubah komentar dan nama rute
	// Rute untuk Activity
	activityRoutes := api.Group("/activities") 
	{
		activityRoutes.GET("/", activityHandler.GetActivitys)
		activityRoutes.GET("/:id", activityHandler.GetActivity)
		activityRoutes.POST("/", activityHandler.CreateActivity)
		activityRoutes.PUT("/:id", activityHandler.UpdateActivity)
		activityRoutes.DELETE("/:id", activityHandler.DeleteActivity)
	}

	return router
}