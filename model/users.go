package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama         string
	Password     string
	Status_users string
	Barangs      []Barang    `gorm:"foreignKey:UserID"`
	Transaksis   []Transaksi `gorm:"foreignKey:UserID"`
}
