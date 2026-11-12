package main

import (
	"errors"
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

	// Membuat aplikasi Fiber dengan Global Error Handler
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c fiber.Ctx, err error) error {
			// Status default untuk galat yang tidak dikenali
			statusCode := fiber.StatusInternalServerError
			pesan := "Terjadi kesalahan internal pada server"

			// Memeriksa apakah galat berasal dari Fiber
			var fiberErr *fiber.Error
			if errors.As(err, &fiberErr) {
				statusCode = fiberErr.Code

				// Galat 4xx menggunakan pesan yang sesuai
				// Galat 5xx tidak menampilkan detail teknis
				if statusCode < fiber.StatusInternalServerError {
					pesan = fiberErr.Message
				}
			}

			// Mencatat detail galat internal di terminal server
			if statusCode >= fiber.StatusInternalServerError {
				log.Printf("Galat internal: %v", err)
			}

			// Mengirim respons JSON terstruktur
			return c.Status(statusCode).JSON(fiber.Map{
				"sukses": false,
				"pesan":  pesan,
			})
		},
	})

	// Mengatur routing API
	router.SetupRoutes(app)

	// Menjalankan server
	log.Fatal(app.Listen(":" + cfg.AppPort))
}
