package model

import (
	"gorm.io/gorm"
)

type Metode_Pembayaran struct {
	gorm.Model
	Method_name string
}
