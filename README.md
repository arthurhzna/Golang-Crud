# Go REST API

A RESTful API implementation using Go (Golang) with PostgreSQL as the database. This project implements clean architecture principles and includes authentication middleware.

## 🚀 Features

- RESTful API endpoints for Category management
- PostgreSQL database integration
- Authentication middleware
- Request validation using go-playground/validator
- Clean architecture implementation
- JSON response formatting
- Error handling middleware

## 📋 Prerequisites

- Go 1.x
- PostgreSQL
- Git

## 🛠️ Tech Stack

- [Go](https://golang.org/) - Programming Language
- [pgx](https://github.com/jackc/pgx) - PostgreSQL Driver
- [go-playground/validator](https://github.com/go-playground/validator) - Request Validation

## 📁 Project Structure
```plaintext
<main golder>/
├── app/
│   ├── database.go        # Database configuration
│   └── router.go          # Router setup
├── controller/            # HTTP request handlers
├── exception/             # Error handling
├── helper/                # Utility functions
├── middleware/            # HTTP middleware
├── model/                 # Data models
│   ├── domain/           # Domain models
│   └── web/              # Request/Response models
├── repository/            # Data access layer
├── service/               # Business logic layer
└── main.go               # Application entry point
```


## 🚀 Getting Started

1. Clone the repository:
```bash
git clone <your-repository-url>
```

2. Install dependencies:
```bash
go mod download
```

3. Set up your PostgreSQL database

4. Run the application:
```bash
go run main.go
```

The server will start on `localhost:3000`

## 🔒 Authentication

The API is protected with authentication middleware. Make sure to include the appropriate authentication headers when making requests.

## 📝 API Endpoints

### Category Management

- `GET /api/categories` - Get all categories
- `GET /api/categories/:id` - Get category by ID
- `POST /api/categories` - Create new category
- `PUT /api/categories/:id` - Update category
- `DELETE /api/categories/:id` - Delete category

## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check [issues page](your-issues-url).

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
