package handler

import (
	"inventory-manajemen-system/dto"
	"inventory-manajemen-system/entity"
	"inventory-manajemen-system/helper"
	"inventory-manajemen-system/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService service.Service
}

func NewUserHandler(userService service.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) AddUser(ctx *gin.Context) {
	var input entity.UserInput

	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}

	user, err := h.userService.AddUser(input)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, nil)
	}

	res := helper.Respons(dto.ResponsParams{
		StatusCode: http.StatusCreated,
		Message:    "Account has been added succesfully",
		Data:       user,
	})

	ctx.JSON(http.StatusOK, res)

}
