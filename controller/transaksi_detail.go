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

func (tds *TransaksiDetailSystem) AddTransaksiDetail(notaTransaksi uint, idBarang uint, jumlahBarang uint, totalHarga float64, statusPembayaran string) (model.TransaksiDetail, bool) {
	var hargaBarang float64
	err := tds.DB.Model(&model.Barang{}).Where("id = ?", idBarang).Pluck("harga_barang", &hargaBarang).Error
	if err != nil {
		fmt.Println("Error mengambil harga barang:", err)
		return model.TransaksiDetail{}, false
	}
	totalHarga = hargaBarang * float64(jumlahBarang)

	var barang model.Barang
	if err := tds.DB.First(&barang, idBarang).Error; err != nil {
		fmt.Println("Error mengambil data barang:", err)
		return model.TransaksiDetail{}, false
	}

	if barang.Stock < jumlahBarang {
		fmt.Println("Stock tidak mencukupi untuk jumlah barang yang diminta")
		return model.TransaksiDetail{}, false
	}

	barang.Stock -= jumlahBarang

	if err := tds.DB.Save(&barang).Error; err != nil {
		fmt.Println("Error mengupdate stock barang:", err)
		return model.TransaksiDetail{}, false
	}

	var newTransaksiDetail = model.TransaksiDetail{
		Nota_transaksi:    notaTransaksi,
		Id_barang:         idBarang,
		Jumlah_barang:     jumlahBarang,
		Total_harga:       totalHarga,
		Status_pembayaran: statusPembayaran,
	}

	err = tds.DB.Create(&newTransaksiDetail).Error
	if err != nil {
		fmt.Println("Input error:", err.Error())
		return model.TransaksiDetail{}, false
	}

	return newTransaksiDetail, true
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

func (tds *TransaksiDetailSystem) UpdateTransaksiDetail(notaTransaksi uint, idBarang uint, jumlahBarang uint, totalHarga float64, statusPembayaran string) (model.TransaksiDetail, bool) {
	existingTransaksiDetail := model.TransaksiDetail{}
	err := tds.DB.Where("nota_transaksi = ? AND id_barang = ?", notaTransaksi, idBarang).First(&existingTransaksiDetail).Error
	if err != nil {
		return model.TransaksiDetail{}, false
	}

	existingTransaksiDetail.Jumlah_barang = jumlahBarang
	existingTransaksiDetail.Total_harga = totalHarga
	existingTransaksiDetail.Status_pembayaran = statusPembayaran

	err = tds.DB.Save(&existingTransaksiDetail).Error
	if err != nil {
		return model.TransaksiDetail{}, false
	}

	return existingTransaksiDetail, true
}

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
