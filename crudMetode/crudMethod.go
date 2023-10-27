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

	for {
		fmt.Println("Pilih operasi yang ingin Anda lakukan:")
		fmt.Println("1. Tambah Metode Pembayaran")
		fmt.Println("2. Tampilkan Metode Pembayaran")
		fmt.Println("3. Update Metode Pembayaran")
		fmt.Println("4. Hapus Metode Pembayaran")
		fmt.Println("5. Keluar")
		fmt.Print("Pilihan: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addMetodePembayaran(db)
		case 2:
			showMetodePembayaran(db)
		case 3:
			updateMetodePembayaran(db)
		case 4:
			deleteMetodePembayaran(db)
		case 5:
			fmt.Println("Terima kasih! Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

func addMetodePembayaran(db *gorm.DB) {
	var metodePembayaran model.Metode_Pembayaran

	fmt.Print("Nama Metode Pembayaran: ")
	fmt.Scanln(&metodePembayaran.Method_name)

	// Tambahkan metode pembayaran ke database
	db.Create(&metodePembayaran)

	fmt.Println("Metode Pembayaran telah ditambahkan.")
}

func showMetodePembayaran(db *gorm.DB) {
	var metodePembayaran []model.Metode_Pembayaran
	db.Find(&metodePembayaran)

	fmt.Println("Daftar Metode Pembayaran:")
	for _, m := range metodePembayaran {
		fmt.Printf("ID: %d, Nama Metode Pembayaran: %s\n", m.ID, m.Method_name)
	}
}

func updateMetodePembayaran(db *gorm.DB) {
	var metodePembayaran model.Metode_Pembayaran

	fmt.Print("Masukkan ID Metode Pembayaran yang ingin diupdate: ")
	var idInput string
	fmt.Scanln(&idInput)

	id, err := strconv.ParseUint(idInput, 10, 64)
	if err != nil {
		fmt.Println("ID Metode Pembayaran tidak valid.")
		return
	}

	err = db.First(&metodePembayaran, id).Error
	if err != nil {
		fmt.Println("Metode Pembayaran tidak ditemukan.")
		return
	}

	fmt.Print("Nama Metode Pembayaran (Enter untuk tidak mengubah): ")
	var namaInput string
	fmt.Scanln(&namaInput)
	if strings.TrimSpace(namaInput) != "" {
		metodePembayaran.Method_name = namaInput
	}

	// Update metode pembayaran di database
	db.Save(&metodePembayaran)

	fmt.Println("Metode Pembayaran telah diupdate.")
}

func deleteMetodePembayaran(db *gorm.DB) {
	var metodePembayaran model.Metode_Pembayaran

	fmt.Print("Masukkan ID Metode Pembayaran yang ingin dihapus: ")
	var idInput string
	fmt.Scanln(&idInput)

	id, err := strconv.ParseUint(idInput, 10, 64)
	if err != nil {
		fmt.Println("ID Metode Pembayaran tidak valid.")
		return
	}

	err = db.First(&metodePembayaran, id).Error
	if err != nil {
		fmt.Println("Metode Pembayaran tidak ditemukan.")
		return
	}

	// Hapus metode pembayaran dari database
	db.Delete(&metodePembayaran)

	fmt.Println("Metode Pembayaran telah dihapus.")
}
