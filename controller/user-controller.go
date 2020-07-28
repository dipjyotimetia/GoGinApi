package controller

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type UserController interface {
	InsertUser(ctx *gin.Context) error
	GetAllUsers() []entity.User
	GetUser(ctx *gin.Context) entity.User
	UpdateUser(ctx *gin.Context) error
	DeleteUser(ctx *gin.Context) error
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

func (c *userController) GetUser(ctx *gin.Context) entity.User {
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

	return c.service.GetUser(user.ID)
}

func (c *userController) UpdateUser(ctx *gin.Context) error {
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
	c.service.UpdateUser(user.ID, user)
	return nil
}

func (c *userController) DeleteUser(ctx *gin.Context) error {
	var user entity.User
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)

	if err != nil {
		return err
	}
	user.ID = id

	c.service.DeleteUser(user.ID)
	return nil
}
