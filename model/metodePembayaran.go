package model

type metodePembayaran struct {
	payment_method_id uint `gorm:"primaryKey"`
	method_name       string
}
