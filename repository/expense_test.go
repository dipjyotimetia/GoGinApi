package repository

import (
	"github.com/GoGinApi/v2/entity"
	mockrepository "github.com/GoGinApi/v2/mocks/database"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestName(t *testing.T) {
	userResp1 := entity.User{
		ID:       1,
		Name:     "DemoName",
		Location: "NewLocation",
		Age:      20,
	}

	userResp := []entity.User{userResp1}

	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := mockrepository.NewMockDataStore(ctrl)

	gomock.InOrder(m.EXPECT().GetAllUsers().Return(userResp))
}
