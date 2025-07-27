package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// Koneksi ke database PostgreSQL menggunakan GORM
func ConnectDB() {
	// Mengecek apakah file .env ada dan memuat variabel
	if err := godotenv.Load(); err != nil {
		log.Println("Gagal memuat file .env, pastikan file tersebut ada:", err)
	}

	// Validasi environment variables untuk koneksi database
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")

	if dbHost == "" || dbUser == "" || dbName == "" || dbPass == "" {
		log.Fatal("Variabel lingkungan untuk koneksi database tidak lengkap. Pastikan DB_HOST, DB_USER, DB_NAME, dan DB_PASS sudah diatur.")
	}

	// Membuat string koneksi untuk PostgreSQL
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=5432 sslmode=disable",
		dbHost, dbUser, dbName, dbPass)

	// Memulai koneksi ke database PostgreSQL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Gagal terhubung ke database:", err)
	}
	fmt.Println("Berhasil terhubung ke database")

	// Mendapatkan instance SQL dari GORM untuk operasi lebih lanjut
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Gagal mendapatkan database instance:", err)
	}

	// Mengatur pengaturan koneksi
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatal("Gagal menutup koneksi database:", err)
		}
		fmt.Println("Koneksi database ditutup")
	}()
}
