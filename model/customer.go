package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	Id        uint `gorm:"primaryKey"`
	No_hp     uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Nama      string
	Alamat    string
	Email     string
}
