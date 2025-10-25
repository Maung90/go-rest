package main

import (
	"go-rest/internal/activity" // <-- Tambahkan import activity
	"go-rest/internal/service"
	"go-rest/internal/router"
	"go-rest/internal/user"
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

	// Inisialisasi Repository seperti biasa
	userRepository := user.NewRepository(db)
	activityRepository := activity.NewRepository(db)

	// Buat instance Generic Service untuk setiap model
	// service.NewService[user.User](...) -> Membuat service khusus untuk tipe User
	userService := service.NewService[user.User](userRepository)
	// service.NewService[activity.Activity](...) -> Membuat service khusus untuk tipe Activity
	activityService := service.NewService[activity.Activity](activityRepository)

	// Inisialisasi Handler seperti biasa, tapi sekarang menerima Generic Service
	userHandler := user.NewHandler(userService)
	activityHandler := activity.NewHandler(activityService)

	// 2. Setup router dengan menyertakan userHandler dan activityHandler
	appRouter := router.SetupRouter(userHandler, activityHandler)
	
	// ==============================================================

	log.Println("Starting server on :8080")
	if err := appRouter.Run(":8080"); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}