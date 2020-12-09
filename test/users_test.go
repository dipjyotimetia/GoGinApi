package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/GoGinApi/v2/internal/entity"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

var token string

// TestUserLogin
func TestUserLogin(t *testing.T) {
	req, _ := json.Marshal(entity.Login{
		Password: "password1",
		Email:    "test1@gmail.com",
	})

	res, err := PostRequest("http://localhost:8082/api/login", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, 200, res.StatusCode, "Verified Status")
	token = res.Header.Get("Cookie")
}

// TestGetExpense
func TestGetExpense(t *testing.T) {
	var data entity.Expense
	resp, err := GetRequest("http://localhost:8082/api/v1/getExpense")
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, 200, resp.StatusCode, "Verified Status")
	resBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(resBody, &data)
	assert.Equal(t, 1, data.ExpenseID)
	assert.Equal(t, "Test1", data.ExpenseType)
	assert.Equal(t, 20.5, data.ExpenseAmount)
	assert.Equal(t, 1, data.ClientID)
}
