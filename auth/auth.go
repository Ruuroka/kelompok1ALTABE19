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

func (as *AuthSystem) Show(userID uint) ([]model.User, bool) {
	var user []model.User
	as.DB.Find(&user)
	if len(user) == 0 {
		fmt.Println("Daftar User kosong.")
		return nil, false
	}
	return user, true
}

func (as *AuthSystem) Update(userID uint) (model.User, bool) {
	var existingUser = model.User{}
	err := as.DB.First(&existingUser, userID).Error
	if err != nil {
		return model.User{}, false
	}
	fmt.Print("Masukkan Nama = ")
	fmt.Scanln(&existingUser.Nama)
	fmt.Print("Masukkan Password = ")
	fmt.Scanln(&existingUser.Password)
	// err := as.DB.Table("pelanggan").Create(existingUser).Error
	err = as.DB.Save(existingUser).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.User{}, false
	}

	return existingUser, true
}

func (as *AuthSystem) Delete(userID uint) bool {
	existingUser := model.User{}
	err := as.DB.Where("id = ?", userID).First(&existingUser).Error
	if err != nil {
		return false
	}

	err = as.DB.Delete(&existingUser).Error
	if err != nil {
		return false
	}

	return true
}
