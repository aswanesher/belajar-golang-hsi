package main

import (
	"fmt"
	"tugas-pertemuan-3/mahasiswa"
)

func main() {
	// Membuat mahasiswa
	m1 := mahasiswa.BuatMahasiswa("Ali", 20, 90, 85, 80)
	m2 := mahasiswa.BuatMahasiswa("Siti", 22, 88, 76, 81)

	// Array untuk closure
	mahasiswas := []mahasiswa.Deskripsi{m1, m2}

	// Tampilkan data
	for _, m := range mahasiswas {
		fmt.Println(m.Info())
		fmt.Printf("Rata-rata nilai: %.2f\n", m.RataRata())
		fmt.Println("---")
	}

	// Versi dan nilai maksimum
	fmt.Println("Versi Package:", mahasiswa.Versi)
	fmt.Println("Nilai Maksimum:", mahasiswa.GetMaxNilai())

	// Closure untuk menghitung total umur
	totalUmur := func(data []mahasiswa.Deskripsi) int {
		total := 0
		for _, m := range data {
			total += m.GetUmur()
		}
		return total
	}

	fmt.Println("Total Umur Mahasiswa:", totalUmur(mahasiswas))
}
