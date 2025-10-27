package main

import (
	"go-rest/internal/habitLog"
	"go-rest/internal/service"
	"go-rest/internal/router"
	"go-rest/internal/habit"
	"go-rest/internal/user"
	"go-rest/internal/auth"
	"go-rest/pkg/database"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.ConnectDB()
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	userRepository := user.NewRepository(db)
	habitRepository := habit.NewRepository(db)
	authRepository := auth.NewRepository(db)
	habitLogRepository := habitLog.NewRepository(db)

	authService := auth.NewService(authRepository)
	habitLogService := habitLog.NewHabitLogService(habitLogRepository)
	userService := service.NewService[user.User](userRepository)
	habitService := service.NewService[habit.Habit](habitRepository)

	userHandler := user.NewHandler(userService)
	authHandler := auth.NewHandler(authService)
	habitLogHandler := habitLog.NewHandler(habitLogService, habitService)
	habitHandler := habit.NewHandler(habitService)

	appRouter := router.SetupRouter(userHandler, habitHandler, authHandler, habitLogHandler)
	
	
	log.Println("Starting server on :8080")
	if err := appRouter.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}