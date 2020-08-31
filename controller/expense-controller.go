package controller

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

//ExpenseController interface
type ExpenseController interface {
	AddExpense(ctx *gin.Context) error
	GetAllExpense() []entity.Expense
}

type expenseController struct {
	service service.ExpenseService
}

var expenseValidate *validator.Validate

//NewExpense expenseController
func NewExpense(service service.ExpenseService) ExpenseController {
	expenseValidate = validator.New()
	return &expenseController{service: service}
}

//AddExpense adding expense
func (ec *expenseController) AddExpense(ctx *gin.Context) error {
	var expense entity.Expense
	err := ctx.ShouldBindJSON(&expense)
	if err != nil {
		return err
	}

	err = expenseValidate.Struct(expense)
	if err != nil {
		return err
	}
	ec.service.AddExpense(expense)
	return nil
}

//GetAllExpense get all expenses
func (ec *expenseController) GetAllExpense() []entity.Expense {
	return ec.service.GetAllExpense()
}
