package validasi

import "testing"

type DataUjiNIM struct {
	Nim string `validate:"required,min=5,max=20,nim"`
}

func TestValidasiNIM(t *testing.T) {
	kasus := []struct {
		nama  string
		nim   string
		valid bool
	}{
		{"NIM H1H024068", "H1H024068", true},
		{"NIM H12345678", "H12345678", true},
		{"NIM diawali A", "A1H024068", false},
		{"Karakter kedua bukan angka", "HABC12345", false},
		{"NIM diawali angka", "12345678", false},
	}

	for _, uji := range kasus {
		t.Run(uji.nama, func(t *testing.T) {
			data := DataUjiNIM{Nim: uji.nim}
			hasil := Periksa(data)

			if (len(hasil) == 0) != uji.valid {
				t.Errorf(
					"NIM %q: diharapkan valid=%v, tetapi hasil valid=%v",
					uji.nim,
					uji.valid,
					len(hasil) == 0,
				)
			}
		})
	}
}
