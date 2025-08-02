package models

import (
	"gorm.io/gorm"
)

// Tugas struct untuk representasi data tugas
type Tugas struct {
	gorm.Model
	Judul       string    `json:"judul" gorm:"type:varchar(100);not null"`
	Deskripsi   string    `json:"deskripsi" gorm:"type:text;null"`
	MahasiswaID uint      `json:"mahasiswa_id" gorm:"not null"`
	Mahasiswa   Mahasiswa `json:"mahasiswa" gorm:"foreignKey:MahasiswaID"`
	Hasil       Hasil     `json:"hasil" gorm:"foreignKey:TugasID"` // Relasi ke model Hasil (HasOne)
}
