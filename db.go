package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
)

// ConnectToDB establishes a connection to the MongoDB database.
func ConnectToDB() error {
	// Set your MongoDB connection string
	uri := "mongodb+srv://atir:atir@cluster0.jyf1i6w.mongodb.net/?retryWrites=true&w=majority"

	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	client, err = mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println("Connected to MongoDB!")

	database = client.Database("zocket01")
	collection = database.Collection("products")

	return nil
}

// InsertProduct inserts a product into the MongoDB collection.
func InsertProduct(product Product) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := collection.InsertOne(ctx, product)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

// GetProducts retrieves all products from the MongoDB collection.
func GetProductsFromDB() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []Product
	if err := cursor.All(ctx, &products); err != nil {
		log.Println(err)
		return nil, err
	}

	return products, nil
}
