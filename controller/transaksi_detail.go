// transaksi_detail_controller.go

package controller

import (
	"fmt"
	"kelompok1ALTABE19/model"

	"gorm.io/gorm"
)

type TransaksiDetailSystem struct {
	DB *gorm.DB
}

func (tds *TransaksiDetailSystem) AddTransaksiDetail(userNama string) (model.TransaksiDetail, bool) {
	var newTransaksiDetail = new(model.TransaksiDetail)
	var transaksiCustomer model.Customer
	var transaksiMetode model.Metode_Pembayaran
	fmt.Print("Masukkan Id Customer: ")
	fmt.Scanln(&transaksiCustomer.Id)
	fmt.Print("Masukkan Metode Pembayaran: ")
	fmt.Scanln(&transaksiMetode.ID)
	fmt.Print("Masukkan Nota Transaksi: ")
	fmt.Scanln(&newTransaksiDetail.Nota_transaksi)
	fmt.Print("Masukkan ID Barang: ")
	fmt.Scanln(&newTransaksiDetail.Id_barang)
	fmt.Print("Masukkan Jumlah Barang: ")
	fmt.Scanln(&newTransaksiDetail.Jumlah_barang)
	newTransaksiDetail.NamaUser = userNama

	if err := tds.DB.First(&transaksiCustomer, transaksiCustomer.Id).Error; err != nil {
		fmt.Println("Error mengambil data customer:", err)
		return model.TransaksiDetail{}, false
	}
	newTransaksiDetail.NamaCustomer = transaksiCustomer.Nama

	if err := tds.DB.First(&transaksiMetode, transaksiMetode.ID).Error; err != nil {
		fmt.Println("Error mengambil data customer:", err)
		return model.TransaksiDetail{}, false
	}
	newTransaksiDetail.NamaMetode = transaksiMetode.Method_name

	var transaksiBarangbaru model.Barang
	var transaksiBaru model.TransaksiDetail

	if err := tds.DB.First(&transaksiBarangbaru, newTransaksiDetail.Id_barang).Error; err != nil {
		fmt.Println("Error mengambil data barang:", err)
		return model.TransaksiDetail{}, false
	}

	if transaksiBarangbaru.Stock < transaksiBaru.Jumlah_barang {
		fmt.Println("Stock tidak mencukupi untuk jumlah barang yang diminta")
		return model.TransaksiDetail{}, false
	}

	transaksiBarangbaru.Stock -= transaksiBaru.Jumlah_barang

	if err := tds.DB.Save(&transaksiBarangbaru).Error; err != nil {
		fmt.Println("Error mengupdate stock barang:", err)
		return model.TransaksiDetail{}, false
	}

	// var hargaBarang float64
	err := tds.DB.Model(&transaksiBarangbaru).Where("id = ?", transaksiBarangbaru.ID).Pluck("harga_barang", &transaksiBarangbaru.Harga_barang).Error
	if err != nil {
		fmt.Println("Error mengambil harga barang:", err)
		return model.TransaksiDetail{}, false
	}
	transaksiBaru.Total_harga += transaksiBarangbaru.Harga_barang * transaksiBaru.Jumlah_barang

	// newTransaksiDetail.NamaCustomer = newTransaksiDetail.NamaCustomer

	err = tds.DB.Create(&newTransaksiDetail).Error
	if err != nil {
		fmt.Println("Input error:", err.Error())
		return model.TransaksiDetail{}, false
	}

	return *newTransaksiDetail, true
}

func (tds *TransaksiDetailSystem) ShowTransaksiDetail(notaTransaksi uint) ([]model.TransaksiDetail, bool) {
	var transaksiDetails []model.TransaksiDetail
	tds.DB.Where("nota_transaksi = ?", notaTransaksi).Find(&transaksiDetails)

	if len(transaksiDetails) == 0 {
		fmt.Println("Daftar transaksi detail kosong.")
		return nil, false
	}

	return transaksiDetails, true
}

// func (tds *TransaksiDetailSystem) UpdateTransaksiDetail(notaTransaksi uint, idBarang uint, jumlahBarang uint, totalHarga float64, statusPembayaran string) (model.TransaksiDetail, bool) {
// 	existingTransaksiDetail := model.TransaksiDetail{}
// 	err := tds.DB.Where("nota_transaksi = ? AND id_barang = ?", notaTransaksi, idBarang).First(&existingTransaksiDetail).Error
// 	if err != nil {
// 		return model.TransaksiDetail{}, false
// 	}

// 	var hargaBarang int
// 	err = tds.DB.Model(&model.Barang{}).Where("id = ?", idBarang).Pluck("harga_barang", &hargaBarang).Error
// 	if err != nil {
// 		fmt.Println("Error mengambil harga barang:", err)
// 		return model.TransaksiDetail{}, false
// 	}
// 	totalHarga = hargaBarang * jumlahBarang

// 	var barang model.Barang
// 	if err := tds.DB.First(&barang, idBarang).Error; err != nil {
// 		fmt.Println("Error mengambil data barang:", err)
// 		return model.TransaksiDetail{}, false
// 	}

// 	if barang.Stock < jumlahBarang {
// 		fmt.Println("Stock tidak mencukupi untuk jumlah barang yang diminta")
// 		return model.TransaksiDetail{}, false
// 	}

// 	barang.Stock -= jumlahBarang
// 	// Disesuaikan dengan nama field yang benar di model Barang
// 	barang.Harga_barang = totalHarga // Update harga_barang sesuai kebutuhan

// 	err = tds.DB.Save(&barang).Error
// 	if err != nil {
// 		fmt.Println("Error mengupdate data barang:", err)
// 		return model.TransaksiDetail{}, false
// 	}

// 	existingTransaksiDetail.Jumlah_barang = jumlahBarang
// 	existingTransaksiDetail.Total_harga = totalHarga
// 	existingTransaksiDetail.Status_pembayaran = statusPembayaran

// 	err = tds.DB.Save(&existingTransaksiDetail).Error
// 	if err != nil {
// 		return model.TransaksiDetail{}, false
// 	}

// 	return existingTransaksiDetail, true
// }

func (tds *TransaksiDetailSystem) DeleteTransaksiDetail(notaTransaksi uint, idBarang uint) bool {
	existingTransaksiDetail := model.TransaksiDetail{}
	err := tds.DB.Where("nota_transaksi = ? AND id_barang = ?", notaTransaksi, idBarang).First(&existingTransaksiDetail).Error
	if err != nil {
		return false
	}

	err = tds.DB.Delete(&existingTransaksiDetail).Error
	if err != nil {
		return false
	}

	return true
}
