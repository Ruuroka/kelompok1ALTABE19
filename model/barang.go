package model

import "gorm.io/gorm"

type Barang struct {
	gorm.Model
	Nama_barang      string
	Desc_barang      string
	Harga_barang     string
	Stock            uint
	UserID           uint
	User             User
	TransaksiDetails []TransaksiDetail `gorm:"foreignKey:Id_barang"`
}
