package models

import (
	"gorm.io/gorm"
)

type Hasil struct {
	gorm.Model
	TugasID uint `json:"tugas_id" gorm:"not null"`
	Nilai   int  `json:"nilai" gorm:"not null"`
}
