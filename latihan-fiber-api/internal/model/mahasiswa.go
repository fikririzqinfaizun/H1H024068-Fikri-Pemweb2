package model

type Mahasiswa struct {
	ID       int     `json:"id"`
	NIM      string  `json:"nim"`
	Nama     string  `json:"nama"`
	Prodi    string  `json:"prodi"`
	Angkatan int     `json:"angkatan"`
	IPK      float64 `json:"ipk"`
}

var Daftar = []Mahasiswa{
	{
		ID:       1,
		NIM:      "H1H024001",
		Nama:     "Andi Pratama",
		Prodi:    "Teknik Komputer",
		Angkatan: 2024,
		IPK:      3.65,
	},
	{
		ID:       2,
		NIM:      "H1H023002",
		Nama:     "Bunga Lestari",
		Prodi:    "Teknik Komputer",
		Angkatan: 2023,
		IPK:      3.80,
	},
	{
		ID:       3,
		NIM:      "H1H023003",
		Nama:     "Citra Permata",
		Prodi:    "Teknik Informatika",
		Angkatan: 2023,
		IPK:      3.45,
	},
}
