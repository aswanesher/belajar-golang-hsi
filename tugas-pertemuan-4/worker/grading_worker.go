package worker

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
	"worker-assignment-perhitungan-nilai-mahasiswa/config"
	"worker-assignment-perhitungan-nilai-mahasiswa/models"
)

func NilaiMahasiswa(ch chan models.Tugas, hasilChannel chan models.Hasil, wg *sync.WaitGroup) {
	// Menggunakan WaitGroup untuk menunggu semua goroutine selesai
	defer wg.Done()

	var wgDB sync.WaitGroup

	for tugas := range ch {
		wgDB.Add(1)
		go func(tugas models.Tugas) {
			defer wgDB.Done()
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			nilai := r.Intn(101) // Generate nilai acak antara 0-100
			hasil := models.Hasil{
				TugasID: tugas.ID,
				Nilai:   nilai,
			}
			// Simpan hasil ke database
			config.DB.Create(&hasil)

			// Tampilkan pesan hasil penilaian
			var mahasiswa models.Mahasiswa
			config.DB.First(&mahasiswa, tugas.MahasiswaID)
			fmt.Printf("Nilai untuk tugas '%s' mahasiswa '%s' adalah %d\n", tugas.Judul, mahasiswa.Nama, nilai)

		}(tugas)
	}

	wgDB.Wait()

}
