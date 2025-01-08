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
	DB = initDB()
	DB.AutoMigrate(&Book{})

	r := gin.Default()
	setupHandlers(r)

	bookJSON := `{"title":"Test Book", "author":"Test Author"}`
	req, _ := http.NewRequest("POST", "/api/books", bytes.NewBufferString(bookJSON))
	req.Header.Set("Content-Type",  "application/json")	

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}

	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	_, idPresent := response["id"]
	assert.True(t, idPresent, "Id should be present in the response")
	assert.Equal(t, "Test Book", response["title"])
	assert.Equal(t, "Test Author", response["author"])
}