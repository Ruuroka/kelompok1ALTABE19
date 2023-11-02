package model

import (
	"time"

	"gorm.io/gorm"
)

type TransaksiDetail struct {
	Nota_transaksi uint `gorm:"primaryKey"`
	NamaCustomer   string
	NamaUser       string
	NamaMetode     string
	Id_barang      uint
	Jumlah_barang  uint
	Total_harga    uint
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}
