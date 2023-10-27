package main

import (
	"fmt"
	"kelompok1ALTABE19/config"
	"kelompok1ALTABE19/model"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Something happend", err.Error())
		return
	}

	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Barang{})
	db.AutoMigrate(&model.Customer{})
	db.AutoMigrate(&model.Transaksi{})
}
