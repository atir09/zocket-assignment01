// main.go

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// User represents a user in the system.
type User struct {
	ID        int       `json:"id" bson:"_id"`
	Name      string    `json:"name" bson:"name"`
	Mobile    string    `json:"mobile" bson:"mobile"`
	Latitude  float64   `json:"latitude" bson:"latitude"`
	Longitude float64   `json:"longitude" bson:"longitude"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

// Product represents a product in the system.
type Product struct {
	ProductID               int       `json:"product_id" bson:"_id"`
	ProductName             string    `json:"product_name" bson:"product_name"`
	ProductDescription      string    `json:"product_description" bson:"product_description"`
	ProductImages           []string  `json:"product_images" bson:"product_images"`
	ProductPrice            float64   `json:"product_price" bson:"product_price"`
	CompressedProductImages []string  `json:"compressed_product_images" bson:"compressed_product_images"`
	CreatedAt               time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt               time.Time `json:"updated_at" bson:"updated_at"`
}

func main() {
	// Initialize the MongoDB connection
	if err := ConnectToDB(); err != nil {
		fmt.Println("Failed to connect to MongoDB.")
		return
	}

	// Set up Gin router
	router := gin.Default()

	// API routes
	router.POST("/api/products", CreateProduct)
	router.GET("/api/products", GetProducts)

	// Run the server
	router.Run(":8080")
}

// CreateProduct handles the creation of a new product.
func CreateProduct(c *gin.Context) {
	var product Product
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Add timestamp values
	now := time.Now()
	product.CreatedAt = now
	product.UpdatedAt = now

	// Insert the product into the database
	if err := InsertProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusCreated, product)
}

// GetProducts retrieves all products from the database.
func GetProducts(c *gin.Context) {
	products, err := GetProductsFromDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}

	c.JSON(http.StatusOK, products)
}
