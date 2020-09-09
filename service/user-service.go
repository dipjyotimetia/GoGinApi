package service

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/repository"
)

type UserService interface {
	ResetPassword(resetPassword entity.ResetPassword) error
	Create(user entity.Register) error
	Login(name, email, password, createdAt, updatedAt string, user entity.Login) error
	CheckUserExist(user entity.Register) bool
	CheckAndRetrieveUserIDViaEmail(createReset entity.CreateReset) (int, bool)
}

type userService struct {
	userRepository repository.DataStore
}

func NewUser(repo repository.DataStore) UserService {
	return &userService{userRepository: repo}
}

func (u userService) ResetPassword(resetPassword entity.ResetPassword) error {
	return u.userRepository.ResetPassword(resetPassword)
}

func (u userService) Create(user entity.Register) error {
	return u.userRepository.Create(user)
}

func (u userService) Login(name, email, password, createdAt, updatedAt string, user entity.Login) error {
	return u.userRepository.Login(name, email, password, createdAt, updatedAt, user)
}

func (u userService) CheckUserExist(user entity.Register) bool {
	return u.userRepository.CheckUserExist(user)
}

func (u userService) CheckAndRetrieveUserIDViaEmail(createReset entity.CreateReset) (int, bool) {
	return u.userRepository.CheckAndRetrieveUserIDViaEmail(createReset)
}
