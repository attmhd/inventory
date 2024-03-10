package config

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	dsn := fmt.Sprintf("host= %v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Shanghai", ENV.DB_URL, ENV.DB_Username, ENV.DB_Password, ENV.DB_Name, ENV.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("koneksi gagal")
	}

	fmt.Println("Koneksi Berhasil")

	DB = db

}
