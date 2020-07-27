package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
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
	w := performRequest(router, "GET", "/")
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
