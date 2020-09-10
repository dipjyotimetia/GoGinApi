package service

import (
	"github.com/GoGinApi/v2/entity"
	"github.com/GoGinApi/v2/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUserServices(t *testing.T) {
	for _, c := range []TestCase{
		{"Successfully create user", testCreateUser},
		{"Successful reset password", testResetPassword},
		{"Successful test login", testLogin},
		{"Successful check user exist", testCheckUserExist},
		{"Successful check and retrieve userid", testCheckAndRetrieveUserIDViaEmail},
	} {
		t.Run(c.name, c.test)
	}
}

func testCreateUser(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	user := entity.Register{
		Password: "Password1",
		Email:    "test1@gmail.com",
		Name:     "test1",
	}

	mockRepo.On("Create").Return(nil)

	testService := NewUser(mockRepo)
	res := testService.Create(user)
	assert.Equal(t, res, nil)
}

func testResetPassword(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	reset := entity.ResetPassword{
		Password: "Password1",
	}

	mockRepo.On("ResetPassword").Return(nil)

	testService := NewUser(mockRepo)
	res := testService.ResetPassword(reset)
	assert.Equal(t, res, nil)
}

func testLogin(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	login := entity.Login{
		Password: "Password1",
		Email:    "test1@gmail.com",
	}

	mockRepo.On("Login").Return(nil)

	testService := NewUser(mockRepo)
	res := testService.Login("", "", "", "", "", login)
	assert.Equal(t, res, nil)
}

func testCheckUserExist(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	user := entity.Register{
		Password: "Password1",
		Email:    "test1@gmail.com",
		Name:     "test1",
	}

	mockRepo.On("CheckUserExist").Return(true)

	testService := NewUser(mockRepo)
	res := testService.CheckUserExist(user)
	assert.Equal(t, res, true)
}

func testCheckAndRetrieveUserIDViaEmail(t *testing.T) {
	mockRepo := new(mocks.MyMockedObject)

	user := entity.CreateReset{
		Email: "test1@gmail.com",
	}

	mockRepo.On("CheckAndRetrieveUserIDViaEmail").Return(1, true)

	testService := NewUser(mockRepo)
	res, res2 := testService.CheckAndRetrieveUserIDViaEmail(user)
	assert.Equal(t, res, 1)
	assert.Equal(t, res2, true)
}
