package repository

import (
	"latihan-fiber/internal/config"
	"latihan-fiber/internal/model"
)

type ProgramStudiDenganJumlah struct {
	model.ProgramStudi
	JumlahMahasiswa int64 `json:"jumlah_mahasiswa"`
}

func AmbilSemuaProgramStudi() ([]ProgramStudiDenganJumlah, error) {
	var daftar []ProgramStudiDenganJumlah

	err := config.DB.
		Model(&model.ProgramStudi{}).
		Select("program_studis.*, COUNT(mahasiswas.id) AS jumlah_mahasiswa").
		Joins("LEFT JOIN mahasiswas ON mahasiswas.program_studi_id = program_studis.id").
		Group("program_studis.id").
		Order("program_studis.nama ASC").
		Scan(&daftar).Error

	return daftar, err
}
