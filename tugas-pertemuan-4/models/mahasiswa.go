package models

import (
	"gorm.io/gorm"
)

// Mahasiswa struct untuk representasi data mahasiswa
type Mahasiswa struct {
	gorm.Model
	Nama  string  `json:"nama" gorm:"type:varchar(100);not null"`
	Tugas []Tugas `json:"tugas" gorm:"foreignKey:MahasiswaID"` // Relasi ke model Tugas (HasMany)
}
