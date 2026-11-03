package router

import (
	"latihan-fiber/internal/handler"
	"latihan-fiber/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	// Endpoint status API
	app.Get("/api/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "berjalan",
			"pesan":  "API Fiber berhasil dijalankan",
		})
	})

	// Grup endpoint mahasiswa yang dilindungi API Key
	mahasiswa := app.Group("/api/mahasiswa", middleware.KunciApi())

	// Endpoint daftar mahasiswa
	mahasiswa.Get("/", handler.GetMahasiswa)

	// Endpoint pencarian mahasiswa berdasarkan NIM
	mahasiswa.Get("/nim/:nim", handler.GetMahasiswaByNIM)

	// Endpoint detail mahasiswa berdasarkan ID
	mahasiswa.Get("/:id", handler.GetMahasiswaByID)

	// Catch-all untuk endpoint yang tidak ditemukan
	app.Use(func(c fiber.Ctx) error {
		return fiber.NewError(
			fiber.StatusNotFound,
			"Endpoint tidak ditemukan",
		)
	})
}
