package main

import (
	"go-rest/internal/dailyStory"
	"go-rest/internal/habitLog"
	"go-rest/internal/service"
	"go-rest/internal/router"
	"go-rest/internal/habit"
	"go-rest/internal/sleep"
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
	userService := service.NewService[user.User](userRepository)
	userHandler := user.NewHandler(userService)

	sleepRepository := sleep.NewRepository(db)
	sleepService := sleep.NewService(sleepRepository)
	sleepHandler := sleep.NewHandler(sleepService)

	authRepository := auth.NewRepository(db)
	authService := auth.NewAuthService(authRepository)
	authHandler := auth.NewHandler(authService, userService)

	habitRepository := habit.NewRepository(db)
	habitService := service.NewService[habit.Habit](habitRepository)
	habitHandler := habit.NewHandler(habitService)

	habitLogRepository := habitLog.NewRepository(db)
	habitLogService := habitLog.NewHabitLogService(habitLogRepository)
	habitLogHandler := habitLog.NewHandler(habitLogService, habitService)

	dailyStoryRepository := dailyStory.NewRepository(db)
	dailyStoryService := dailyStory.NewService(dailyStoryRepository)
	dailyStoryHandler := dailyStory.NewHandler(dailyStoryService)

	appRouter := router.SetupRouter(
		authHandler,	
		userHandler, 
		habitHandler,
		habitLogHandler, 
		sleepHandler,
		dailyStoryHandler,
	)


	log.Println("Starting server on :9001")
	if err := appRouter.Run(":9001"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}