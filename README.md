# Go-Ecom

Go-Ecom is a Golang-based eCommerce API that handles essential operations for an online store. The API supports product management, order handling, and user authentication. This repository demonstrates how to build scalable and efficient backend services for eCommerce using Golang.

## Features

- Product CRUD operations (Create, Read, Update, Delete)
- User authentication and authorization
- Order management
- Database integration with PostgreSQL
- Error handling and validation

## Tech Stack

- **Backend:** Golang
- **Database:** PostgreSQL
- **API Framework:** net/http
- **Dependency Management:** Go Modules

## Project Structure

```
├── cmd
├── config
├── db
├── services
├── types
├── utils
└── vendor
```

- `cmd`: Application entry points.
- `config`: Configuration settings.
- `db`: Database setup and migrations.
- `services`: Business logic and service layers.
- `types`: Structs and types used across the application.
- `utils`: Utility functions and helpers.
- `vendor`: Third-party libraries.

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/SIRIUS-webkit/go-ecom.git
   ```
2. Navigate to the project directory:
   ```bash
   cd go-ecom
   ```
3. Install dependencies:
   ```bash
   go mod download
   ```

4. Setup PostgreSQL and configure your database connection in the `config` file.

5. Run the application:
   ```bash
   go run cmd/main.go
   ```

## Usage

- **Product Management:** Use RESTful API endpoints to create, update, delete, and list products.
- **Order Processing:** Handle orders with endpoints that create and manage customer orders.
- **User Authentication:** Register and authenticate users with JWT-based security.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request.
