package controller

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"strconv"
)

// ExpenseController interface
type ExpenseController interface {
	AddExpense(ctx *gin.Context) error
	GetAllExpense() []entity.Expense
	GetExpense(ctx *gin.Context) (entity.Expense, error)
	UpdateExpense(ctx *gin.Context) error
	DeleteExpense(ctx *gin.Context) error
}

type expenseController struct {
	service service.ExpenseService
}

var expenseValidate *validator.Validate

// NewExpense expenseController
func NewExpense(service service.ExpenseService) ExpenseController {
	expenseValidate = validator.New()
	return &expenseController{service: service}
}

// AddExpense adding expense
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

// GetAllExpense get all expenses
func (ec *expenseController) GetAllExpense() []entity.Expense {
	return ec.service.GetAllExpense()
}

func (ec *expenseController) GetExpense(ctx *gin.Context) (entity.Expense, error) {
	var expense entity.Expense
	err := ctx.ShouldBindJSON(&expense)
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return expense, err
	}
	expense.ExpenseID = id
	return ec.service.GetExpense(id)
}

func (ec *expenseController) UpdateExpense(ctx *gin.Context) error {
	var expense entity.Expense
	err := ctx.ShouldBindJSON(&expense)
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	expense.ExpenseID = id
	return ec.service.UpdateExpense(id, expense)
}

func (ec *expenseController) DeleteExpense(ctx *gin.Context) error {
	var expense entity.Expense
	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	expense.ExpenseID = id
	return ec.service.DeleteExpense(id)
}
