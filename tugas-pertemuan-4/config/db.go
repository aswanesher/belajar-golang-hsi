package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"worker-assignment-perhitungan-nilai-mahasiswa/models"
)

var DB *gorm.DB // Variabel global untuk menyimpan instance GORM

func createDSN() string {
	// Memuat variabel koneksi DB dari file .env
	if err := godotenv.Load(); err != nil {
		log.Println("Gagal memuat file .env, pastikan file tersebut ada:", err)
	}

	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbPort := os.Getenv("DB_PORT")
	if dbPort == "" {
		dbPort = "5432" // gunakan port default PostgreSQL jika tidak diatur
	}

	if dbHost == "" || dbUser == "" || dbName == "" || dbPass == "" {
		log.Fatal("Variabel lingkungan untuk koneksi database tidak lengkap. Pastikan DB_HOST, DB_USER, DB_NAME, dan DB_PASS sudah diatur.")
	}

	return fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%s sslmode=disable",
		dbHost, dbUser, dbName, dbPass, dbPort)
}

// ConnectDB menginisialisasi koneksi ke database PostgreSQL menggunakan GORM
func ConnectDB() {
	// panggil fungsi init
	dsn := createDSN()

	// Buka koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	fmt.Println("Berhasil terhubung ke database")

	DB = db // Simpan instance GORM ke variabel global untuk digunakan di tempat lain
}

// Fungsi untuk melakukan migrasi model ke database
func Migrate() {
	dsn := createDSN()

	// Buka koneksi ke database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database untuk migrasi:", err)
	}

	// Lakukan migrasi model ke database
	err = db.AutoMigrate(&models.Mahasiswa{}, &models.Tugas{}, &models.Hasil{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi:", err)
	}
	fmt.Println("Migrasi berhasil")
}

// Fungsi untuk seeding data awal ke database tabel mahasiswas
func SeedMahasiswa() {
	dsb := createDSN()

	db, err := gorm.Open(postgres.Open(dsb), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database untuk seeding:", err)
	}

	var count int64
	db.Model(&models.Mahasiswa{}).Count(&count)
	if count > 0 {
		fmt.Println("Seeding data mahasiswa tidak diperlukan, data sudah ada.")
		return
	}

	// Buat contoh data mahasiswa
	mahasiswa := []models.Mahasiswa{
		{Nama: "Budi Santoso"},
		{Nama: "Siti Aminah"},
		{Nama: "Joko Widodo"},
		{Nama: "Dewi Sartika"},
		{Nama: "Rina Marlina"},
	}

	// Lakukan seeding data mahasiswa
	for _, m := range mahasiswa {
		if err := db.Create(&m).Error; err != nil {
			log.Printf("Gagal menambahkan mahasiswa %s: %v\n", m.Nama, err)
		} else {
			fmt.Printf("Mahasiswa %s berhasil ditambahkan\n", m.Nama)
		}
	}

	// Tutup koneksi database
	if sqlDB, err := db.DB(); err != nil {
		log.Fatal("Gagal mendapatkan database instance untuk seeding:", err)
	} else {
		defer sqlDB.Close()
		fmt.Println("Koneksi database ditutup setelah seeding")
	}
}
