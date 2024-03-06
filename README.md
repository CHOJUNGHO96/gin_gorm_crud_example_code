# gin_gorm_crud_example_code

A simple CRUD implementation gorm example code using gin, the go framework.

# Gin and GORM Go CRUD Example

This project demonstrates a simple CRUD (Create, Read, Update, Delete) application using the Gin web framework and GORM ORM in Go. It's designed to provide a straightforward example for those learning how to implement RESTful APIs with Go, Gin, and GORM with PostgreSQL as the database.

## Getting Started

### Prerequisites

Before you begin, ensure you have met the following requirements:
- Go (version 1.18 or later)
- PostgreSQL

### Installation

Follow these steps to get your development environment set up:

1. Clone the repository to your local machine:
   ```sh
   git clone https://github.com/CHOJUNGHO96/gin_gorm_crud_example_code.git

2. Navigate to the project directory:
   ```sh
   cd gin-gorm-crud-example
   ```

3. Install the required Go dependencies:
   ```sh
   go get .
   ```
   
### Configuring the Database
Create a PostgreSQL database and user, and update the DSN in the init function within main.go with your database details:
```sh
dsn := "host=localhost user=youruser dbname=yourdb password='yourpassword' port=5432 sslmode=disable TimeZone=Asia/Seoul"
```

### Usage
To run the server, execute:
```sh
go run main.go
```
The API endpoints are accessible at http://localhost:8080/ as follows:

GET /users - Retrieve all users </br>
GET /users/:id - Retrieve a single user by ID </br>
POST /users - Create a new user </br>
PUT /users/:id - Update an existing user </br>
DELETE /users/:id - Delete a user </br>
You can test these endpoints using tools like Postman or cURL. Here's an example of creating a user with cURL:
```sh
curl -X POST -H "Content-Type: application/json" -d '{"name":"test", "email":"test@test.com", "reg_date":"2024-01-01"}' http://localhost:8080/users
```
