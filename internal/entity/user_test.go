package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
