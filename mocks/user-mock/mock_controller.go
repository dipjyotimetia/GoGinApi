// Code generated by MockGen. DO NOT EDIT.
// Source: controller/user-controller.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	entity "github.com/GoGinApi/v2/entity"
	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserController is a mock of UserController interface
type MockUserController struct {
	ctrl     *gomock.Controller
	recorder *MockUserControllerMockRecorder
}

// MockUserControllerMockRecorder is the mock recorder for MockUserController
type MockUserControllerMockRecorder struct {
	mock *MockUserController
}

// NewMockUserController creates a new mock instance
func NewMockUserController(ctrl *gomock.Controller) *MockUserController {
	mock := &MockUserController{ctrl: ctrl}
	mock.recorder = &MockUserControllerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserController) EXPECT() *MockUserControllerMockRecorder {
	return m.recorder
}

// InsertUser mocks base method
func (m *MockUserController) InsertUser(ctx *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertUser", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertUser indicates an expected call of InsertUser
func (mr *MockUserControllerMockRecorder) InsertUser(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertUser", reflect.TypeOf((*MockUserController)(nil).InsertUser), ctx)
}

// GetAllUsers mocks base method
func (m *MockUserController) GetAllUsers() []entity.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]entity.User)
	return ret0
}

// GetAllUsers indicates an expected call of GetAllUsers
func (mr *MockUserControllerMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockUserController)(nil).GetAllUsers))
}

// GetUser mocks base method
func (m *MockUserController) GetUser(ctx *gin.Context) entity.User {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", ctx)
	ret0, _ := ret[0].(entity.User)
	return ret0
}

// GetUser indicates an expected call of GetUser
func (mr *MockUserControllerMockRecorder) GetUser(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockUserController)(nil).GetUser), ctx)
}

// UpdateUser mocks base method
func (m *MockUserController) UpdateUser(ctx *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUserControllerMockRecorder) UpdateUser(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserController)(nil).UpdateUser), ctx)
}

// DeleteUser mocks base method
func (m *MockUserController) DeleteUser(ctx *gin.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", ctx)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockUserControllerMockRecorder) DeleteUser(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserController)(nil).DeleteUser), ctx)
}
