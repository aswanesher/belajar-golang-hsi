package mahasiswa

import "fmt"

type Mahasiswa struct {
	Nama     string
	Nilai    []int
	umur     int     // private
	nilaiAvg float64 // private
}

type Deskripsi interface {
	Info() string
	RataRata() float64
	GetUmur() int
}

// Implementasi interface
func (m Mahasiswa) Info() string {
	return fmt.Sprintf("Nama: %s, Umur: %d", m.Nama, m.umur)
}

func (m Mahasiswa) RataRata() float64 {
	return m.nilaiAvg
}

func (m Mahasiswa) GetUmur() int {
	return m.umur
}

// Private variabel
var maxNilai int = 100

// Fungsi akses variabel private
func GetMaxNilai() int {
	return maxNilai
}

// Public variabel versi
var Versi string = "v1.0.0"
