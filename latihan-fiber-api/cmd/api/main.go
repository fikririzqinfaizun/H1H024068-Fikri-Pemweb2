package main

import (
	"log"

	"latihan-fiber/internal/config"
	"latihan-fiber/internal/router"

	"github.com/gofiber/fiber/v3"
)

func main() {
	// Memuat konfigurasi aplikasi
	cfg := config.LoadConfig()

	// Menghubungkan aplikasi ke database MySQL
	config.KoneksiDatabase(cfg)

	// Menjalankan migrasi database
	config.Migrasi()

	// Memasukkan data awal program studi
	config.DataAwal()

	// Membuat aplikasi Fiber
	app := fiber.New()

	// Mengatur routing API
	router.SetupRoutes(app)

	// Menjalankan server
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
