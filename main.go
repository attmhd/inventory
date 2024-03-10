package main

import (
	"inventory-manajemen-system/config"
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

	//Load Database
	config.LoadConfig()
	config.DBConnection()

	//Set up router(gin)
	r := gin.Default()
	api := r.Group("/api")

	//handler
	api.GET("", handler.IndexHandler)
	api.POST("/user", handler.UserHandler)

	//start server
	err := r.Run()
	if err != nil {
		panic(err.Error())
	}
}
