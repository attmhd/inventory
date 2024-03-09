package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "HElo Niadu")
}

func LoginHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "ini Halaman Login")
}
