# Go Auth API

A RESTful authentication service built with Golang. This project demonstrates a clean and scalable backend architecture with secure authentication using JSON Web Tokens (JWT). It is suitable for portfolio, learning reference, and real-world backend service development.

## Tech Stack
- Golang
- Gin Web Framework
- GORM ORM
- MySQL
- JSON Web Token (JWT)

## Features
- User registration with secure password hashing
- User login with JWT generation
- JWT-based authentication
- Middleware-protected endpoints
- Clean and maintainable project structure

## Project Structure
cmd/
  server/
internal/
  handler/
  service/
  repository/
  model/
  middleware/
config/

## How to Run
go mod tidy
go run cmd/server/main.go

## Environment Variables
Create a `.env` file based on `.env.example` and configure:
- APP_PORT
- DB_HOST
- DB_USER
- DB_PASSWORD
- DB_NAME
- JWT_SECRET

## API Endpoints
POST   /api/auth/register   Register new user  
POST   /api/auth/login      Authenticate user  
GET    /api/user/me         Get authenticated user  

## Purpose
This project demonstrates backend API development using Golang, implementation of authentication and authorization, and clean project structure following industry practices.

## License
MIT License
