package controller

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ExpenseController interface {
	AddExpense(ctx *gin.Context) error
	GetAllExpense() []entity.Expense
}

type expenseController struct {
	service service.ExpenseService
}

var expenseValidate *validator.Validate

func NewExpense(service service.ExpenseService) ExpenseController {
	expenseValidate = validator.New()
	return &expenseController{service: service}
}

func (e *expenseController) AddExpense(ctx *gin.Context) error {
	var expense entity.Expense
	err := ctx.ShouldBindJSON(&expense)
	if err != nil {
		return err
	}

	err = expenseValidate.Struct(expense)
	if err != nil {
		return err
	}
	e.service.AddExpense(expense)
	return nil
}

func (e *expenseController) GetAllExpense() []entity.Expense {
	return e.service.GetAllExpense()
}
