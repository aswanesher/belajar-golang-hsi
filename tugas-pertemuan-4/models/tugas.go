package models

import (
	"gorm.io/gorm"
	"time"
)

// Mahasiswa struct untuk representasi data mahasiswa
type Tugas struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Judul       string    `json:"judul" gorm:"type:varchar(100);not null"`
	Deskripsi   string    `json:"deskripsi" gorm:"type:text;null"`
	MahasiswaID int       `json:"mahasiswa_id" gorm:"not null"`
	Mahasiswa   Mahasiswa `json:"mahasiswa" gorm:"foreignKey:MahasiswaID;references:ID"`
	CreatedAt   time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"type:timestamp;default:NULL"`
}

// TableName mengembalikan nama tabel untuk model Tugas
func (t *Tugas) BeforeCreate(tx *gorm.DB) (err error) {
	// Mengatur ID otomatis jika belum diatur
	if t.ID == 0 {
		var lastID int64
		// Asumsikan ada fungsi untuk mendapatkan ID terakhir dari tabel Tugas
		tx.Model(&Tugas{}).Select("MAX(id)").Scan(&lastID)
		t.ID = int(lastID + 1)
		if t.ID < 1 {
			t.ID = 1 // Pastikan ID tidak kurang dari 1
		}
	}
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()
	return
}
