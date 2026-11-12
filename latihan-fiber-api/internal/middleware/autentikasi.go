package middleware

import (
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"

	"latihan-fiber/internal/auth"
	"latihan-fiber/internal/config"
	"latihan-fiber/internal/model"
)

// Autentikasi memeriksa token JWT pada setiap permintaan.
func Autentikasi() fiber.Handler {
	cfg := config.LoadConfig()

	return func(c fiber.Ctx) error {
		header := c.Get("Authorization")

		if header == "" {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Token autentikasi tidak ditemukan",
			)
		}

		bagian := strings.SplitN(header, " ", 2)

		if len(bagian) != 2 ||
			!strings.EqualFold(bagian[0], "Bearer") ||
			strings.TrimSpace(bagian[1]) == "" {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Format token autentikasi tidak valid",
			)
		}

		// Mengambil token dari header Authorization.
		nilaiToken := strings.TrimSpace(bagian[1])

		// Memeriksa validitas token JWT.
		klaim, err := auth.PeriksaToken(
			nilaiToken,
			cfg.JwtSecret,
		)
		if err != nil {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Token tidak valid atau sudah kedaluwarsa",
			)
		}

		// Memeriksa apakah token sudah masuk blacklist.
		var tokenBlacklist model.TokenBlacklist
		err = config.DB.Where("token = ?", nilaiToken).
			First(&tokenBlacklist).Error

		if err == nil {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Token sudah dicabut dan tidak dapat digunakan",
			)
		}

		// Memastikan kesalahan database bukan karena data tidak ditemukan.
		if err != nil && err.Error() != "record not found" {
			return fiber.NewError(
				fiber.StatusInternalServerError,
				"Gagal memeriksa status token",
			)
		}

		// Mengambil ID pengguna dari subject token.
		idPengguna, err := strconv.ParseUint(klaim.Subject, 10, 64)
		if err != nil || idPengguna == 0 {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Identitas pengguna pada token tidak valid",
			)
		}

		// Menyimpan identitas pengguna dan token ke context.
		c.Locals("id_pengguna", uint(idPengguna))
		c.Locals("peran", klaim.Peran)
		c.Locals("token", nilaiToken)
		c.Locals("kadaluarsa_token", klaim.ExpiresAt.Time)

		return c.Next()
	}
}

// HanyaPeran membatasi akses berdasarkan peran pengguna.
func HanyaPeran(peranDiizinkan ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		peranPengguna, ok := c.Locals("peran").(string)
		if !ok || peranPengguna == "" {
			return fiber.NewError(
				fiber.StatusUnauthorized,
				"Identitas pengguna tidak ditemukan",
			)
		}

		for _, peran := range peranDiizinkan {
			if peranPengguna == peran {
				return c.Next()
			}
		}

		return fiber.NewError(
			fiber.StatusForbidden,
			"Anda tidak memiliki izin untuk mengakses sumber daya ini",
		)
	}
}
