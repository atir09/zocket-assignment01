// db_test.go

package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConnectToDB(t *testing.T) {
	assert.NoError(t, ConnectToDB())
}

func TestInsertProduct(t *testing.T) {
	// Your test code for InsertProduct function
	// Example:
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

	err := InsertProduct(product)
	assert.NoError(t, err)
}

func TestGetProductsFromDB(t *testing.T) {
	// Your test code for GetProductsFromDB function
	// Example:
	products, err := GetProductsFromDB()
	assert.NoError(t, err)
	assert.NotEmpty(t, products)
}
