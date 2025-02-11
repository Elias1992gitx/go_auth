JWT Authentication in Go

A secure JWT-based authentication system built with Go, Fiber, and GORM.

Features

User registration with password hashing
Secure login with JWT token generation
Protected routes with JWT verification
User information retrieval
Secure logout mechanism
CORS support for cross-origin requests
MySQL database integration

Prerequisites

Go 1.23.6 or higher
MySQL database
Git (for cloning the repository)

Installation

Clone the repository:
git clone https://github.com/yourusername/JWT-Authentication-go.git
cd JWT-Authentication-go

Install dependencies:
go mod download

Configure your database connection in database/connection.go:
dsn := "your_username:your_password@tcp(localhost:3306)/your_database"

Run the application:
go run main.go

Security Features

Password hashing using bcrypt
JWT token-based authentication
HTTP-only cookies for token storage
CORS protection
Secure cookie flags

Project Structure

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

Environment Variables

For production deployment, consider moving these configurations to environment variables:
Database connection string
JWT secret key
Server port
CORS settings
