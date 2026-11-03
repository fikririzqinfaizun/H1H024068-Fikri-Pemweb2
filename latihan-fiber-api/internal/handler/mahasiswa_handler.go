package handler

import (
	"sort"
	"strconv"

	"latihan-fiber/internal/model"

	"github.com/gofiber/fiber/v3"
)

// Menampilkan seluruh data mahasiswa
func GetMahasiswa(c fiber.Ctx) error {
	angkatan := c.Query("angkatan")
	urut := c.Query("urut")

	var hasil []model.Mahasiswa

	// Filter berdasarkan angkatan jika parameter tersedia
	if angkatan != "" {
		tahun, err := strconv.Atoi(angkatan)
		if err != nil {
			return fiber.NewError(
				fiber.StatusBadRequest,
				"Angkatan harus berupa angka",
			)
		}

		for _, mhs := range model.Daftar {
			if mhs.Angkatan == tahun {
				hasil = append(hasil, mhs)
			}
		}
	} else {
		hasil = append(hasil, model.Daftar...)
	}

	// Mengurutkan data berdasarkan nama atau IPK
	switch urut {
	case "nama":
		sort.Slice(hasil, func(i, j int) bool {
			return hasil[i].Nama < hasil[j].Nama
		})

	case "ipk":
		sort.Slice(hasil, func(i, j int) bool {
			return hasil[i].IPK < hasil[j].IPK
		})

	case "":
		// Tidak melakukan pengurutan tambahan

	default:
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Parameter urut hanya menerima nama atau ipk",
		)
	}

	return c.JSON(hasil)
}

// Menampilkan detail mahasiswa berdasarkan ID
func GetMahasiswaByID(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
	}

	for _, mhs := range model.Daftar {
		if mhs.ID == id {
			return c.JSON(mhs)
		}
	}

	return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
}

// Menampilkan detail mahasiswa berdasarkan NIM
func GetMahasiswaByNIM(c fiber.Ctx) error {
	nim := c.Params("nim")

	for _, mhs := range model.Daftar {
		if mhs.NIM == nim {
			return c.JSON(mhs)
		}
	}

	return fiber.NewError(
		fiber.StatusNotFound,
		"Mahasiswa dengan NIM tersebut tidak ditemukan",
	)
}
