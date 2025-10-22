// main.go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"go-rest/internal/config"
	"go-rest/internal/connection"
	// "strconv"
)

func main() {
	cnf := config.Get()

	// Panggil GetDatabase dan terima DUA nilai balikan
	dbConnection, err := connection.GetDatabase(cnf.Database)
	// Periksa apakah ada error saat koneksi
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// Anda bisa menutup koneksi di sini menggunakan defer
	// Ini akan dieksekusi saat fungsi main selesai
	defer dbConnection.Close()

	app := fiber.New()

	app.Get("/developers", developers)

	serverAddress := ":" + cnf.Server.Port
	log.Printf("Server berjalan di port %d", cnf.Server.Port)
	app.Listen(serverAddress)
}

func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}
