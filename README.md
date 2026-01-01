# Go Auth API

RESTful authentication service built with Golang, implementing secure JWT-based authentication and clean architecture principles.

## Tech Stack
- Golang
- Gin (HTTP Framework)
- GORM (ORM)
- MySQL
- JWT (JSON Web Token)
- bcrypt (Password Hashing)

## Features
- User registration
- User login
- Password hashing with bcrypt
- JWT authentication
- Protected routes with middleware
- Clean architecture (handler, service, repository)

## Project Structure
```text
cmd/
 └─ server/
    └─ main.go
config/
internal/
 ├─ handler/
 ├─ service/
 ├─ repository/
 ├─ middleware/
 └─ model/
```

## Environment Variables
Create `.env` file:

```env
APP_PORT=8080
DB_HOST=localhost
DB_USER=root
DB_PASS=
DB_NAME=go_auth
JWT_SECRET=supersecretkey
```

## How to Run
```bash
go mod tidy
go run cmd/server/main.go
```

Server will run on:
```
http://localhost:8080
```

## Authentication Flow

### Register
**POST**
```
/api/auth/register
```

Body:
```json
{
  "name": "Zums",
  "email": "zums@mail.com",
  "password": "password123"
}
```

### Login
**POST**
```
/api/auth/login
```

Body:
```json
{
  "email": "zums@mail.com",
  "password": "password123"
}
```

Response:
```json
{
  "access_token": "JWT_TOKEN_HERE"
}
```

### Protected Route
**GET**
```
/api/user/me
```

Header:
```
Authorization: Bearer <JWT_TOKEN>
```

Response:
```json
{
  "user_id": 1
}
```

## Notes
- Passwords are securely hashed using bcrypt
- JWT is used for stateless authentication
- Middleware protects private endpoints
- Suitable for backend portfolio and real-world reference

## Author
Zums  
Backend / Go Developer
