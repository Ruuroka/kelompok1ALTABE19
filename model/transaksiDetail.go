package model

type TransaksiDetail struct {
	Nota_transaksi    uint `gorm:"primaryKey"`
	Id_barang         uint
	Jumlah_barang     uint
	Total_harga       float64
	Status_pembayaran string
}
