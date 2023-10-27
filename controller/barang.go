package controller

import (
	"fmt"
	"kelompok1ALTABE19/model"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

type BarangSystem struct {
	DB *gorm.DB
}

func (ts *BarangSystem) addBarang(db *gorm.DB) (model.Barang, bool) {
	var newBarang = new(model.Barang)

	fmt.Print("Nama Barang: ")
	fmt.Scanln(&newBarang.Nama_barang)
	fmt.Print("Deskripsi Barang: ")
	fmt.Scanln(&newBarang.Desc_barang)
	fmt.Print("Harga Barang: ")
	fmt.Scanln(&newBarang.Harga_barang)
	fmt.Print("Stok Barang: ")
	fmt.Scanln(&newBarang.Stock)

	// Tambahkan barang ke database
	err := ts.DB.Create(newBarang).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Barang{}, false
	}

	return *newBarang, true
}

func showBarang(db *gorm.DB) {
	var barang []model.Barang
	db.Find(&barang)

	fmt.Println("Daftar Barang:")
	for _, b := range barang {
		fmt.Printf("ID: %d, Nama: %s, Harga: %s, Stok: %d\n", b.ID, b.Nama_barang, b.Harga_barang, b.Stock)
	}
}

func updateBarang(db *gorm.DB) {
	var barang model.Barang

	fmt.Print("Masukkan ID Barang yang ingin diupdate: ")
	var idInput string
	fmt.Scanln(&idInput)

	id, err := strconv.ParseUint(idInput, 10, 64)
	if err != nil {
		fmt.Println("ID Barang tidak valid.")
		return
	}

	err = db.First(&barang, id).Error
	if err != nil {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	fmt.Print("Nama Barang (Enter untuk tidak mengubah): ")
	var namaInput string
	fmt.Scanln(&namaInput)
	if strings.TrimSpace(namaInput) != "" {
		barang.Nama_barang = namaInput
	}

	fmt.Print("Deskripsi Barang (Enter untuk tidak mengubah): ")
	var descInput string
	fmt.Scanln(&descInput)
	if strings.TrimSpace(descInput) != "" {
		barang.Desc_barang = descInput
	}

	fmt.Print("Harga Barang (Enter untuk tidak mengubah): ")
	var hargaInput string
	fmt.Scanln(&hargaInput)
	if strings.TrimSpace(hargaInput) != "" {
		barang.Harga_barang = hargaInput
	}

	fmt.Print("Stok Barang (Enter untuk tidak mengubah): ")
	var stokInput string
	fmt.Scanln(&stokInput)
	if strings.TrimSpace(stokInput) != "" {
		stok, err := strconv.Atoi(stokInput)
		if err == nil {
			barang.Stock = stok
		}
	}

	// Update barang di database
	db.Save(&barang)

	fmt.Println("Barang telah diupdate.")
}

func deleteBarang(db *gorm.DB) {
	var barang model.Barang

	fmt.Print("Masukkan ID Barang yang ingin dihapus: ")
	var idInput string
	fmt.Scanln(&idInput)

	id, err := strconv.ParseUint(idInput, 10, 64)
	if err != nil {
		fmt.Println("ID Barang tidak valid.")
		return
	}

	err = db.First(&barang, id).Error
	if err != nil {
		fmt.Println("Barang tidak ditemukan.")
		return
	}

	// Hapus barang dari database
	db.Delete(&barang)

	fmt.Println("Barang telah dihapus.")
}
