package router

import (
	"go-rest/internal/middleware"
	"go-rest/internal/habitLog"
	"github.com/gin-gonic/gin"
	"go-rest/internal/habit"
	"go-rest/internal/sleep"
	"go-rest/internal/auth"
	"go-rest/internal/user"

)

func SetupRouter(
	authHandler *auth.Handler, 
	userHandler *user.Handler, 
	habitHandler *habit.Handler, 
	habitLogHandler *habitLog.Handler,
	sleepHandler *sleep.Handler, 
) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	authRoutes := api.Group("auth") 
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.POST("/refresh", authHandler.Refresh)
		authRoutes.POST("/forgot-password", authHandler.ForgotPassword)
		authRoutes.POST("/reset-password", authHandler.ResetPassword)
	}

	protected := api.Group("/")
	protected.Use(middleware.AuthMiddleware())
	{

		authRoutes := protected.Group("auth")
		{
			authRoutes.POST("/logout", authHandler.Logout)
			authRoutes.POST("/profile", authHandler.UpdateProfile)
			authRoutes.GET("/me", authHandler.GetProfile)
		}

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

		sleepRoutes := protected.Group("/sleeps")
		{
			sleepRoutes.GET("/", sleepHandler.GetSleeps)
			sleepRoutes.GET("/:id", sleepHandler.GetSleep)
			sleepRoutes.POST("/", sleepHandler.CreateSleep)
			sleepRoutes.PUT("/:id", sleepHandler.UpdateSleep)
			sleepRoutes.DELETE("/:id", sleepHandler.DeleteSleep)
		}
	}

	return router
}