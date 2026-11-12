package router

import (
	"latihan-fiber/internal/handler"
	"latihan-fiber/internal/middleware"

	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	// =====================================================
	// ENDPOINT PUBLIK
	// =====================================================

	// Endpoint status API
	app.Get("/api/status", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "berjalan",
			"pesan":  "API Fiber berhasil dijalankan",
		})
	})

	// Endpoint registrasi pengguna
	app.Post("/api/auth/registrasi", handler.Registrasi)

	// Endpoint login pengguna
	app.Post("/api/auth/login", handler.Login)

	// =====================================================
	// ENDPOINT AUTENTIKASI JWT
	// =====================================================

	// Grup endpoint yang memerlukan token JWT
	autentikasi := app.Group("/api", middleware.Autentikasi())

	// Endpoint profil pengguna yang sedang login
	autentikasi.Get("/auth/profil", handler.Profil)

	// Endpoint memperbarui token JWT
	autentikasi.Post("/auth/segarkan", handler.SegarkanToken)

	// Endpoint logout untuk mencabut token JWT
	autentikasi.Post("/auth/logout", handler.Logout)

	// =====================================================
	// ENDPOINT MAHASISWA
	// =====================================================

	// Seluruh endpoint mahasiswa memerlukan JWT
	mahasiswa := autentikasi.Group("/mahasiswa")

	// Endpoint daftar mahasiswa
	mahasiswa.Get("/", handler.DaftarMahasiswa)

	// Endpoint detail mahasiswa berdasarkan ID
	mahasiswa.Get("/:id", handler.DetailMahasiswa)

	// Endpoint tambah mahasiswa: admin dan dosen
	mahasiswa.Post(
		"/",
		middleware.HanyaPeran("admin", "dosen"),
		handler.BuatMahasiswa,
	)

	// Endpoint memperbarui mahasiswa: admin dan dosen
	mahasiswa.Put(
		"/:id",
		middleware.HanyaPeran("admin", "dosen"),
		handler.PerbaruiMahasiswa,
	)

	// Endpoint menghapus mahasiswa: hanya admin
	mahasiswa.Delete(
		"/:id",
		middleware.HanyaPeran("admin"),
		handler.HapusMahasiswa,
	)

	// Endpoint memulihkan mahasiswa: hanya admin
	mahasiswa.Post(
		"/:id/pulihkan",
		middleware.HanyaPeran("admin"),
		handler.PulihkanMahasiswa,
	)

	// =====================================================
	// ENDPOINT MATA KULIAH
	// =====================================================

	// Tetap menggunakan middleware API Key
	matakuliah := app.Group(
		"/api/matakuliah",
		middleware.KunciApi(),
	)

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

	// =====================================================
	// ENDPOINT PROGRAM STUDI
	// =====================================================

	// Tetap menggunakan middleware API Key
	programStudi := app.Group(
		"/api/program-studi",
		middleware.KunciApi(),
	)

	// Endpoint daftar program studi beserta jumlah mahasiswa
	programStudi.Get("/", handler.DaftarProgramStudi)

	// =====================================================
	// PENANGANAN ENDPOINT TIDAK DITEMUKAN
	// =====================================================

	app.Use(func(c fiber.Ctx) error {
		return fiber.NewError(
			fiber.StatusNotFound,
			"Endpoint tidak ditemukan",
		)
	})
}
