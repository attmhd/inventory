package main

import (
	"inventory-manajemen-system/admingd"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=localhost user=postgres password=postgres dbname=inventory port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	tstRepo := admingd.NewRepository(db)
	user := admingd.Admgd{
		Id:   1,
		Name: "Testing dlu",
	}

	tstRepo.Save(user)

	//fmt.Println("Koneksi sukses")

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080
}
