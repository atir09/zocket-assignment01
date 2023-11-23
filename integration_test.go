package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateProductIntegration(t *testing.T) {
	router := gin.Default()

	// Your test code for creating a product via HTTP request
	product := Product{
		ProductID:               1,
		ProductName:             "Test Product",
		ProductDescription:      "Description for testing",
		ProductImages:           []string{"image1.jpg", "image2.jpg"},
		ProductPrice:            19.99,
		CompressedProductImages: nil,
		CreatedAt:               time.Now(),
		UpdatedAt:               time.Now(),
	}

	productJSON, err := json.Marshal(product)
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "/api/products", bytes.NewBuffer(productJSON))
	assert.NoError(t, err)

	req.Header.Set("Content-Type", "application/json")

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
}

func TestGetProductsIntegration(t *testing.T) {

	router := gin.Default()

	// Your test code for retrieving products via HTTP request
	req, err := http.NewRequest("GET", "/api/products", nil)
	assert.NoError(t, err)

	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

}
