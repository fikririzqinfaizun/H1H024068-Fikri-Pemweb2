package handler

import (
	"latihan-fiber/internal/repository"

	"github.com/gofiber/fiber/v3"
)

func DaftarProgramStudi(c fiber.Ctx) error {
	daftar, err := repository.AmbilSemuaProgramStudi()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data program studi",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
	})
}
