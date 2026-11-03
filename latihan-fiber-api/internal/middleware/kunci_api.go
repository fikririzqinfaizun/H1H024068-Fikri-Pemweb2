package middleware

import (
	"os"

	"github.com/gofiber/fiber/v3"
)

// Middleware untuk memeriksa API Key
func KunciApi() fiber.Handler {
	return func(c fiber.Ctx) error {
		apiKey := os.Getenv("API_KEY")

		// Memastikan API Key sudah dikonfigurasi
		if apiKey == "" {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"API Key belum dikonfigurasi",
			)
		}

		// Memeriksa header X-Api-Key
		kunci := c.Get("X-Api-Key")

		if kunci != apiKey {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"API Key tidak valid",
			)
		}

		// Melanjutkan permintaan jika API Key sesuai
		return c.Next()
	}
}
