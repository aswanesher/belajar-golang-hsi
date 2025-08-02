package main

import (
	"fmt"
	"sync"
	"worker-assignment-perhitungan-nilai-mahasiswa/config"
	"worker-assignment-perhitungan-nilai-mahasiswa/models"
	"worker-assignment-perhitungan-nilai-mahasiswa/worker"
)

func main() {
	config.ConnectDB()
	// Hapus data lama dari tabel hasil dan tugas
	config.DB.Exec("DELETE FROM hasils")
	config.DB.Exec("DELETE FROM tugas")
	// jalankan perintah migrasi
	// config.DB.AutoMigrate(&models.Mahasiswa{}, &models.Tugas{}, &models.Hasil{})

	var count int64
	config.DB.Model(&models.Mahasiswa{}).Count(&count)
	if count == 0 {
		config.SeedMahasiswa()
	}

	// Ambil semua mahasiswa dari database
	var mahasiswas []models.Mahasiswa
	if err := config.DB.Find(&mahasiswas).Error; err != nil {
		fmt.Println("Error fetching mahasiswa", err)
		return
	}

	// Assign tugas ke mahasiswa
	tugasChannel := make(chan models.Tugas)
	hasilChannel := make(chan models.Hasil)
	var wgAssign sync.WaitGroup
	var wgNilai sync.WaitGroup
	wgAssign.Add(1)
	wgNilai.Add(1)

	// Jalankan goroutine untuk assign tugas dan nilai mahasiswa
	go worker.AssignTugasToMahasiswa(mahasiswas, tugasChannel, &wgAssign)
	go worker.NilaiMahasiswa(tugasChannel, hasilChannel, &wgNilai)

	// Tunggu sampai semua tugas diassign, lalu tutup channel
	wgAssign.Wait()
	close(tugasChannel)

	// Tunggu sampai semua penilaian selesai
	go func() {
		wgNilai.Wait()
		close(hasilChannel)
	}()

	// Simpan hasil ke database
	for hasil := range hasilChannel {
		config.DB.Create(&hasil)
	}

	// Cetak hasil akhir
	fmt.Println("\nHasil Tugas Mahasiswa:")

	// Ambil semua mahasiswa dan tugas mereka
	for _, m := range mahasiswas {
		var tugas models.Tugas                                  // untuk menyimpan tugas yang diassign
		var hasil models.Hasil                                  // untuk menyimpan hasil penilaian
		config.DB.Where("mahasiswa_id = ?", m.ID).First(&tugas) // Ambil tugas pertama yang diassign ke mahasiswa
		config.DB.Where("tugas_id = ?", tugas.ID).First(&hasil) // Ambil hasil penilaian untuk tugas tersebut
		fmt.Printf("Mahasiswa: %s, Tugas: %s, Nilai: %d\n", m.Nama, tugas.Judul, hasil.Nilai)
	}
}
