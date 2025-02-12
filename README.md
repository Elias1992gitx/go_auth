# JWT Authentication in Go

A secure JWT-based authentication system built with Go, Fiber, and GORM (Object Relational Mapping (ORM) library for Golang).

# Features
-  User registration with password hashing
-  Secure login with JWT token generation
-  Protected routes with JWT verification
-  User information retrieval
-  Secure logout mechanism
-  CORS support for cross-origin requests
-  MySQL database integration

# Prerequisites
Ensure you have the following installed:

- Go (1.23.6 or higher)
- MySQL database
- Git (for cloning the repository)

# Installation
 1 Clone the repository

git clone (https://github.com/Elias1992gitx/go_auth.git)
cd JWT-Authentication-go

 2 Install dependencies

go mod download

 3 Configure your database connection
Modify database/connection.go with your credentials:

dsn := "dsn := "your_username:your_password@tcp(localhost:3306)/your_database""


### 4 Run the application
go run main.go

# Security Features
- Password hashing using bcrypt
- JWT token-based authentication
- HTTP-only cookies for token storage
- CORS protection
- Secure cookie flags

# Project Structure

JWT-Authentication-go/
├── controllers/
│   └── authController.go
├── database/
│   └── connection.go
├── models/
│   └── user.go
├── routes/
│   └── routes.go
├── main.go
├── go.mod
└── go.sum

# Environment Variables
For production deployment, move sensitive configurations to environment variables:

- Database connection string
- JWT secret key
- Server port
- CORS settings

