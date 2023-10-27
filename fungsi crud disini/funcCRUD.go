package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Model Barang
type Product struct {
	ID          uint
	NamaBarang  string
	DescBarang  string
	HargaBarang float64
	Stock       int
	IDUsers     uint
}

func main() {
	// Ganti sesuai dengan detail koneksi database Anda.
	dsn := "user=username dbname=mydb sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	// AutoMigrate untuk membuat tabel barang
	db.AutoMigrate(&Product{})

	// Create (Membuat Barang Baru)
	newProduct := Product{
		NamaBarang:  "Contoh Barang",
		DescBarang:  "Deskripsi Contoh Barang",
		HargaBarang: 100.00,
		Stock:       50,
		IDUsers:     1,
	}
	db.Create(&newProduct)

	// Read (Membaca Barang)
	var readProduct Product
	db.First(&readProduct, 1) // Mengambil barang dengan ID 1
	fmt.Println(readProduct)

	// Update (Memperbarui Barang)
	db.Model(&readProduct).Updates(Product{
		NamaBarang: "Barang Diperbarui",
		Stock:      60,
	})

	// Delete (Menghapus Barang)
	db.Delete(&readProduct)

	// Anda juga dapat menambahkan fitur tambahan seperti mengambil daftar barang, dll.
}
