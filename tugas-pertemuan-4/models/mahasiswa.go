package models

import (
	"gorm.io/gorm"
	"time"
)

// Mahasiswa struct untuk representasi data mahasiswa
type Mahasiswa struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Nama      string    `json:"nama" gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `json:"created_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:timestamp;default:CURRENT_TIMESTAMP"`
	DeletedAt time.Time `json:"deleted_at" gorm:"type:timestamp;default:NULL"`
}

// TableName mengembalikan nama tabel untuk model Mahasiswa
func (m *Mahasiswa) BeforeCreate(tx *gorm.DB) (err error) {
	// Mengatur ID otomatis jika belum diatur
	if m.ID == 0 {
		var lastID int64
		tx.Model(&Mahasiswa{}).Select("MAX(id)").Scan(&lastID)
		m.ID = int(lastID + 1)
		if m.ID < 1 {
			m.ID = 1 // Pastikan ID tidak kurang dari 1
		}
	}
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	return
}
