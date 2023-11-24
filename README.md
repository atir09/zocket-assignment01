# Project Title

Zocket-Product-Management-App

## Project Description

A backend system using the Go programming
language for a product management application



## Project Structure

cmd<br>
    -main.go<br>
api<br>
    -routes.go<br>
    -handlers.go<br>
    -handlers_test.go<br>
db<br>
    -database.go<br>
    -db_test.go
models<br>
    -models.go<br>



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

Endpoint: POST /users<br>
Payload:{
    "name":string,
	"mobile":string,
	"latitude":string,
	"longitude" :string
}



**Create Product:**

Endpoint: POST /products <br>
Payload:{
  "product_name" : string
  "product_description" : string
 "product_images": []string
 "product_price" : int,
 "compressed_product_images": []string
}


**Get Products:**

Endpoint: GET /products

**Get Product by ID:**

Endpoint: GET /products/{productID}


## Testing<br>

     go test -v ./...
