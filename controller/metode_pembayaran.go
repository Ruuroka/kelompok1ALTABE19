package controller

import (
	"bufio"
	"fmt"
	"kelompok1ALTABE19/model"
	"os"
	"strings"

	"gorm.io/gorm"
)

type MetodeSystem struct {
	DB *gorm.DB
}

func (ts *MetodeSystem) AddMetode() (model.Metode_Pembayaran, bool) {
	var newMetode = new(model.Metode_Pembayaran)
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Nama Metode Pembayaran: ")
	newMetode.Method_name, _ = reader.ReadString('\n')
	newMetode.Method_name = strings.TrimSpace(newMetode.Method_name)

	err := ts.DB.Create(newMetode).Error
	if err != nil {
		fmt.Println("input error:", err.Error())
		return model.Metode_Pembayaran{}, false
	}
	return *newMetode, true
}

func (ts *MetodeSystem) ShowMetode() ([]model.Metode_Pembayaran, bool) {
	var metode []model.Metode_Pembayaran
	// ts.DB.InnerJoins("INNER JOIN users ON barangs.user_id = users.id").Find(&barang)
	ts.DB.Find(&metode)

	if len(metode) == 0 {
		fmt.Println("Daftar barang kosong.")
		return nil, false
	}
	return metode, true
}

func (ts *MetodeSystem) UpdateMetode(metodeID uint) (model.Metode_Pembayaran, bool) {
	existingMetode := model.Metode_Pembayaran{}
	err := ts.DB.First(&existingMetode, metodeID).Error
	if err != nil {
		return model.Metode_Pembayaran{}, false
	}
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Nama Metode: ")
	existingMetode.Method_name, _ = reader.ReadString('\n')
	existingMetode.Method_name = strings.TrimSpace(existingMetode.Method_name)

	err = ts.DB.Save(&existingMetode).Error
	if err != nil {
		return model.Metode_Pembayaran{}, false
	}

	return existingMetode, true
}

func (ts *MetodeSystem) DeleteMetode(metodeID uint) bool {
	existingMetode := model.Metode_Pembayaran{}
	err := ts.DB.Where("id = ?", metodeID).First(&existingMetode).Error
	if err != nil {
		return false
	}

	err = ts.DB.Delete(&existingMetode).Error
	if err != nil {
		return false
	}

	return true
}
