package service

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/repository"
)

type ExpenseService interface {
	AddExpense(expense entity.Expense) int64
	GetAllExpense() []entity.Expense
	GetExpense(id int64) (entity.Expense, error)
	UpdateExpense(id int64, expense entity.Expense) error
	DeleteExpense(id int64) error
}

type expenseService struct {
	expenseRepository repository.DataStore
}

func NewExpense(repo repository.DataStore) ExpenseService {
	return &expenseService{expenseRepository: repo}
}

func (e expenseService) AddExpense(expense entity.Expense) int64 {
	e.expenseRepository.AddExpense(expense)
	return expense.ExpenseID
}

func (e expenseService) GetAllExpense() []entity.Expense {
	return e.expenseRepository.GetAllExpense()
}

func (e expenseService) GetExpense(id int64) (entity.Expense, error) {
	return e.expenseRepository.GetExpense(id)
}

func (e expenseService) UpdateExpense(id int64, expense entity.Expense) error {
	return e.expenseRepository.UpdateExpense(id, expense)
}

func (e expenseService) DeleteExpense(id int64) error {
	return e.expenseRepository.DeleteExpense(id)
}
