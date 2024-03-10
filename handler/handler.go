package handler

import (
	"net/http"

	"inventory-manajemen-system/model"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HElo Niadu")
}

func LoginHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ini Halaman Login")
}

func UserHandler(ctx *gin.Context) {
	user := new(model.User)
	err := ctx.Bind(user)
	if err != nil {
		panic(err)
	}

	ctx.JSON(http.StatusOK, user)
}
