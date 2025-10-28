package router

import (
	"go-rest/internal/middleware"
	"go-rest/internal/habitLog"
	"github.com/gin-gonic/gin"
	"go-rest/internal/habit"
	"go-rest/internal/auth"
	"go-rest/internal/user"

)

func SetupRouter(
	userHandler *user.Handler, 
	habitHandler *habit.Handler, 
	authHandler *auth.Handler, 
	habitLogHandler *habitLog.Handler,
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	authRoutes := api.Group("auth") 
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
		// authRoutes.GET("/:email", authHandler.GetUserByEmail)
		authRoutes.POST("/refresh", authHandler.Refresh)
	}

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{
		userRoutes := protected.Group("/users")
		{
			userRoutes.GET("/", userHandler.GetUsers)
			userRoutes.GET("/:id", userHandler.GetUser)
			userRoutes.POST("/", userHandler.CreateUser)
			userRoutes.PUT("/:id", userHandler.UpdateUser)
			userRoutes.DELETE("/:id", userHandler.DeleteUser)
		}

		habitRoutes := protected.Group("habits") 
		{
			habitRoutes.GET("/", habitHandler.GetHabits)
			habitRoutes.POST("/history", habitLogHandler.GetLogsByDate)
			habitRoutes.GET("/:id", habitHandler.GetHabit)
			habitRoutes.POST("/", habitHandler.CreateHabit)
			habitRoutes.POST("/:id/complete", habitLogHandler.CreateLogs)
			habitRoutes.PUT("/:id", habitHandler.UpdateHabit)
			habitRoutes.DELETE("/:id", habitHandler.DeleteHabit)
		}
	}

	return router
}