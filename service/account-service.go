package service

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/repository"
)

type AccountService interface {
	AddAccountDetails(account entity.Account) error
	GetAccountDetails(accountID int64) (entity.Account, error)
	UpdateAccountDetails(accountID int64, account entity.Account) error
}

type accountService struct {
	accountRepository repository.DataStore
}

func NewAccountService(repo repository.DataStore) AccountService {
	return &accountService{accountRepository: repo}
}

func (a accountService) AddAccountDetails(account entity.Account) error {
	return a.accountRepository.AddAccountDetails(account)
}

func (a accountService) GetAccountDetails(accountID int64) (entity.Account, error) {
	return a.accountRepository.GetAccountDetails(accountID)
}

func (a accountService) UpdateAccountDetails(accountID int64, account entity.Account) error {
	return a.accountRepository.UpdateAccountDetails(accountID, account)
}
