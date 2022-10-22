package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/titoyudha/go_ecommerce_api/dto"
	"github.com/titoyudha/go_ecommerce_api/helper"
	"github.com/titoyudha/go_ecommerce_api/service"
)

type UserHandler interface {
	GetUser(ctx *gin.Context) error
	GetUserByID(ctx *gin.Context) error
	CreateUser(ctx *gin.Context) error
	DeleteUserByID(ctx *gin.Context) error
}

type userHandler struct {
	UserService service.UserService
}

func NewHttpHandler(userService service.UserService) UserHandler {
	return &userHandler{
		UserService: userService,
	}
}

// CreateUser implements UserHandler
func (services userHandler) CreateUser(ctx *gin.Context) error {
	userReqBody := new(dto.UserRequestBody)
	if err := ctx.ShouldBindJSON(userReqBody); err != nil {
		helper.Log.Errorln(err)
		return helper.JSONResponse(ctx, http.StatusUnprocessableEntity, err.Error(), nil)
	}
	user, err := services.UserService.CreateNewUser(*userReqBody)
	if err != nil {
		helper.Log.Errorln(err)
		return helper.JSONResponse(ctx, http.StatusBadRequest, err.Error(), nil)
	}
	userResponse := dto.ParseFromEntity(user)
	return helper.JSONResponse(ctx, http.StatusCreated, "", userResponse)
}

// DeleteUserByID implements UserHandler
func (services userHandler) DeleteUserByID(ctx *gin.Context) error {
	userID := ctx.Param("user_id")
	if userID != "" {
		helper.Log.Error()
		return helper.JSONResponse(ctx, http.StatusBadRequest, "Error Finding User ID", nil)
	}
	err := services.UserService.DeleteUser(userID)
	if err != nil {
		helper.Log.Errorln(err)
		return helper.JSONResponse(ctx, http.StatusBadRequest, "Failed Deleting User", nil)
	}
	return helper.JSONResponse(ctx, http.StatusOK, "Success Delete User", nil)
}

// GetUser implements UserHandler
func (services userHandler) GetUser(ctx *gin.Context) error {
	//get all user
	users := services.UserService.GetUser()
	userResponse := dto.ParseFromEntities(users)

	return helper.JSONResponse(ctx, http.StatusOK, "", userResponse)
}

// GetUserByID implements UserHandler
func (service userHandler) GetUserByID(ctx *gin.Context) error {
	userID := ctx.Param("user_id")
	if userID != "" {
		helper.Log.Error()
		return helper.JSONResponse(ctx, http.StatusBadRequest, "Error Finding User ID", nil)
	}

	user := service.UserService.GetUserByID(userID)
	userResponse := dto.ParseFromEntity(user)
	return helper.JSONResponse(ctx, http.StatusOK, "Succes Find User", userResponse)

}
