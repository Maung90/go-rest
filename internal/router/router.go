package router

import (
	"go-rest/internal/habit"
	"go-rest/internal/auth"
	"go-rest/internal/user"

	"github.com/gin-gonic/gin"
)

func SetupRouter(userHandler *user.Handler, habitHandler *habit.Handler, authHandler *auth.Handler) *gin.Engine {
	router := gin.Default()

	api := router.Group("/api/v1")

	userRoutes := api.Group("/users")
	{
		userRoutes.GET("/", userHandler.GetUsers)
		userRoutes.GET("/:id", userHandler.GetUser)
		userRoutes.POST("/", userHandler.CreateUser)
		userRoutes.PUT("/:id", userHandler.UpdateUser)
		userRoutes.DELETE("/:id", userHandler.DeleteUser)
	}

	habitRoutes := api.Group("habits") 
	{
		habitRoutes.GET("/", habitHandler.GetHabits)
		habitRoutes.GET("/:id", habitHandler.GetHabit)
		habitRoutes.POST("/", habitHandler.CreateHabit)
		habitRoutes.PUT("/:id", habitHandler.UpdateHabit)
		habitRoutes.DELETE("/:id", habitHandler.DeleteHabit)
	}

	authRoutes := api.Group("auth") 
	{
		authRoutes.POST("/login", authHandler.Login)
		authRoutes.POST("/register", authHandler.Register)
		authRoutes.GET("/:email", authHandler.GetUserByEmail)
	}
	return router
}