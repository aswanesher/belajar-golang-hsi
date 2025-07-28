package models

import (
	"gorm.io/gorm"
	"time"
)

type Hasil struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	TugasID   int       `json:"tugas_id" gorm:"not null"`
	Tugas     Tugas     `json:"tugas" gorm:"foreignKey:TugasID;references:ID"`
	Nilai     int       `json:"nilai" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"not null"`
	UpdatedAt time.Time `json:"updated_at" gorm:"not null"`
	DeletedAt time.Time `json:"deleted_at" gorm:"default:NULL"`
}

// BeforeCreate mengatur nilai default sebelum membuat entri baru
func (h *Hasil) BeforeCreate(tx *gorm.DB) (err error) {
	if h.ID == 0 {
		var lastID int64
		tx.Model(&Hasil{}).Select("MAX(id)").Scan(&lastID)
		h.ID = int(lastID + 1)
		if h.ID < 1 {
			h.ID = 1 // Pastikan ID tidak kurang dari 1
		}
	}
	h.CreatedAt = time.Now()
	h.UpdatedAt = time.Now()
	return
}
