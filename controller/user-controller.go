package controller

import (
	"strconv"

	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//UserController having user function
type UserController interface {
	InsertUser(ctx *gin.Context) error
	GetAllUsers() []entity.User
	GetUser(ctx *gin.Context) entity.User
	UpdateUser(ctx *gin.Context) error
	DeleteUser(ctx *gin.Context) error
}

//userController is having serservice
type userController struct {
	service service.UserService
}

var userValidate *validator.Validate

//NewUser implementing userController
func NewUser(service service.UserService) UserController {
	userValidate = validator.New()
	return &userController{service: service}
}

//InsertUser for create user
func (uc *userController) InsertUser(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	err = userValidate.Struct(user)
	if err != nil {
		return err
	}
	uc.service.InsertUser(user)
	return nil
}

//GetAllUsers get all users
func (uc *userController) GetAllUsers() []entity.User {
	return uc.service.GetAllUsers()
}

//GetUser get user
func (uc *userController) GetUser(ctx *gin.Context) entity.User {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)

	if err != nil {
		return entity.User{}
	}

	user.ID = id
	err = userValidate.Struct(user)

	if err != nil {
		return entity.User{}
	}

	return uc.service.GetUser(user.ID)
}

//UpdateUser update user
func (uc *userController) UpdateUser(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		return err
	}
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)

	if err != nil {
		return err
	}
	user.ID = id

	err = userValidate.Struct(user)

	if err != nil {
		return err
	}
	uc.service.UpdateUser(user.ID, user)
	return nil
}

//DeleteUser delete user
func (uc *userController) DeleteUser(ctx *gin.Context) error {
	var user entity.User
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)

	if err != nil {
		return err
	}
	user.ID = id

	uc.service.DeleteUser(user.ID)
	return nil
}
