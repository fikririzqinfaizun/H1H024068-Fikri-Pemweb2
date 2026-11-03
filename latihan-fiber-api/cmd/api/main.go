package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"github.com/gofiber/fiber/v3/middleware/requestid"

	"latihan-fiber/internal/config"
	"latihan-fiber/internal/middleware"
	"latihan-fiber/internal/router"
)

func main() {
	// Memuat konfigurasi aplikasi
	cfg := config.LoadConfig()

	// Membuat instance Fiber
	app := fiber.New(fiber.Config{
		AppName: cfg.AppName,
	})

	// Middleware bawaan Fiber
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(requestid.New())

	// Middleware buatan sendiri
	app.Use(middleware.WaktuProses)

	// Global error handler
	app.Use(func(c fiber.Ctx) error {
		err := c.Next()

		if err != nil {
			code := fiber.StatusInternalServerError

			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}

			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return nil
	})

	// Mendaftarkan seluruh route
	router.SetupRoutes(app)

	// Menjalankan server
	log.Printf("%s berjalan pada port %s", cfg.AppName, cfg.AppPort)

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
