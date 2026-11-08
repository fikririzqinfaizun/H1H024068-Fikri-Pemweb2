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
	mahasiswa.Get("/", handler.DaftarMahasiswa)

	// Endpoint detail mahasiswa berdasarkan ID
	mahasiswa.Get("/:id", handler.DetailMahasiswa)

	// Endpoint tambah mahasiswa
	mahasiswa.Post("/", handler.BuatMahasiswa)

	// Endpoint memperbarui mahasiswa
	mahasiswa.Put("/:id", handler.PerbaruiMahasiswa)

	// Endpoint menghapus mahasiswa
	mahasiswa.Delete("/:id", handler.HapusMahasiswa)

	// Endpoint memulihkan mahasiswa yang sudah dihapus
	mahasiswa.Post("/:id/pulihkan", handler.PulihkanMahasiswa)

	// Grup endpoint mata kuliah yang dilindungi API Key
	matakuliah := app.Group("/api/matakuliah", middleware.KunciApi())

	// Endpoint daftar mata kuliah
	matakuliah.Get("/", handler.DaftarMatakuliah)

	// Endpoint detail mata kuliah berdasarkan ID
	matakuliah.Get("/:id", handler.DetailMatakuliah)

	// Endpoint tambah mata kuliah
	matakuliah.Post("/", handler.BuatMatakuliah)

	// Endpoint memperbarui mata kuliah
	matakuliah.Put("/:id", handler.PerbaruiMatakuliah)

	// Endpoint menghapus mata kuliah
	matakuliah.Delete("/:id", handler.HapusMatakuliah)

	// Grup endpoint program studi yang dilindungi API Key
	programStudi := app.Group("/api/program-studi", middleware.KunciApi())

	// Endpoint daftar program studi beserta jumlah mahasiswa
	programStudi.Get("/", handler.DaftarProgramStudi)

	// Catch-all untuk endpoint yang tidak ditemukan
	app.Use(func(c fiber.Ctx) error {
		return fiber.NewError(
			fiber.StatusNotFound,
			"Endpoint tidak ditemukan",
		)
	})
}
