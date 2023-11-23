package main

import (
	"fmt"
	"zocket-api/api"
	"zocket-api/db"
)

func main() {
	fmt.Println("Hello, Product Management Application!")

	// Connect to MongoDB
	db.Connect()

	// Setup API routes
	api.SetupRoutes()
}
