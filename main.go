package main

import (
	"inventory-manajemen-system/handler"

	"github.com/gin-gonic/gin"
)

func main() {

	// tstRepo := admin.NewRepository(db)
	// user := admin.Admgd{
	// 	Id:   1,
	// 	Name: "Testing dlu",
	// }

	// tstRepo.Save(user)

	//fmt.Println("Koneksi sukses")

	// r := gin.Default()
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	// r.Run() // listen and serve on 0.0.0.0:8080

	r := gin.Default()

	r.GET("", handler.IndexHandler)

	err := r.Run()
	if err != nil {
		panic(err.Error())
	}
}
