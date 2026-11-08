package handler

import (
	"errors"
	"strconv"

	"latihan-fiber/internal/model"
	"latihan-fiber/internal/repository"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func DaftarMatakuliah(c fiber.Ctx) error {
	daftar, err := repository.AmbilSemuaMatakuliah()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   daftar,
	})
}

func DetailMatakuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID tidak valid",
		})
	}

	matakuliah, err := repository.AmbilMatakuliahPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data mata kuliah tidak ditemukan",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   matakuliah,
	})
}

func BuatMatakuliah(c fiber.Ctx) error {
	var matakuliah model.Matakuliah

	if err := c.Bind().Body(&matakuliah); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format data tidak valid",
		})
	}

	if err := repository.SimpanMatakuliah(&matakuliah); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menyimpan data mata kuliah",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mata kuliah berhasil dibuat",
		"data":   matakuliah,
	})
}

func PerbaruiMatakuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID tidak valid",
		})
	}

	matakuliah, err := repository.AmbilMatakuliahPerId(uint(id))
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data mata kuliah tidak ditemukan",
		})
	}
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal mengambil data mata kuliah",
		})
	}

	if err := c.Bind().Body(&matakuliah); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Format data tidak valid",
		})
	}

	if err := repository.PerbaruiMatakuliah(&matakuliah); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal memperbarui data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mata kuliah berhasil diperbarui",
		"data":   matakuliah,
	})
}

func HapusMatakuliah(c fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 32)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "ID tidak valid",
		})
	}

	err = repository.HapusMatakuliah(uint(id))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Gagal menghapus data mata kuliah",
		})
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Data mata kuliah berhasil dihapus",
	})
}
