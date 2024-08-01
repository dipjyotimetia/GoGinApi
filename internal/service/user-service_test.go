package service

import (
	"github.com/GoGinApi/v2/internal/entity"
	"github.com/GoGinApi/v2/internal/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/pact-foundation/pact-go/v2/dsl"
)

func TestUserServices(t *testing.T) {
	for _, c := range []TestCase{
		{"Successfully create user", testCreateUser},
		{"Successful reset password", testResetPassword},
		{"Successful test login", testLogin},
		{"Successful check user exist", testCheckUserExist},
		{"Successful check and retrieve userid", testCheckAndRetrieveUserIDViaEmail},
		{"Successful Pact test for user service", testPactUserService},
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

func testPactUserService(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "UserServiceConsumer",
		Provider: "UserServiceProvider",
	}

	defer pact.Teardown()

	pact.AddInteraction().
		Given("User with email test1@gmail.com exists").
		UponReceiving("A request to check user existence").
		WithRequest(dsl.Request{
			Method: "POST",
			Path:   dsl.String("/users/check"),
			Body: dsl.Match(&map[string]interface{}{
				"email": dsl.Like("test1@gmail.com"),
			}),
		}).
		WillRespondWith(dsl.Response{
			Status:  200,
			Headers: dsl.MapMatcher{"Content-Type": dsl.String("application/json")},
			Body: dsl.Match(&map[string]interface{}{
				"exists": dsl.Like(true),
			}),
		})

	err := pact.Verify(func() error {
		// Make request to the provider
		// This is where you would call your actual service
		return nil
	})

	assert.NoError(t, err)
}
