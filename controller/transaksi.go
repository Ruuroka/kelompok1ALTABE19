// transaksi_controller.go

package controller

import (
	"fmt"
	"kelompok1ALTABE19/model"
	"time"

	"gorm.io/gorm"
)

type TransaksiSystem struct {
	DB *gorm.DB
}

func (ts *TransaksiSystem) AddTransaksi(userID uint, no_Hp uint, metodeID uint, newTanggal time.Time) (model.Transaksi, bool) {
	var newTransaksi = model.Transaksi{
		UserID:    userID,
		No_hp:     no_Hp,
		Id_metode: metodeID,
	}

	err := ts.DB.Create(&newTransaksi).Error
	if err != nil {
		fmt.Println("Input error:", err.Error())
		return model.Transaksi{}, false
	}

	return newTransaksi, true
}

func (ts *TransaksiSystem) ShowTransaksi(userID uint) ([]model.Transaksi, bool) {
	var transaksi []model.Transaksi
	ts.DB.Joins("User").Find(&transaksi)
	// ts.DB.Joins("Customer").Find(&transaksi)
	// ts.DB.Joins("Metode_Pembayaran").Find(&transaksi)

	if len(transaksi) == 0 {
		fmt.Println("Daftar transaksi kosong.")
		return nil, false
	}
	return transaksi, true
}

func (ts *TransaksiSystem) UpdateTransaksi(userID uint, no_Nota uint, no_Hp uint, metodeID uint, newTanggal time.Time) (model.Transaksi, bool) {
	existingTransaksi := model.Transaksi{}
	err := ts.DB.Where("user_id = ? AND no_nota = ?", userID, no_Nota).First(&existingTransaksi).Error
	if err != nil {
		return model.Transaksi{}, false
	}
	existingTransaksi.UserID = userID
	existingTransaksi.No_hp = no_Hp
	existingTransaksi.Id_metode = metodeID

	// existingTransaksi.No_hp = no_Hp
	// existingTransaksi.Id_metode = metodeID

	err = ts.DB.Save(&existingTransaksi).Error
	if err != nil {
		return model.Transaksi{}, false
	}

	return existingTransaksi, true
}

func (ts *TransaksiSystem) DeleteTransaksi(userID uint, no_Nota uint) bool {
	existingTransaksi := model.Transaksi{}
	err := ts.DB.Where("user_id = ? AND no_nota = ?", userID, no_Nota).First(&existingTransaksi).Error
	if err != nil {
		return false
	}

	err = ts.DB.Delete(&existingTransaksi).Error
	if err != nil {
		return false
	}

	return true
}
