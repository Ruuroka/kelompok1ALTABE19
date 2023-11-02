package main

import (
	"fmt"
	"kelompok1ALTABE19/auth"
	"kelompok1ALTABE19/config"
	"kelompok1ALTABE19/controller"
	"time"
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
	var transaksi = controller.TransaksiSystem{DB: db}
	var transaksi_detail = controller.TransaksiDetailSystem{DB: db}

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
							fmt.Println("5. Menu Transaksi")
							fmt.Println("6. Detail Transaksi")
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
									fmt.Println("5. Update Stock")
									fmt.Println("0. Kembali")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 0:
										permit = false
									case 1:
										result, permit := barang.AddBarang(result.ID)
										if permit {
											fmt.Println("Barang berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("ID: %d\nNama: %s\nHarga: %2.f\nStok: %d\n", result.ID, result.Nama_barang, result.Harga_barang, result.Stock)
										}
									case 2:
										result, permit := barang.ShowBarang(result.ID)
										if permit {
											for _, b := range result {
												fmt.Println("===Daftar Barang===")
												fmt.Printf("ID: %d\nNama: %s\nHarga: %2.f\nStok: %d\nNama Editor: %d \n", b.ID, b.Nama_barang, b.Harga_barang, b.Stock, b.UserID)
											}
										}
									case 3:
										var barangID uint
										fmt.Print("Masukkan ID barang yang akan diperbarui: ")
										fmt.Scanln(&barangID)
										result, permit := barang.UpdateBarang(result.ID, barangID)
										if permit {
											fmt.Println("Barang berhasil diperbarui dengan detail berikut:")
											fmt.Printf("ID: %d\nNama: %s\nHarga: %2.f\nStok: %d\n Nama Editor:%d", result.ID, result.Nama_barang, result.Harga_barang, result.Stock, result.UserID)
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
										var barangID uint
										fmt.Print("Masukkan ID barang yang akan diperbarui: ")
										fmt.Scanln(&barangID)
										result, permit := barang.UpdateStock(result.ID, barangID)
										if permit {
											fmt.Println("Stock Barang berhasil diperbarui dengan detail berikut:")
											fmt.Printf("ID: %d\nNama: %s\nStok: %d\n Nama Editor:%d", result.ID, result.Nama_barang, result.Stock, result.UserID)
										}
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
									fmt.Println("0. kembali")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 0:
										permit = false
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
									fmt.Println("0. Kembali")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 0:
										permit = false
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
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
							case 5:
								for permit {
									fmt.Println("======Menu Transaksi========")
									fmt.Println("1. Tambah Data Transaksi")
									fmt.Println("2. Tampilkan Daftar Transaksi")
									fmt.Println("3. Update Data Transaksi")
									fmt.Println("4. Hapus Data Transaksi")
									fmt.Println("0. Kembali")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 0:
										permit = false
									case 1:
										var no_Hp uint
										var metodeID uint
										var tanggalTransaksi string
										fmt.Print("Masukkan No HP Customer: ")
										fmt.Scanln(&no_Hp)
										fmt.Print("Masukkan ID Metode Pembayaran: ")
										fmt.Scanln(&metodeID)
										fmt.Print("Masukkan Tanggal Transaksi baru: ")
										fmt.Scanln(&tanggalTransaksi)
										var dateFormat = "2006-01-02"
										newTanggal, err := time.Parse(dateFormat, tanggalTransaksi)
										if err != nil {
											fmt.Println("Error:", err)
											return
										}
										result, permit := transaksi.AddTransaksi(result.ID, no_Hp, metodeID, newTanggal)
										if permit {
											fmt.Println("Data Transaksi berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("No Nota: %d\nNama Editor: %d\n", result.No_nota, result.UserID)
										}
									case 2:
										result, permit := transaksi.ShowTransaksi(result.ID)
										if permit {
											for _, m := range result {
												fmt.Println("===Daftar Customer===")
												fmt.Printf("No Nota: %d\nTanggal Transaksi : %s\nUser ID: %d\nNo HP: %d\nID Metode :%d\n", m.No_nota, m.Tanggal_transaksi, m.UserID, m.No_hp, m.Id_metode)
											}
										}
									case 3:
										var no_Nota uint
										var no_Hp uint
										var metodeID uint
										var tanggalTransaksi string
										fmt.Print("Masukkan No Nota yang akan diperbarui: ")
										fmt.Scanln(&no_Nota)
										fmt.Print("Masukkan No HP baru: ")
										fmt.Scanln(&no_Hp)
										fmt.Print("Masukkan ID Metode baru: ")
										fmt.Scanln(&metodeID)
										fmt.Print("Masukkan Tanggal Transaksi baru: ")
										fmt.Scanln(&tanggalTransaksi)
										var dateFormat = "2006-01-02"
										newTanggal, err := time.Parse(dateFormat, tanggalTransaksi)
										if err != nil {
											fmt.Println("Error:", err)
											return
										}
										result, permit := transaksi.UpdateTransaksi(result.ID, no_Nota, no_Hp, metodeID, newTanggal)
										if permit {
											fmt.Println("Data dengan No. Nota ini berhasil diperbarui dengan detail berikut:")
											fmt.Printf("No Nota: %d\nTanggal Transaksi : %s\nUser ID: %d\nNo HP: %d\nID Metode :%d\n", result.No_nota, result.Tanggal_transaksi, result.UserID, result.No_hp, result.Id_metode)
										}
									case 4:
										var no_Nota uint
										fmt.Print("Masukkan No Hp Customer yang akan dihapus: ")
										fmt.Scanln(&no_Nota)

										permit := transaksi.DeleteTransaksi(result.ID, no_Nota)
										if permit {
											fmt.Println("Data Customer berhasil dihapus")
										}
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
							case 6:
								for permit {
									fmt.Println("======Menu Detail Transaksi========")
									fmt.Println("1. Tambah Data Transaksi Detail")
									fmt.Println("2. Cari Detail Transaksi")
									fmt.Println("3. Update Detail Transaksi")
									fmt.Println("4. Hapus Detail Transaksi")
									fmt.Println("0. Kembali")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 0:
										permit = false
									case 1:
										resulBarang, permitBarang := barang.ShowBarang(result.ID)
										if permitBarang {
											for _, b := range resulBarang {
												fmt.Println("===Daftar Barang===")
												fmt.Printf("ID: %d\nNama: %s\nDeskripsi:%s\nHarga: %.2f\nStok: %d\nNama Editor: %d \n", b.ID, b.Nama_barang, b.Desc_barang, b.Harga_barang, b.Stock, b.UserID)
											}
										}
										fmt.Println("====Masukan Detail Transaksi====")
										var notaTransaksi uint
										var idBarang uint
										var jumlahBarang uint
										var statusPembayaran string
										var totalHarga float64
										fmt.Print("Masukkan Nota Transaksi: ")
										fmt.Scanln(&notaTransaksi)
										fmt.Print("Masukkan ID Barang: ")
										fmt.Scanln(&idBarang)
										fmt.Print("Masukkan Jumlah Barang: ")
										fmt.Scanln(&jumlahBarang)
										fmt.Print("Masukkan Status Pembayaran: ")
										fmt.Scanln(&statusPembayaran)
										resultTransaksi, permitTransaksi := transaksi_detail.AddTransaksiDetail(notaTransaksi, idBarang, jumlahBarang, totalHarga, statusPembayaran)
										if permitTransaksi {
											fmt.Println("Data Transaksi berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("No Nota: %d\nID Barang: %d\nJumlah Barang: %d\nTotal Harga: %.2f\nStatus Pembayaran: %s\n",
												resultTransaksi.Nota_transaksi, resultTransaksi.Id_barang, resultTransaksi.Jumlah_barang, resultTransaksi.Total_harga, resultTransaksi.Status_pembayaran)
										}
									case 2:
										var notaTransaksi uint
										fmt.Print("Masukkan Nota Transaksi: ")
										fmt.Scanln(&notaTransaksi)
										result, permit := transaksi_detail.ShowTransaksiDetail(notaTransaksi)
										if permit {
											for _, transaksi := range result {
												fmt.Println("===Daftar Detail Transaksi===")
												fmt.Printf("Nota Transaksi: %d, ID Barang: %d, Jumlah Barang: %d, Total Harga: %f, Status Pembayaran: %s\n",
													transaksi.Nota_transaksi, transaksi.Id_barang, transaksi.Jumlah_barang, transaksi.Total_harga, transaksi.Status_pembayaran)
											}
										}
									case 3:
										var notaTransaksi uint
										var idBarang uint
										var jumlahBarang uint
										var statusPembayaran string
										var totalHarga float64
										fmt.Print("Masukkan Nota Transaksi: ")
										fmt.Scanln(&notaTransaksi)
										fmt.Print("Masukkan ID Barang: ")
										fmt.Scanln(&idBarang)
										fmt.Print("Masukkan Jumlah Barang: ")
										fmt.Scanln(&jumlahBarang)
										fmt.Print("Masukkan Status Pembayaran: ")
										fmt.Scanln(&statusPembayaran)
										resultTransaksi, permitTransaksi := transaksi_detail.UpdateTransaksiDetail(notaTransaksi, idBarang, jumlahBarang, totalHarga, statusPembayaran)
										if permitTransaksi {
											fmt.Println("Data Transaksi berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("No Nota: %d\nID Barang: %d\nJumlah Barang: %d\nTotal Harga: %.2f\nStatus Pembayaran: %s\n",
												resultTransaksi.Nota_transaksi, resultTransaksi.Id_barang, resultTransaksi.Jumlah_barang, resultTransaksi.Total_harga, resultTransaksi.Status_pembayaran)
										}
									case 4:
										var notaTransaksi uint
										var idBarang uint
										fmt.Print("Masukkan Nota Transaksi yang akan dihapus detailnya: ")
										fmt.Scanln(&notaTransaksi)
										fmt.Print("Masukkan ID Barang yang akan dihapus detailnya: ")
										fmt.Scanln(&idBarang)

										permit := transaksi_detail.DeleteTransaksiDetail(notaTransaksi, idBarang)
										if permit {
											fmt.Println("Data Customer berhasil dihapus")
										}
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
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
							fmt.Println("4. Menu Transaksi")
							fmt.Println("5. Menu Detail Transaksi")
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
									fmt.Println("4. Update Stock")
									fmt.Println("0. Kembali")
									fmt.Print("Pilihan: ")

									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 0:
										permit = false
									case 1:
										result, permit := barang.AddBarang(result.ID)
										if permit {
											fmt.Println("Barang berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("ID: %d\nNama: %s\nHarga: %.2f\nStok: %d\n", result.ID, result.Nama_barang, result.Harga_barang, result.Stock)
										}
									case 2:
										result, permit := barang.ShowBarang(result.ID)
										if permit {
											for _, b := range result {
												fmt.Println("===Daftar Barang===")
												fmt.Printf("ID: %d\nNama: %s\nDeskripsi:%s\nHarga: %.2f\nStok: %d\nNama Editor: %d \n", b.ID, b.Nama_barang, b.Desc_barang, b.Harga_barang, b.Stock, b.UserID)
											}
										}
									case 3:
										var barangID uint
										fmt.Print("Masukkan ID barang yang akan diperbarui: ")
										fmt.Scanln(&barangID)
										result, permit := barang.UpdateBarang(result.ID, barangID)
										if permit {
											fmt.Println("Barang berhasil diperbarui dengan detail berikut:")
											fmt.Printf("ID: %d\nNama: %s\nDeskripsi: %s\nHarga: %.2f\nStok: %d\n Nama Editor:%d", result.ID, result.Nama_barang, result.Desc_barang, result.Harga_barang, result.Stock, result.UserID)
										}
									case 4:
										var barangID uint
										fmt.Print("Masukkan ID barang yang akan diperbarui: ")
										fmt.Scanln(&barangID)
										result, permit := barang.UpdateStock(result.ID, barangID)
										if permit {
											fmt.Println("Stock Barang berhasil diperbarui dengan detail berikut:")
											fmt.Printf("ID: %d\nNama: %s\nStok: %d\n Nama Editor:%d", result.ID, result.Nama_barang, result.Stock, result.UserID)
										}
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
									fmt.Println("0. Kembali")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 0:
										permit = false
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
									fmt.Println("0. Kembali")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 0:
										permit = false
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
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
							case 4:
								for permit {
									fmt.Println("======Menu Transaksi========")
									fmt.Println("1. Tambah Data Transaksi")
									fmt.Println("2. Tampilkan Daftar Transaksi")
									fmt.Println("0. Kembali")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 0:
										permit = false
									case 1:
										var no_Hp uint
										var metodeID uint
										var tanggalTransaksi string
										fmt.Print("Masukkan No HP Customer: ")
										fmt.Scanln(&no_Hp)
										fmt.Print("Masukkan ID Metode Pembayaran: ")
										fmt.Scanln(&metodeID)
										fmt.Print("Masukkan Tanggal Transaksi baru: ")
										fmt.Scanln(&tanggalTransaksi)
										var dateFormat = "2006-01-02"
										newTanggal, err := time.Parse(dateFormat, tanggalTransaksi)
										if err != nil {
											fmt.Println("Error:", err)
											return
										}
										result, permit := transaksi.AddTransaksi(result.ID, no_Hp, metodeID, newTanggal)
										if permit {
											fmt.Println("Data Transaksi berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("No Nota: %d\nNama Editor: %d\n", result.No_nota, result.UserID)
										}
									case 2:
										result, permit := transaksi.ShowTransaksi(result.ID)
										if permit {
											for _, m := range result {
												fmt.Println("===Daftar Customer===")
												fmt.Printf("No Nota: %d\nTanggal Transaksi : %s\nUser ID: %d\nNo HP: %d\nID Metode :%d\n", m.No_nota, m.Tanggal_transaksi, m.UserID, m.No_hp, m.Id_metode)
											}
										}
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
							case 5:
								for permit {
									fmt.Println("======Menu Detail Transaksi========")
									fmt.Println("1. Tambah Data Transaksi Detail")
									fmt.Println("2. Cari Detail Transaksi")
									fmt.Println("0. Kembali")
									fmt.Println("5. Keluar")
									fmt.Print("Pilihan: ")
									var choice int
									fmt.Scanln(&choice)

									switch choice {
									case 0:
										permit = false
									case 1:
										resulBarang, permitBarang := barang.ShowBarang(result.ID)
										if permitBarang {
											for _, b := range resulBarang {
												fmt.Println("===Daftar Barang===")
												fmt.Printf("ID: %d\nNama: %s\nDeskripsi:%s\nHarga: %.2f\nStok: %d\nNama Editor: %d \n", b.ID, b.Nama_barang, b.Desc_barang, b.Harga_barang, b.Stock, b.UserID)
											}
										}
										fmt.Println("====Masukan Detail Transaksi====")
										var notaTransaksi uint
										var idBarang uint
										var jumlahBarang uint
										var statusPembayaran string
										var totalHarga float64
										fmt.Print("Masukkan Nota Transaksi: ")
										fmt.Scanln(&notaTransaksi)
										fmt.Print("Masukkan ID Barang: ")
										fmt.Scanln(&idBarang)
										fmt.Print("Masukkan Jumlah Barang: ")
										fmt.Scanln(&jumlahBarang)
										fmt.Print("Masukkan Status Pembayaran: ")
										fmt.Scanln(&statusPembayaran)
										resultTransaksi, permitTransaksi := transaksi_detail.AddTransaksiDetail(notaTransaksi, idBarang, jumlahBarang, totalHarga, statusPembayaran)
										if permitTransaksi {
											fmt.Println("Data Transaksi berhasil ditambahkan dengan detail berikut:")
											fmt.Printf("No Nota: %d\nID Barang: %d\nJumlah Barang: %d\nTotal Harga: %.2f\nStatus Pembayaran: %s\n",
												resultTransaksi.Nota_transaksi, resultTransaksi.Id_barang, resultTransaksi.Jumlah_barang, resultTransaksi.Total_harga, resultTransaksi.Status_pembayaran)
										}
									case 2:
										var notaTransaksi uint
										fmt.Print("Masukkan Nota Transaksi: ")
										fmt.Scanln(&notaTransaksi)
										result, permit := transaksi_detail.ShowTransaksiDetail(notaTransaksi)
										if permit {
											for _, transaksi := range result {
												fmt.Println("===Daftar Detail Transaksi===")
												fmt.Printf("Nota Transaksi: %d, ID Barang: %d, Jumlah Barang: %d, Total Harga: %f, Status Pembayaran: %s\n",
													transaksi.Nota_transaksi, transaksi.Id_barang, transaksi.Jumlah_barang, transaksi.Total_harga, transaksi.Status_pembayaran)
											}
										}
									default:
										fmt.Println("Pilihan tidak valid. Silakan pilih lagi.")
									}
								}
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
