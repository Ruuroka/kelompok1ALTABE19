package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nama         string
	Password     string
	status_users string
	barangs      []Barang
}
