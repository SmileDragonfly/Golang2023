package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gormtest/GormTest/model"
	"log"
)

func main() {
	dsn := "host=localhost user=postgres password=123@123A dbname=gormtest port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&model.Wallet{}, &model.Bank{}, &model.LinkInfo{})
	//wallet := model.Wallet{NationalID: "123456789", Mobile: "0358698789", Gender: "Male", Language: "VietNam", UserName: "User1"}
	//db.Create(&wallet)
	//db.Create(&Wallet{})
	var wallet model.Wallet
	db.First(&wallet, 1)
	log.Println(wallet)
}
