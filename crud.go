package main

import (
	"fmt"
	"kelompok1ALTABE19/config"
	"kelompok1ALTABE19/model"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err.Error())
		return
	}

	defer db.Close()

	for {
		fmt.Println("Pilih operasi yang ingin Anda lakukan:")
		fmt.Println("1. Tambah Barang")
		fmt.Println("2. Tampilkan Barang")
		fmt.Println("3. Update Barang")
		fmt.Println("4. Hapus Barang")
		fmt.Println("5. Keluar")
		fmt.Print("Pilihan: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addBarang(db)
		case 2:
			showBarang(db)
		case 3:
			updateBarang(db)
		case 4:
			deleteBarang(db)
		case 5:
			fmt.Println("Terima kasih! Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

func addBarang(db *gorm.DB) {
	var barang model.Barang

	fmt.Print("Nama Barang: ")
	fmt.Scanln(&barang.Nama_barang)
	fmt.Print("Deskripsi Barang: ")
	fmt.Scanln(&barang.Desc_barang)
	fmt.Print("Harga Barang: ")
	fmt.Scanln(&barang.Harga_barang)
	fmt.Print("Stok Barang: ")
	fmt.Scanln(&barang.Stock)

	// Tambahkan barang ke database
	db.Create(&barang)

	fmt.Println("Barang telah ditambahkan.")
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
