# Project Title

Zocket-Product-Management-App

## Project Description

A backend system using the Go programming
language for a product management application

## Project Structure


Certainly! Here's the README template in Markdown format:

markdown
Copy code
# Project Title

[Your Project Name]

## Project Description

[Provide a brief description of your project.]

## Project Structure

cmd
    -main.go
api
    -routes.go
    -handlers.go
    -handlers_test.go
db
    -database.go
models
    -models.go



## Getting Started

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/atir09/zocket-assignment01.git
   cd zocket-assignment01

2. **Install Dependencies:**

    ```bash
    go mod tidy


3. **Configure MongoDB:**

Make sure you have a MongoDB instance running.
Update the MongoDB connection string in db/database.go.

4. **Run the Application:**

    ```bash
    go run cmd/main.go


## API Endpoints

**Create User:**

Endpoint: POST /users
Payload:{
  "product_name" : string
  "product_description" : string
 "product_images": []string
 "product_price" : int,
 "compressed_product_images": []string
}


**Create Product:**

Endpoint: POST /products


**Get Products:**

Endpoint: GET /products

**Get Product by ID:**

Endpoint: GET /products/{productID}


## Testing

    ```bash
    go test -v ./...