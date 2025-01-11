# Inventory Sales System

## Overview
Inventory Sales System is a RESTful API-based application designed to manage product inventories. It supports operations such as listing, creating, updating, and deleting products, as well as user authentication using JWT.

## Features
- **User Authentication**: Login with a username and password to obtain a JWT token.
- **CRUD Operations on Products**: Create, Read, Update, and Delete product information.
- **Secure API Access**: Protect product routes with JWT authentication.
- **Database Integration**: Uses MySQL with GORM for database management.

## Prerequisites

1. **Install Go**
   - Download and install Go from [golang.org](https://golang.org/dl/).
2. **Install MySQL**
   - Ensure MySQL is installed and running locally.
3. **Install Git**
   - Download and install Git from [git-scm.com](https://git-scm.com/downloads).

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/your-username/InventorySalesSystem.git
cd InventorySalesSystem
```

### Configure the Database

1. Open `config/db.go`.
2. Update the `dsn` string with your MySQL credentials:
   ```go
   dsn := "root:password@tcp(127.0.0.1:3306)/inventory_db?charset=utf8mb4&parseTime=True&loc=Local"
   ```
3. Ensure the `inventory_db` database exists:
   ```sql
   CREATE DATABASE inventory_db;
   ```

### Install Dependencies

Run the following command to install the required Go modules:
```bash
go mod tidy
```

### Run the Application

Start the server using:
```bash
go run main.go
```

### Test API Endpoints

Use Postman or a similar tool to test the API endpoints. Below are the available endpoints:

#### Authentication
- **POST /login**
  - Request Body:
    ```json
    {
        "username": "admin",
        "password": "password123"
    }
    ```
  - Response:
    ```json
    {
        "token": "<JWT Token>"
    }
    ```

#### Products
- **GET /products/**: List all products (requires JWT).
- **GET /products/:id**: Get product details by ID (requires JWT).
- **POST /products/**: Add a new product (requires JWT).
  - Example Body:
    ```json
    {
        "name": "Product Name",
        "price": 100,
        "stock": 50
    }
    ```
- **PUT /products/:id**: Update product details by ID (requires JWT).
- **DELETE /products/:id**: Delete a product by ID (requires JWT).

### JWT Authentication
- Include the token in the `Authorization` header for protected routes:
  ```
  Authorization: Bearer <JWT Token>
  ```

## Project Structure

```
InventorySalesSystem/
├── config/
│   └── db.go
├── controller/
│   ├── auth_controller.go
│   ├── product_controller.go
│   └── register_routes.go
├── entity/
│   ├── product.go
│   └── user.go
├── middleware/
│   └── jwt.go
├── repository/
│   └── product_repository.go
├── router/
│   └── router.go
├── service/
│   └── product_service.go
├── main.go
└── go.mod
```

## Technologies Used
- **Backend**: Go (Gin Framework, GORM)
- **Database**: MySQL
- **Authentication**: JWT

## License
This project is licensed under the MIT License.

## Contributing
Feel free to fork this project and submit pull requests for improvements or bug fixes.

## Author
- **Your Name** - [Your GitHub Profile](https://github.com/your-username)

