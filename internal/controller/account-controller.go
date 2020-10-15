package controller

import (
	"github.com/GoGinApi/v2/internal/entity"
	"github.com/GoGinApi/v2/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

type AccountController interface {
	AddAccountDetails(ctx *gin.Context) error
	GetAccountDetails(ctx *gin.Context) (entity.Account, error)
	UpdateAccountDetails(ctx *gin.Context) error
}

type accountController struct {
	service service.AccountService
}

var _ *validator.Validate

func NewAccount(service service.AccountService) AccountController {
	_ = validator.New()
	return &accountController{service: service}
}

func (ac *accountController) AddAccountDetails(ctx *gin.Context) error {
	var account entity.Account
	err := ctx.ShouldBindJSON(&account)
	if err != nil {
		return nil
	}
	ac.service.AddAccountDetails(account)
	return nil
}

func (ac *accountController) GetAccountDetails(ctx *gin.Context) (entity.Account, error) {
	var account entity.Account
	ctx.ShouldBindJSON(&account)
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return entity.Account{}, nil
	}
	account.ClientID = int(id)
	return ac.service.GetAccountDetails(id)
}

func (ac *accountController) UpdateAccountDetails(ctx *gin.Context) error {
	var account entity.Account
	ctx.ShouldBindJSON(&account)
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	account.ClientID = int(id)
	return ac.service.UpdateAccountDetails(id, account)
}
