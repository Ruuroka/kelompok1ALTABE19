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
	var metode = controller.MetodeSystem{DB: db}
	var customer = controller.CustomerSystem{DB: db}

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
							fmt.Println("3. Menu Metode Transaksi")
							fmt.Println("4. Menu Customer")
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
											for _, b := range result {
												fmt.Println("===Daftar Barang===")
												fmt.Printf("ID: %d\nNama: %s\nHarga: %s\nStok: %d\nNama Editor: %d \n", b.ID, b.Nama_barang, b.Harga_barang, b.Stock, b.UserID)
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
								for permit {
									fmt.Println("======Menu Metode Transaksi========")
									fmt.Println("1. Tambah Metode Transaksi")
									fmt.Println("2. Tampilkan Metode Transaksi")
									fmt.Println("3. Update Metode Transaksi")
									fmt.Println("4. Hapus Metode Transaksi")
									fmt.Println("5. Keluar")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 1:
										result, permit := metode.AddMetode()
										if permit {
											fmt.Println("Metode Pembayaran berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("Nama Metode: %s\n", result.Method_name)
										}

									case 2:
										result, permit := metode.ShowMetode()
										if permit {
											for _, m := range result {
												fmt.Println("===Daftar Metode Transaksi===")
												fmt.Printf("ID: %d\nNama Metode: %s\n", m.ID, m.Method_name)
											}
										}
									case 3:
										var metodeID uint
										fmt.Print("Masukkan ID metode yang akan diperbarui: ")
										fmt.Scanln(&metodeID)
										result, permit := metode.UpdateMetode(metodeID)
										if permit {
											fmt.Println("Metode Pembayaran berhasil diperbarui dengan detail berikut:")
											fmt.Printf("ID: %d\nNama Metode: %s\n", result.ID, result.Method_name)
										}
									case 4:
										var metodeID uint
										fmt.Print("Masukkan ID metode yang akan dihapus: ")
										fmt.Scanln(&metodeID)

										permit := metode.DeleteMetode(metodeID)
										if permit {
											fmt.Println("Metode Pembayaran berhasil dihapus")
										}
									case 5:
										return
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
							case 4:
								for permit {
									fmt.Println("======Menu Customer========")
									fmt.Println("1. Tambah Data Customer")
									fmt.Println("2. Tampilkan Daftar Customer")
									fmt.Println("3. Update Data Customer")
									fmt.Println("4. Hapus Data Customer")
									fmt.Println("5. Keluar")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 1:
										result, permit := customer.AddCustomer()
										if permit {
											fmt.Println("Data Customer berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("Nama Customer: %s\n", result.Nama)
										}

									case 2:
										result, permit := customer.ShowCustomer()
										if permit {
											for _, m := range result {
												fmt.Println("===Daftar Customer===")
												fmt.Printf("No HP: %d\nNama : %s\nAlamat: %s\nEmail: %s\n", m.No_hp, m.Nama, m.Alamat, m.Email)
											}
										}
									case 3:
										var no_Hp uint
										fmt.Print("Masukkan No Hp Customer yang akan diperbarui: ")
										fmt.Scanln(&no_Hp)
										result, permit := customer.UpdateCustomer(no_Hp)
										if permit {
											fmt.Println("Data Customer berhasil diperbarui dengan detail berikut:")
											fmt.Printf("No HP: %d\nNama : %s\nAlamat: %s\nEmail: %s\n", result.No_hp, result.Nama, result.Alamat, result.Email)
										}
									case 4:
										var no_Hp uint
										fmt.Print("Masukkan No Hp Customer yang akan dihapus: ")
										fmt.Scanln(&no_Hp)

										permit := customer.DeleteCustomer(no_Hp)
										if permit {
											fmt.Println("Data Customer berhasil dihapus")
										}
									case 5:
										return
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
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
							fmt.Println("1. Menu Barang")
							fmt.Println("2. Menu Metode Transaksi")
							fmt.Println("3. Menu Customer")
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
												fmt.Println("===Daftar Barang===")
												fmt.Printf("ID: %d\nNama: %s\nHarga: %s\nStok: %d\nNama Editor: %d \n", b.ID, b.Nama_barang, b.Harga_barang, b.Stock, b.UserID)
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
								for permit {
									fmt.Println("======Menu Metode Transaksi========")
									fmt.Println("1. Tambah Metode Transaksi")
									fmt.Println("2. Tampilkan Metode Transaksi")
									fmt.Println("3. Update Metode Transaksi")
									fmt.Println("4. Hapus Metode Transaksi")
									fmt.Println("5. Keluar")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 1:
										result, permit := metode.AddMetode()
										if permit {
											fmt.Println("Metode Pembayaran berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("Nama Metode: %s\n", result.Method_name)
										}

									case 2:
										result, permit := metode.ShowMetode()
										if permit {
											for _, m := range result {
												fmt.Println("===Daftar Metode Transaksi===")
												fmt.Printf("ID: %d\nNama Metode: %s\n", m.ID, m.Method_name)
											}
										}
									case 3:
										var metodeID uint
										fmt.Print("Masukkan ID metode yang akan diperbarui: ")
										fmt.Scanln(&metodeID)
										result, permit := metode.UpdateMetode(metodeID)
										if permit {
											fmt.Println("Metode Pembayaran berhasil diperbarui dengan detail berikut:")
											fmt.Printf("ID: %d\nNama Metode: %s\n", result.ID, result.Method_name)
										}
									case 4:
										var metodeID uint
										fmt.Print("Masukkan ID metode yang akan dihapus: ")
										fmt.Scanln(&metodeID)

										permit := metode.DeleteMetode(metodeID)
										if permit {
											fmt.Println("Metode Pembayaran berhasil dihapus")
										}
									case 5:
										return
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
							case 3:
								for permit {
									fmt.Println("======Menu Customer========")
									fmt.Println("1. Tambah Data Customer")
									fmt.Println("2. Tampilkan Daftar Customer")
									fmt.Println("3. Update Data Customer")
									fmt.Println("4. Hapus Data Customer")
									fmt.Println("5. Keluar")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 1:
										result, permit := customer.AddCustomer()
										if permit {
											fmt.Println("Barang berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("Nama Customer: %s\n", result.Nama)
										}

									case 2:
										result, permit := customer.ShowCustomer()
										if permit {
											for _, m := range result {
												fmt.Println("===Daftar Customer===")
												fmt.Printf("No HP: %d\nNama : %s\nAlamat: %s\nEmail: %s\n", m.No_hp, m.Nama, m.Alamat, m.Email)
											}
										}
									case 3:
										var no_Hp uint
										fmt.Print("Masukkan No HP Customer yang akan diperbarui: ")
										fmt.Scanln(&no_Hp)
										result, permit := customer.UpdateCustomer(no_Hp)
										if permit {
											fmt.Println("Data Customer berhasil diperbarui dengan detail berikut:")
											fmt.Printf("No HP: %d\nNama : %s\nAlamat: %s\nEmail: %s\n", result.No_hp, result.Nama, result.Alamat, result.Email)
										}
									case 4:
										var no_Hp uint
										fmt.Print("Masukkan No HP Customer yang akan dihapus: ")
										fmt.Scanln(&no_Hp)

										permit := customer.DeleteCustomer(no_Hp)
										if permit {
											fmt.Println("Data Customer berhasil dihapus")
										}
									case 5:
										return
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
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
