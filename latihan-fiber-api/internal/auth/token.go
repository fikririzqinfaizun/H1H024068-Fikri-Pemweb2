package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Klaim struct {
	Peran string `json:"peran"`
	jwt.RegisteredClaims
}

// TerbitkanToken membuat JWT untuk pengguna yang berhasil login.
func TerbitkanToken(
	idPengguna uint,
	peran string,
	rahasia string,
	ttlMenit int,
) (string, error) {
	if rahasia == "" {
		return "", errors.New("JWT secret tidak boleh kosong")
	}

	if ttlMenit <= 0 {
		return "", errors.New("masa berlaku token harus lebih dari nol")
	}

	waktuSekarang := time.Now()

	klaim := Klaim{
		Peran: peran,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(idPengguna), 10),
			Issuer:    "pemweb2-fiber-api",
			IssuedAt:  jwt.NewNumericDate(waktuSekarang),
			ExpiresAt: jwt.NewNumericDate(waktuSekarang.Add(time.Duration(ttlMenit) * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, klaim)

	return token.SignedString([]byte(rahasia))
}

// PeriksaToken memvalidasi JWT dan mengembalikan klaimnya.
func PeriksaToken(nilaiToken string, rahasia string) (*Klaim, error) {
	if nilaiToken == "" || rahasia == "" {
		return nil, errors.New("token atau secret tidak boleh kosong")
	}

	token, err := jwt.ParseWithClaims(
		nilaiToken,
		&Klaim{},
		func(token *jwt.Token) (any, error) {
			if token.Method != jwt.SigningMethodHS256 {
				return nil, errors.New("metode tanda tangan token tidak valid")
			}

			return []byte(rahasia), nil
		},
	)

	if err != nil {
		return nil, err
	}

	klaim, ok := token.Claims.(*Klaim)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return klaim, nil
}
