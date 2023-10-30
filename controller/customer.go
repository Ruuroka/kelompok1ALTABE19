package controller

import (
	"bufio"
	"fmt"
	"kelompok1ALTABE19/model"
	"os"
	"strings"

	"gorm.io/gorm"
)

type CustomerSystem struct {
	DB *gorm.DB
}

func (ts *CustomerSystem) AddCustomer() (model.Customer, bool) {
	var newCustomer = new(model.Customer)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("No HP : ")
	fmt.Scanln(&newCustomer.No_hp)
	fmt.Print("Masukan Nama : ")
	newCustomer.Nama, _ = reader.ReadString('\n')
	newCustomer.Nama = strings.TrimSpace(newCustomer.Nama)
	fmt.Print("Masukan Alamat : ")
	newCustomer.Alamat, _ = reader.ReadString('\n')
	newCustomer.Alamat = strings.TrimSpace(newCustomer.Alamat)
	fmt.Print("Masukan Email : ")
	newCustomer.Email, _ = reader.ReadString('\n')
	newCustomer.Email = strings.TrimSpace(newCustomer.Email)

	err := ts.DB.Create(newCustomer).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Customer{}, false
	}
	return *newCustomer, true
}

func (ts *CustomerSystem) ShowCustomer() ([]model.Customer, bool) {
	var customer []model.Customer
	// ts.DB.InnerJoins("INNER JOIN users ON barangs.user_id = users.id").Find(&barang)
	ts.DB.Find(&customer)

	if len(customer) == 0 {
		fmt.Println("Daftar barang kosong.")
		return nil, false
	}
	return customer, true
}

func (ts *CustomerSystem) UpdateCustomer(no_Hp uint) (model.Customer, bool) {
	exitingCustomer := model.Customer{}
	err := ts.DB.First(&exitingCustomer, no_Hp).Error
	if err != nil {
		return model.Customer{}, false
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Masukan Nama : ")
	exitingCustomer.Nama, _ = reader.ReadString('\n')
	exitingCustomer.Nama = strings.TrimSpace(exitingCustomer.Nama)
	fmt.Print("Masukan Alamat : ")
	exitingCustomer.Alamat, _ = reader.ReadString('\n')
	exitingCustomer.Alamat = strings.TrimSpace(exitingCustomer.Alamat)
	fmt.Print("Masukan Email : ")
	exitingCustomer.Email, _ = reader.ReadString('\n')
	exitingCustomer.Email = strings.TrimSpace(exitingCustomer.Email)

	err = ts.DB.Save(&exitingCustomer).Error
	if err != nil {
		return model.Customer{}, false
	}

	return exitingCustomer, true
}

func (ts *CustomerSystem) DeleteCustomer(no_Hp uint) bool {
	exitingCustomer := model.Customer{}
	err := ts.DB.Where("no_hp = ?", no_Hp).First(&exitingCustomer).Error
	if err != nil {
		return false
	}

	err = ts.DB.Delete(&exitingCustomer).Error
	if err != nil {
		return false
	}

	return true
}
