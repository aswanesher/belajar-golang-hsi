package main

import (
	"os"
	"worker-assignment-perhitungan-nilai-mahasiswa/config"
)

func main() {
	// Test koneksi database
	config.ConnectDB()

	// jalankan perintah migrasi
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		config.Migrate()
		return
	}

	// jalankan perintah seeding
	if len(os.Args) > 1 && os.Args[1] == "seed" {
		config.SeedMahasiswa()
		return
	}
}
