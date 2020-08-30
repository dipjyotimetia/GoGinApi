package service

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/repository"
)

type UserService interface {
	InsertUser(user entity.User) int64
	GetAllUsers() []entity.User
	GetUser(id int64) entity.User
	UpdateUser(id int64, user entity.User) int64
	DeleteUser(id int64) int64
}

type userService struct {
	userRepository repository.DataStore
}

func NewUser(repo repository.DataStore) UserService {
	return &userService{userRepository: repo}
}

func (u userService) InsertUser(user entity.User) int64 {
	u.userRepository.InsertUser(user)
	return user.ID
}

func (u userService) GetAllUsers() []entity.User {
	return u.userRepository.GetAllUsers()
}

func (u userService) GetUser(id int64) entity.User {
	return u.userRepository.GetUser(id)
}

func (u userService) UpdateUser(id int64, user entity.User) int64 {
	u.userRepository.UpdateUser(id, user)
	return user.ID
}

func (u userService) DeleteUser(id int64) int64 {
	u.userRepository.DeleteUser(id)
	return id
}
