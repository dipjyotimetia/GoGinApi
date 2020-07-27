package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	if method == "POST" {
		req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
	}
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPingPong(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"ping": "pong",
	}
	// Grab our router
	router := SetupRouter()
	// Perform a GET request with that handler.
	w := performRequest(router, "GET", "/", nil)
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusOK, w.Code)
	// Convert the JSON response to a map
	var response map[string]string
	var err = json.Unmarshal([]byte(w.Body.String()), &response) // Grab the value & whether or not it exists
	value, exists := response["ping"]
	// Make some assertions on the correctness of the response.
	assert.Nil(t, err)
	assert.True(t, exists)
	assert.Equal(t, body["ping"], value)
}

func TestVideoPost(t *testing.T) {

	var jsonStr = []byte(`{
    "id":1,
    "title": "xyz",
    "description": "iwfhi iwehi",
    "url": "http://localhost:8080/test",
    "Author": {
        "id":1,
        "firstName": "test",
        "LastName": "new",
        "Age": 20,
        "Email": "testnew@gmail.com"
    },
    "PersonID": 1
}`)

	router := SetupRouter()

	w := performRequest(router, "POST", "/videos", jsonStr)
	assert.Equal(t, http.StatusOK, w.Code)
}
