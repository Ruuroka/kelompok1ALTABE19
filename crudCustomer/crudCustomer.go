package main

import (
	"fmt"
	"kelompok1ALTABE19/config"
	"kelompok1ALTABE19/model"
	"strconv"

	"gorm.io/gorm"
)

func main() {
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Terjadi kesalahan:", err.Error())
		return
	}

	for {
		fmt.Println("Pilih operasi yang ingin Anda lakukan pada Customer:")
		fmt.Println("1. Tambah Customer")
		fmt.Println("2. Tampilkan Customer")
		fmt.Println("3. Update Customer")
		fmt.Println("4. Hapus Customer")
		fmt.Println("5. Keluar")
		fmt.Print("Pilihan: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addCustomer(db)
		case 2:
			showCustomers(db)
		case 3:
			updateCustomer(db)
		case 4:
			deleteCustomer(db)
		case 5:
			fmt.Println("Terima kasih! Keluar dari program.")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
		}
	}
}

func addCustomer(db *gorm.DB) {
	var customer model.Customer

	fmt.Print("Nama: ")
	fmt.Scanln(&customer.Nama)
	fmt.Print("Alamat: ")
	fmt.Scanln(&customer.Alamat)
	fmt.Print("Email: ")
	fmt.Scanln(&customer.Email)

	// Tambahkan customer ke database
	db.Create(&customer)

	fmt.Println("Customer telah ditambahkan.")
}

func showCustomers(db *gorm.DB) {
	var customers []model.Customer
	db.Find(&customers)

	fmt.Println("Daftar Customers:")
	for _, c := range customers {
		fmt.Printf("No_hp: %d, Nama: %s, Alamat: %s, Email: %s\n", c.No_hp, c.Nama, c.Alamat, c.Email)
	}
}

func updateCustomer(db *gorm.DB) {
	var customer model.Customer

	fmt.Print("Masukkan No_hp Customer yang ingin diupdate: ")
	var noHPInput string
	fmt.Scanln(&noHPInput)

	noHP, err := strconv.ParseUint(noHPInput, 10, 64)
	if err != nil {
		fmt.Println("No_hp Customer tidak valid.")
		return
	}

	err = db.First(&customer, noHP).Error
	if err != nil {
		fmt.Println("Customer tidak ditemukan.")
		return
	}

	fmt.Print("Nama (Enter untuk tidak mengubah): ")
	var namaInput string
	fmt.Scanln(&namaInput)
	if namaInput != "" {
		customer.Nama = namaInput
	}

	fmt.Print("Alamat (Enter untuk tidak mengubah): ")
	var alamatInput string
	fmt.Scanln(&alamatInput)
	if alamatInput != "" {
		customer.Alamat = alamatInput
	}

	fmt.Print("Email (Enter untuk tidak mengubah): ")
	var emailInput string
	fmt.Scanln(&emailInput)
	if emailInput != "" {
		customer.Email = emailInput
	}

	// Update customer di database
	db.Save(&customer)

	fmt.Println("Customer telah diupdate.")
}

func deleteCustomer(db *gorm.DB) {
	var customer model.Customer

	fmt.Print("Masukkan No_hp Customer yang ingin dihapus: ")
	var noHPInput string
	fmt.Scanln(&noHPInput)

	noHP, err := strconv.ParseUint(noHPInput, 10, 64)
	if err != nil {
		fmt.Println("No_hp Customer tidak valid.")
		return
	}

	err = db.First(&customer, noHP).Error
	if err != nil {
		fmt.Println("Customer tidak ditemukan.")
		return
	}

	// Hapus customer dari database
	db.Delete(&customer)

	fmt.Println("Customer telah dihapus.")
}
