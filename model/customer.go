package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	No_hp      uint `gorm:"primaryKey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Nama       string
	Alamat     string
	Email      string
	Transaksis []Transaksi `gorm:"foreignKey:No_hp"`
}
