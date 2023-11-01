package controller

import (
	"bufio"
	"fmt"
	"kelompok1ALTABE19/model"
	"os"
	"strings"

	"gorm.io/gorm"
)

type BarangSystem struct {
	DB *gorm.DB
}

func (ts *BarangSystem) AddBarang(userID uint) (model.Barang, bool) {
	var newBarang = new(model.Barang)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama Barang: ")
	fmt.Scanln(&newBarang.Nama_barang)
	fmt.Print("Deskripsi Barang: ")
	newBarang.Desc_barang, _ = reader.ReadString('\n')
	newBarang.Desc_barang = strings.TrimSpace(newBarang.Desc_barang)
	fmt.Print("Harga Barang: ")
	fmt.Scanln(&newBarang.Harga_barang)
	fmt.Print("Stok Barang: ")
	fmt.Scanln(&newBarang.Stock)
	newBarang.UserID = userID

	err := ts.DB.Create(newBarang).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Barang{}, false
	}
	return *newBarang, true
}

func (ts *BarangSystem) ShowBarang(userID uint) ([]model.Barang, bool) {
	var barang []model.Barang
	// ts.DB.InnerJoins("INNER JOIN users ON barangs.user_id = users.id").Find(&barang)
	ts.DB.Find(&barang)

	if len(barang) == 0 {
		fmt.Println("Daftar barang kosong.")
		return nil, false
	}
	return barang, true
}

func (ts *BarangSystem) UpdateBarang(userID uint, barangID uint) (model.Barang, bool) {
	existingBarang := model.Barang{}
	err := ts.DB.First(&existingBarang, barangID).Error
	if err != nil {
		return model.Barang{}, false
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nama Barang: ")
	fmt.Scanln(&existingBarang.Nama_barang)
	fmt.Print("Deskripsi Barang: ")
	existingBarang.Desc_barang, _ = reader.ReadString('\n')
	existingBarang.Desc_barang = strings.TrimSpace(existingBarang.Desc_barang)
	fmt.Print("Harga Barang: ")
	fmt.Scanln(&existingBarang.Harga_barang)
	fmt.Print("Stok Barang: ")
	fmt.Scanln(&existingBarang.Stock)
	existingBarang.UserID = userID

	err = ts.DB.Save(&existingBarang).Error
	if err != nil {
		return model.Barang{}, false
	}

	return existingBarang, true
}

func (ts *BarangSystem) UpdateStock(userID uint, barangID uint) (model.Barang, bool) {
	existingBarang := model.Barang{}
	err := ts.DB.First(&existingBarang, barangID).Error
	if err != nil {
		return model.Barang{}, false
	}

	fmt.Print("Stok Barang: ")
	fmt.Scanln(&existingBarang.Stock)
	existingBarang.UserID = userID

	err = ts.DB.Save(&existingBarang).Error
	if err != nil {
		return model.Barang{}, false
	}

	return existingBarang, true
}

func (ts *BarangSystem) DeleteBarang(userID uint, barangID uint) bool {
	existingBarang := model.Barang{}
	err := ts.DB.Where("id = ? AND user_id = ?", barangID, userID).First(&existingBarang).Error
	if err != nil {
		return false
	}

	err = ts.DB.Delete(&existingBarang).Error
	if err != nil {
		return false
	}

	return true
}
