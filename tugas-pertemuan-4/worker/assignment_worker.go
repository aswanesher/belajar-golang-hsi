package worker

import (
	"fmt"
	"sync"
	"worker-assignment-perhitungan-nilai-mahasiswa/config"
	"worker-assignment-perhitungan-nilai-mahasiswa/models"
)

// assign tugas ke mahasiswa
func AssignTugasToMahasiswa(m []models.Mahasiswa, ch chan models.Tugas, wg *sync.WaitGroup) {

	// Menggunakan WaitGroup untuk menunggu semua goroutine selesai
	defer wg.Done()

	// Membuat list tugas yang akan diassign
	listTugas := []string{
		"Tugas Pemrograman Goroutine",
		"Tugas Implementasi WaitGroup",
		"Tugas Implementasi Mutex",
		"Tugas Implementasi Channel",
		"Tugas Implementasi WaitGroup",
	}

	for i, m := range m {
		// Membuat tugas baru untuk setiap mahasiswa
		tugas := models.Tugas{
			Judul:       listTugas[i%len(listTugas)], // Menggunakan modulus untuk menghindari index out of range
			Deskripsi:   "Deskripsi tugas untuk " + listTugas[i%len(listTugas)],
			MahasiswaID: m.ID,
		}
		config.DB.Create(&tugas) // Simpan tugas ke database

		// Menampilkan pesan tugas diberikan
		fmt.Printf("Tugas %s diberikan ke %s\n", tugas.Judul, m.Nama)

		// Kirim tugas ke channel
		ch <- tugas
	}

	// Tutup channel setelah semua tugas diassign
	// close(ch)
}
