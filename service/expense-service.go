package service

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/repository"
)

type ExpenseService interface {
	AddExpense(expense entity.Expense) int64
	GetAllExpense() []entity.Expense
}

type expenseService struct {
	expenseRepository repository.ExpenseRepository
}

func NewExpense(repo repository.ExpenseRepository) ExpenseService {
	return &expenseService{expenseRepository: repo}
}

func (e expenseService) AddExpense(expense entity.Expense) int64 {
	e.expenseRepository.AddExpense(expense)
	return expense.ExpenseID
}

func (e expenseService) GetAllExpense() []entity.Expense {
	return e.expenseRepository.GetAllExpense()
}
