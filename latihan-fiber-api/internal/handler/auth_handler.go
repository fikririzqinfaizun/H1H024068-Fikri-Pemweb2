package handler

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"latihan-fiber/internal/auth"
	"latihan-fiber/internal/config"
	"latihan-fiber/internal/model"
	"latihan-fiber/internal/validasi"
)

// Registrasi digunakan untuk mendaftarkan pengguna baru.
func Registrasi(c fiber.Ctx) error {
	var permintaan PermintaanRegistrasi

	// Mengambil data dari request body.
	err := c.Bind().Body(&permintaan)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Format body tidak valid",
		)
	}

	// Memvalidasi data registrasi.
	galat := validasi.Periksa(permintaan)
	if galat != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data registrasi tidak valid",
			"galat":  galat,
		})
	}

	// Memeriksa apakah email sudah terdaftar.
	var penggunaLama model.Pengguna
	err = config.DB.Where("email = ?", permintaan.Email).
		First(&penggunaLama).Error

	if err == nil {
		return fiber.NewError(
			fiber.StatusConflict,
			"Email sudah terdaftar",
		)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal memeriksa email pengguna",
		)
	}

	// Mengenkripsi kata sandi menggunakan bcrypt.
	hashKataSandi, err := bcrypt.GenerateFromPassword(
		[]byte(permintaan.KataSandi),
		bcrypt.DefaultCost,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal memproses kata sandi",
		)
	}

	// Membuat data pengguna baru.
	pengguna := model.Pengguna{
		Nama:      permintaan.Nama,
		Email:     permintaan.Email,
		KataSandi: string(hashKataSandi),
		Peran:     "mahasiswa",
	}

	// Menyimpan data pengguna ke database.
	err = config.DB.Create(&pengguna).Error
	if err != nil {
		return fiber.NewError(
			fiber.StatusUnprocessableEntity,
			"Gagal menyimpan data pengguna",
		)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Registrasi berhasil",
		"data":   pengguna,
	})
}

// Login digunakan untuk melakukan autentikasi pengguna.
func Login(c fiber.Ctx) error {
	var permintaan PermintaanLogin

	// Mengambil data dari request body.
	err := c.Bind().Body(&permintaan)
	if err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"Format body tidak valid",
		)
	}

	// Memvalidasi data login.
	galat := validasi.Periksa(permintaan)
	if galat != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"sukses": false,
			"pesan":  "Data login tidak valid",
			"galat":  galat,
		})
	}

	// Mencari pengguna berdasarkan email.
	var pengguna model.Pengguna
	err = config.DB.Where("email = ?", permintaan.Email).
		First(&pengguna).Error

	if err != nil {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Email atau kata sandi tidak sesuai",
		)
	}

	// Memeriksa kecocokan kata sandi.
	err = bcrypt.CompareHashAndPassword(
		[]byte(pengguna.KataSandi),
		[]byte(permintaan.KataSandi),
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Email atau kata sandi tidak sesuai",
		)
	}

	// Mengambil konfigurasi JWT.
	cfg := config.LoadConfig()

	// Membuat token JWT.
	token, err := auth.TerbitkanToken(
		pengguna.ID,
		pengguna.Peran,
		cfg.JwtSecret,
		cfg.JwtTtlMenit,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal membuat token autentikasi",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Login berhasil",
		"data": fiber.Map{
			"token":         token,
			"tipe_token":    "Bearer",
			"berlaku_menit": cfg.JwtTtlMenit,
			"pengguna":      pengguna,
		},
	})
}

// Profil digunakan untuk mengambil profil pengguna yang sedang login.
func Profil(c fiber.Ctx) error {
	// Mengambil ID pengguna dari middleware autentikasi.
	idPengguna, ok := c.Locals("id_pengguna").(uint)
	if !ok || idPengguna == 0 {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Pengguna belum terautentikasi",
		)
	}

	// Mengambil data pengguna dari database.
	var pengguna model.Pengguna
	err := config.DB.First(&pengguna, idPengguna).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return fiber.NewError(
				fiber.StatusNotFound,
				"Pengguna tidak ditemukan",
			)
		}

		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mengambil profil pengguna",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Profil pengguna berhasil diambil",
		"data":   pengguna,
	})
}

// SegarkanToken digunakan untuk menerbitkan JWT baru
// berdasarkan token yang masih berlaku.
func SegarkanToken(c fiber.Ctx) error {
	// Mengambil ID pengguna dari middleware autentikasi.
	idPengguna, ok := c.Locals("id_pengguna").(uint)
	if !ok || idPengguna == 0 {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Pengguna belum terautentikasi",
		)
	}

	// Mengambil peran pengguna dari middleware autentikasi.
	peran, ok := c.Locals("peran").(string)
	if !ok || peran == "" {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Peran pengguna tidak ditemukan",
		)
	}

	// Memuat konfigurasi JWT.
	cfg := config.LoadConfig()

	// Menerbitkan token JWT baru.
	tokenBaru, err := auth.TerbitkanToken(
		idPengguna,
		peran,
		cfg.JwtSecret,
		cfg.JwtTtlMenit,
	)
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal menerbitkan token baru",
		)
	}

	// Mengembalikan token baru kepada pengguna.
	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Token berhasil diperbarui",
		"data": fiber.Map{
			"token":         tokenBaru,
			"tipe_token":    "Bearer",
			"berlaku_menit": cfg.JwtTtlMenit,
		},
	})
}

// Logout digunakan untuk mencabut token JWT sebelum masa berlakunya habis.
func Logout(c fiber.Ctx) error {
	// Mengambil token dari middleware autentikasi.
	token, ok := c.Locals("token").(string)
	if !ok || token == "" {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Token autentikasi tidak ditemukan",
		)
	}

	// Mengambil waktu kedaluwarsa token.
	kadaluarsa, ok := c.Locals("kadaluarsa_token").(time.Time)
	if !ok || kadaluarsa.IsZero() {
		return fiber.NewError(
			fiber.StatusUnauthorized,
			"Masa berlaku token tidak ditemukan",
		)
	}

	// Menyimpan token ke tabel blacklist.
	dataBlacklist := model.TokenBlacklist{
		Token:      token,
		Kadaluarsa: kadaluarsa,
	}

	err := config.DB.Create(&dataBlacklist).Error
	if err != nil {
		return fiber.NewError(
			fiber.StatusInternalServerError,
			"Gagal mencabut token autentikasi",
		)
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"pesan":  "Logout berhasil. Token tidak dapat digunakan kembali",
	})
}
