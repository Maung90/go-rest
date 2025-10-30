package main

import (
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
	sleepRepository := sleep.NewRepository(db)
	habitRepository := habit.NewRepository(db)
	authRepository := auth.NewRepository(db)
	habitLogRepository := habitLog.NewRepository(db)

	authService := auth.NewAuthService(authRepository)
	habitLogService := habitLog.NewHabitLogService(habitLogRepository)
	userService := service.NewService[user.User](userRepository)
	sleepService := service.NewService[sleep.Sleep](sleepRepository)
	habitService := service.NewService[habit.Habit](habitRepository)

	userHandler := user.NewHandler(userService)
	authHandler := auth.NewHandler(authService, userService)
	habitLogHandler := habitLog.NewHandler(habitLogService, habitService)
	habitHandler := habit.NewHandler(habitService)
	sleepHandler := sleep.NewHandler(sleepService)

	appRouter := router.SetupRouter(
		authHandler,	
		userHandler, 
		habitHandler,
		habitLogHandler, 
		sleepHandler,
	)


	log.Println("Starting server on :9001")
	if err := appRouter.Run(":9001"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}