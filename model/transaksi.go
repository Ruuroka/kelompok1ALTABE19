package model

import "time"

type Transaksi struct {
	No_nota           uint `gorm:"primaryKey"`
	Tanggal_transaksi time.Time
	UserID            uint
	No_hp             uint
	Id_metode         uint
}
