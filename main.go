package main

import (
	"context"
	"go-rest/internal/config"
	"go-rest/internal/connection"
	"go-rest/internal/repository" // Impor repository
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// 1. Muat Konfigurasi
	cnf := config.Get()
	log.Printf("Konfigurasi dimuat: %+v", cnf.Server)

	// 2. Buat Koneksi Database
	dbConnection, err := connection.GetDatabase(cnf.Database)
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}
	defer dbConnection.Close()
	log.Println("Berhasil terhubung ke database!")

	// 3. Inisialisasi Repository dengan koneksi DB
	customerRepo := repository.NewCustomer(dbConnection)

	// 4. Inisialisasi Fiber App
	app := fiber.New()

	// 5. Buat Handler baru yang menggunakan Repository
	customerHandler := func(c *fiber.Ctx) error {
		// Gunakan repository untuk mengambil semua customer
		customers, err := customerRepo.FindAll(context.Background())
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "gagal mengambil data customer",
			})
		}
		return c.Status(fiber.StatusOK).JSON(customers)
	}

	// 6. Daftarkan Route ke Handler yang baru
	app.Get("/customers", customerHandler)
	app.Get("/developers", developers) // route lama tetap ada

	// 7. Jalankan Server
	serverAddress := cnf.Server.Host + ":" + cnf.Server.Port
	log.Printf("Server berjalan di http://%s", serverAddress)
	app.Listen(serverAddress)
}

// Handler lama sebagai contoh
func developers(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{
		"message": "success",
	})
}
