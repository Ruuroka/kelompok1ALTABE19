package model

import "gorm.io/gorm"

type Barang struct {
	gorm.Model
	Nama_barang  string
	Desc_barang  string
	harga_barang string
	stock        int
	UserID       uint
}
