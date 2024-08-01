package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"github.com/pact-foundation/pact-go/v2/dsl"
)

func TestHashPassword(t *testing.T) {
	tests := []Register{
		{Name: "testtest1", Password: "test@password1", Email: "test@gmail.com"},
		{Name: "testtest2", Password: "test@password2", Email: "test1@gmail.com"},
		{Name: "testtest3", Password: "test@password3", Email: "test2@gmail.com"},
	}

	for _, tt := range tests {
		t.Run("TestHash", func(t *testing.T) {
			pass := tt.Password
			HashPassword(&tt)
			assert.NotEqual(t, pass, tt.Password)
		})
	}

	for _, tt := range tests {
		t.Run("TestCreateHash", func(t *testing.T) {
			pass := CreateHashedPassword(tt.Password)
			assert.NotEqual(t, pass, tt.Password)
		})
	}

	for _, tt := range tests {
		t.Run("TestCheckPasswordHash", func(t *testing.T) {
			pass := tt.Password
			HashPassword(&tt)
			assert.NoError(t, CheckPasswordHash(pass, tt.Password))
			assert.Error(t, CheckPasswordHash("test@password", tt.Password))
		})
	}
}

func TestPact_UserEntity(t *testing.T) {
	pact := &dsl.Pact{
		Consumer: "UserEntityConsumer",
		Provider: "UserEntityProvider",
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
