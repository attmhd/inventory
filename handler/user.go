package handler

import (
	"inventory-manajemen-system/dto"
	"inventory-manajemen-system/entity"
	"inventory-manajemen-system/helper"
	"inventory-manajemen-system/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type userHandler struct {
	userService service.Service
}

func NewUserHandler(userService service.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) GetUser(ctx *gin.Context) {
	result, err := h.userService.GetUser()
	if err != nil {
		res := helper.Respons(dto.ResponsParams{
			StatusCode: http.StatusBadRequest,
			Message:    "Register account has been failed",
		})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.Respons(dto.ResponsParams{
		StatusCode: http.StatusCreated,
		Message:    "Account has been added succesfully",
		Data:       result,
	})

	ctx.JSON(http.StatusOK, res)

}

func (h *userHandler) AddUser(ctx *gin.Context) {
	var input entity.UserInput

	if err := ctx.ShouldBind(&input); err != nil {

		var errors []string

		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, e.Error())
		}

		errorMsg := gin.H{"error": errors}

		res := helper.Respons(dto.ResponsParams{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    errorMsg
		})
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	user, err := h.userService.AddUser(input)
	if err != nil {
		res := helper.Respons(dto.ResponsParams{
			StatusCode: http.StatusBadRequest,
			Message:    "Register account has been failed",
		})
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := helper.Respons(dto.ResponsParams{
		StatusCode: http.StatusCreated,
		Message:    "Account has been added succesfully",
		Data:       user,
	})

	ctx.JSON(http.StatusOK, res)

}

// func (h *userHandler) DeleteUser(ctx *gin.Context) {
// 	id := ctx.Params("id")
// 	ids, err := strconv.Atoi(id)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, nil)
// 	}
// 	h.userService.DeleteUser(ids)

// 	ctx.JSON(http.StatusOK, "ok")
// }
