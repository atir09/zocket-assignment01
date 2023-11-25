# Project Title

Zocket-Product-Management-App

## Project Description

Zocket-Product-Management-App is a backend application built using the Go programming language. The system is designed to manage product information and user data, providing a simple and efficient solution. The application includes API endpoints for creating users, adding products, and retrieving product details. User and product information is stored in a MongoDB database, ensuring data persistence.

## Key Features

**User Management:**
- Create users with essential details such as name, contact number, and location coordinates.

**Product Management:**
- Add products with details like name, description, images, and price.

**Data Storage:**
- Utilize a MongoDB database to store and manage user and product information.

**Testing:**
- The project includes comprehensive unit tests to ensure code quality. The README provides commands for executing tests.

## Project Structure

- cmd<br>
  - main.go<br>
- api<br>
  - routes.go<br>
  - handlers.go<br>
  - handlers_test.go<br>
- db<br>
  - database.go<br>
  - db_test.go
- models<br>
  - models.go<br>



## Getting Started

- Follow the instructions to set up the project locally, configure dependencies, and run the application.

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/atir09/zocket-assignment01.git
   cd zocket-assignment01

2. **Install Dependencies:**

    ```bash
    go mod tidy


3. **Configure MongoDB:**

   Make sure you have a MongoDB instance running. Update the MongoDB connection string in db/database.go.

4. **Run the Application:**

    ```bash
    go run cmd/main.go


## API Endpoints

**Create User:**

- Endpoint : POST /users<br>
- Payload : {
    "name":string, 
	"mobile":string, 
	"latitude":string, 
	"longitude" :string
}



**Create Product:**

- Endpoint : POST /products <br>
- Payload : {
  "product_name" : string, 
  "product_description" : string, 
 "product_images": []string, 
 "product_price" : int, 
}


**Get Products:**

- Endpoint : GET /products

**Get Product by ID:**

- Endpoint : GET /products/{productID}


## Testing<br>

     go test -v ./...
