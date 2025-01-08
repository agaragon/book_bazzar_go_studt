package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"github.com/gin-gonic/gin"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestCreateBookHandler(t *testing.T) {
	// Initialize the database (mock or in-memory for testing)
	DB = initDB()
	DB.AutoMigrate(&Book{})

	// Set up a test router
	r := gin.Default()
	setupHandlers(r)

	// Create a test request with valid book data
	bookJSON := `{"title": "Test Book", "author": "Test Author"}`
	req, _ := http.NewRequest("POST", "/api/books", bytes.NewBufferString(bookJSON))
	req.Header.Set("Content-Type", "application/json")

	// Record the response
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	// Assert that the response code is 201 Created
	assert.Equal(t, http.StatusCreated, w.Code)

	// Check the response body
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// Assert that the Id is present and is a number
	_, idPresent := response["Id"]
	assert.True(t, idPresent, "Id should be present in the response")

	// Assert other fields
	assert.Equal(t, "Test Book", response["title"])
	assert.Equal(t, "Test Author", response["author"])
}
