package repository

import (
	"latihan-fiber/internal/config"
	"latihan-fiber/internal/model"
)

func AmbilSemuaMatakuliah() ([]model.Matakuliah, error) {
	var daftar []model.Matakuliah

	err := config.DB.Order("nama ASC").Find(&daftar).Error

	return daftar, err
}

func AmbilMatakuliahPerId(id uint) (model.Matakuliah, error) {
	var matakuliah model.Matakuliah

	err := config.DB.First(&matakuliah, id).Error

	return matakuliah, err
}

func SimpanMatakuliah(matakuliah *model.Matakuliah) error {
	return config.DB.Create(matakuliah).Error
}

func PerbaruiMatakuliah(matakuliah *model.Matakuliah) error {
	return config.DB.Save(matakuliah).Error
}

func HapusMatakuliah(id uint) error {
	return config.DB.Delete(&model.Matakuliah{}, id).Error
}
