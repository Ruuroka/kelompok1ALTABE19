package main

import (
	"fmt"
	"kelompok1ALTABE19/auth"
	"kelompok1ALTABE19/config"
	"kelompok1ALTABE19/controller"
)

func main() {
	var inputMenu int
	db, err := config.InitDB()
	if err != nil {
		fmt.Println("Something happend", err.Error())
		return
	}
	// db.AutoMigrate(&model.User{})
	// db.AutoMigrate(&model.Barang{})
	// db.AutoMigrate(&model.Customer{})
	// db.AutoMigrate(&model.Transaksi{})
	// db.AutoMigrate(&model.TransaksiDetail{})
	// db.AutoMigrate(&model.Metode_Pembayaran{})

	var auth = auth.AuthSystem{DB: db}
	var barang = controller.BarangSystem{DB: db}

	for {
		fmt.Println("1. Login")
		fmt.Println("99. Exit")
		fmt.Print("Masukkan pilihan:")
		fmt.Scanln(&inputMenu)
		switch inputMenu {
		case 1:
			var menuLogin int
			result, permit := auth.Login()
			if permit {
				fmt.Println("Selamat datang ", result.Nama)
				if result.Status_users == "admin" {
					for permit {
						fmt.Println("======Tampilan Admin========")
						for permit {
							fmt.Println("1. Tambahkann akun Pegawai")
							fmt.Println("2. Menu Barang")
							fmt.Println("3. Update todo")
							fmt.Println("4. Delete todo")
							fmt.Println("0. Logout")
							fmt.Println("99. Exit")
							fmt.Print("Masukkan pilihan:")
							fmt.Scanln(&menuLogin)
							switch menuLogin {
							case 1:
								result, permit := auth.Register()
								if permit {
									fmt.Println(result, "Berhasil menambahkan data")
								}
							case 2:
								for permit {
									fmt.Println("======Menu Barang========")
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
										result, permit := barang.AddBarang(result.ID)
										if permit {
											fmt.Println("Barang berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("ID: %d\nNama: %s\nHarga: %s\nStok: %d\n", result.ID, result.Nama_barang, result.Harga_barang, result.Stock)
										}

									case 2:
										result, permit := barang.ShowBarang(result.ID)
										if permit {
											fmt.Println("===Daftar Barang===")
											for _, b := range result {
												fmt.Printf("ID: %d, Nama: %s, Harga: %s, Stok: %d, Nama Editor: %d\n", b.ID, b.Nama_barang, b.Harga_barang, b.Stock, b.UserID)
											}
										}
									case 3:
										var barangID uint
										fmt.Print("Masukkan ID barang yang akan diperbarui: ")
										fmt.Scanln(&barangID)
										result, permit := barang.UpdateBarang(result.ID, barangID)
										if permit {
											fmt.Println("Barang berhasil diperbarui dengan detail berikut:")
											fmt.Printf("ID: %d\nNama: %s\nHarga: %s\nStok: %d\n Nama Editor:%d", result.ID, result.Nama_barang, result.Harga_barang, result.Stock, result.UserID)
										}
									case 4:
										var barangID uint
										fmt.Print("Masukkan ID barang yang akan dihapus: ")
										fmt.Scanln(&barangID)

										permit := barang.DeleteBarang(result.ID, barangID)
										if permit {
											fmt.Println("Tugas berhasil dihapus")
										}
									case 5:
										return
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
							case 3:
							case 4:
							case 5:
							case 0:
								permit = false
							case 99:
								fmt.Println("Thank you....")
								return
							}
						}
					}
				} else if result.Status_users == "" {
					for permit {
						fmt.Println("======Tampilan Pegawai========")
						for permit {
							fmt.Println("1. Menu barang")
							fmt.Println("3. Update todo")
							fmt.Println("4. Delete todo")
							fmt.Println("0. Logout")
							fmt.Println("99. Exit")
							fmt.Print("Masukkan pilihan:")
							fmt.Scanln(&menuLogin)
							switch menuLogin {
							case 1:
								for {
									fmt.Println("======Menu Barang========")
									fmt.Println("1. Tambah Barang")
									fmt.Println("2. Tampilkan Barang")
									fmt.Println("3. Update Barang")
									fmt.Println("4. Keluar")
									fmt.Print("Pilihan: ")

									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 1:
										result, permit := barang.AddBarang(result.ID)
										if permit {
											fmt.Println("Barang berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("ID: %d\nNama: %s\nHarga: %s\nStok: %d\n", result.ID, result.Nama_barang, result.Harga_barang, result.Stock)
										}
									case 2:
										result, permit := barang.ShowBarang(result.ID)
										if permit {
											for _, b := range result {
												fmt.Printf("ID: %d, Nama: %s, Harga: %s, Stok: %d, Nama Editor:%s", b.ID, b.Nama_barang, b.Harga_barang, b.Stock, b.User.Nama)
											}
										}
									case 3:
										var barangID uint
										fmt.Print("Masukkan ID barang yang akan diperbarui: ")
										fmt.Scanln(&barangID)
										result, permit := barang.UpdateBarang(result.ID, barangID)
										if permit {
											fmt.Println("Barang berhasil diperbarui dengan detail berikut:")
											fmt.Printf("ID: %d\nNama: %s\nHarga: %s\nStok: %d\n Nama Editor:%d", result.ID, result.Nama_barang, result.Harga_barang, result.Stock, result.UserID)
										}
									case 4:
										fmt.Println("Terima kasih! Keluar dari program.")
										return
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
							case 2:
							case 3:
							case 4:
							case 0:
								permit = false
							case 99:
								fmt.Println("Thank you....")
								return
							}
						}
					}
				}
			}
		case 99:
			fmt.Println("Thank you....")
			return
		default:
		}
	}
}
