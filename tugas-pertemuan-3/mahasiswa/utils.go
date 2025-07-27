package mahasiswa

// Variadic function
func hitungRataRata(nilai ...int) float64 {
	total := 0
	for _, n := range nilai {
		total += n
	}
	return float64(total) / float64(len(nilai))
}

// Gunakan pointer agar dapat mengubah umur
func BuatMahasiswa(nama string, umur int, nilai ...int) Mahasiswa {
	rata := hitungRataRata(nilai...)
	m := Mahasiswa{
		Nama:     nama,
		Nilai:    nilai,
		umur:     umur,
		nilaiAvg: rata,
	}
	return m
}
