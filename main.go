package main

import (
	"fmt"
	"kelompok1ALTABE19/auth"
	"kelompok1ALTABE19/config"
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
							fmt.Println("1. Tambahakn akun Pegawai")
							fmt.Println("2. Show List todo")
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
							fmt.Println("2. Show List todo")
							fmt.Println("3. Update todo")
							fmt.Println("4. Delete todo")
							fmt.Println("0. Logout")
							fmt.Println("99. Exit")
							fmt.Print("Masukkan pilihan:")
							fmt.Scanln(&menuLogin)
							switch menuLogin {
							case 1:
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
