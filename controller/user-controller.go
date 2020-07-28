package controller

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserController interface {
	InsertUser(ctx *gin.Context) error
	GetAllUsers() []entity.User
}

//Controller is
type userController struct {
	service service.UserService
}

var userValidate *validator.Validate

func NewUser(service service.UserService) UserController {
	userValidate = validator.New()
	return &userController{service: service}
}

func (c *userController) InsertUser(ctx *gin.Context) error {
	var user entity.User
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		return err
	}

	err = userValidate.Struct(user)
	if err != nil {
		return err
	}
	c.service.InsertUser(user)
	return nil
}

func (c *userController) GetAllUsers() []entity.User {
	return c.service.GetAllUsers()
}
