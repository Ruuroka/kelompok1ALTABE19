package auth

import (
	"fmt"
	"kelompok1ALTABE19/model"

	"gorm.io/gorm"
)

type AuthSystem struct {
	DB *gorm.DB
}

func (as *AuthSystem) Login() (model.User, bool) {
	var currentUser = new(model.User)
	fmt.Print("Masukkan Nama = ")
	fmt.Scanln(&currentUser.Nama)
	fmt.Print("Masukkan Password = ")
	fmt.Scanln(&currentUser.Password)

	// qry := as.DB.Table("pelanggan").Where("hp = ?", hp).Take(currentUser)
	qry := as.DB.Where("nama = ? AND password = ?", currentUser.Nama, currentUser.Password).Take(currentUser)

	err := qry.Error

	if err != nil {
		fmt.Println("login process error:", err.Error())
		return model.User{}, false
	}

	return *currentUser, true
}

func (as *AuthSystem) Register() (model.User, bool) {
	var newUser = new(model.User)
	fmt.Print("Masukkan Nama = ")
	fmt.Scanln(&newUser.Nama)
	fmt.Print("Masukkan Password = ")
	fmt.Scanln(&newUser.Password)
	// err := as.DB.Table("pelanggan").Create(newUser).Error
	err := as.DB.Create(newUser).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.User{}, false
	}

	return *newUser, true
}
