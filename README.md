# Shopline - REST API Framework FOR ECOMMERCE

This is a modular and scalable REST API framework for ecommerce built using **Go**, **Chi** router, **GORM** ORM, **PostgreSQL**, **Redis**, and **JWT authentication**. It supports features like user registration, product management, comments, ratings, dynamic roles, and caching.

---

## Table of Contents

1. [Overview](#overview)
2. [Features](#features)
3. [Architecture](#architecture)
4. [Prerequisites](#prerequisites)
5. [Setup](#setup)
6. [Usage](#usage)
7. [Environment Variables](#environment-variables)

---

## Overview

This application serves as a base framework for building large-scale REST APIs. It includes:
- **User Management**: Registration, login, and role-based access control.
- **Product Management**: CRUD operations with pagination, filtering, and promoted products.
- **Comment System**: Users can submit comments and ratings; admins can publish or reject comments.
- **Dynamic Roles**: Admins can create, update, or delete roles dynamically.
- **Caching**: Redis is used to cache frequently accessed data (e.g., product lists, published comments).
- **Scalability**: Modular design ensures easy extension and maintenance.

---

## Features

- **Authentication**: JWT-based authentication with role-based access control.
- **Authorization**: Fine-grained permissions for different user roles (`admin`, `seller`, `user`).
- **Pagination**: Efficiently fetch large datasets with pagination.
- **Caching**: Use Redis to cache API responses and query results.
- **Dynamic Roles**: Admins can create new roles dynamically.
- **Logging**: Structured logging with `zap` for better observability.
- **Graceful Shutdown**: Ensures no data loss during shutdowns.
- **Docker Support**: Easily deploy the application using Docker and Docker Compose.

---

## Architecture

The application follows a clean and modular architecture:

### Layers:
1. **Handlers**: Handle HTTP requests and responses.
2. **Services**: Encapsulate business logic.
3. **Repositories**: Interact with the database using GORM.
4. **Cache**: Centralized caching mechanism using Redis.
5. **Middleware**: Handles authentication, authorization, and logging.
6. **Models**: Define database schemas and validation rules.

### Packages:
- `cmd`: Entry point of the application.
- `config`: Load configuration settings from `.env`.
- `internal`: Core application logic (handlers, services, repositories, models).
- `pkg`: Utility packages (cache, logger, response, validation, etc.).
- `docker-compose.yml`: Defines the application's services (app, PostgreSQL, Redis).

---

## Prerequisites

Before running the application, ensure you have the following installed:
- **Go**: Version 1.20 or higher.
- **Docker**: For running the application in containers.
- **Make**: Optional, but useful for running setup scripts.

---

## Setup

### 1. Clone the Repository
```bash
git clone https://github.com/alikhalafnejad/shopline.git
cd shopline
```

### 2. Install Dependencies
Run the following command to install Go dependencies:
```bash
go mod tidy
```

### 3. Configure Environment Variables
Create a `.env` file in the root directory and configure the required variables:
```env
DEBUG=true

# Database Settings
DB_HOST=localhost
DB_PORT=5432
DB_USER=myuser
DB_PASSWORD=mypassword
DB_NAME=mydb

# Redis Settings
REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# JWT Settings
JWT_SECRET_KEY=my_secret_key

# Pagination Defaults
DEFAULT_PAGE=1
DEFAULT_LIMIT=10

# Cache TTL
CACHE_TTL_PRODUCT=5m
CACHE_TTL_COMMENT=10m
```

### 4. Run Migrations
Apply database migrations to create tables:
```bash
go run migrations/migrate.go
```

### 5. Start the Application
You can start the application using Docker Compose:
```bash
docker-compose up --build
```

Alternatively, run it locally:
```bash
go run cmd/main.go
```

The server will start on `http://localhost:8080`.

---

## Usage

### 1. Register a New User
Register a new user by sending a POST request:
```bash
POST /api/v1/register
Content-Type: application/json

{
  "username": "john_doe",
  "email": "john@example.com",
  "password": "securepassword123"
}
```

### 2. Login
Login to get a JWT token:
```bash
POST /api/v1/login
Content-Type: application/json

{
  "email": "john@example.com",
  "password": "securepassword123"
}
```

### 3. Manage Products
- **Get Products**: Fetch paginated products.
  ```bash
  GET /api/v1/products?page=1&limit=10
  Authorization: Bearer <token>
  ```
- **Create Product**: Add a new product (requires seller role).
  ```bash
  POST /api/v1/seller/products
  Authorization: Bearer <seller_token>
  Content-Type: application/json

  {
    "name": "Laptop",
    "price": 999.99,
    "category_id": 1
  }
  ```

### 4. Manage Comments
- **Add Comment**: Submit a comment for a product.
  ```bash
  POST /api/v1/products/1/comments
  Authorization: Bearer <user_token>
  Content-Type: application/json

  {
    "text": "Great product!",
    "rating": 5
  }
  ```
- **Publish Comment**: Approve a comment (requires admin role).
  ```bash
  PUT /api/v1/admin/comments/1/publish
  Authorization: Bearer <admin_token>
  ```
- **Reject Comment**: Reject a comment (requires admin role).
  ```bash
  PUT /api/v1/admin/comments/1/reject
  Authorization: Bearer <admin_token>
  ```

### 5. Dynamic Role Management
Admins can create new roles dynamically:
```bash
POST /api/v1/admin/roles
Authorization: Bearer <admin_token>
Content-Type: application/json

{
  "name": "moderator",
  "description": "Can approve or reject comments."
}
```

---

## Environment Variables

The application uses environment variables for configuration. You can customize these values in the `.env` file.

| Variable            | Default Value      | Description                                      |
|---------------------|--------------------|-------------------------------------------------|
| `DEBUG`            | `true`            | Enable debug logging.                           |
| `DB_HOST`          | `localhost`       | PostgreSQL host address.                        |
| `DB_PORT`          | `5432`           | PostgreSQL port.                                |
| `DB_USER`          | `myuser`         | PostgreSQL username.                            |
| `DB_PASSWORD`      | `mypassword`     | PostgreSQL password.                            |
| `DB_NAME`          | `mydb`           | PostgreSQL database name.                       |
| `REDIS_ADDR`       | `localhost:6379` | Redis address.                                  |
| `REDIS_PASSWORD`   |                   | Redis password (leave empty if none).           |
| `REDIS_DB`         | `0`              | Redis database index.                           |
| `JWT_SECRET_KEY`   | `my_secret_key`  | Secret key for JWT token generation.             |
| `DEFAULT_PAGE`     | `1`              | Default page number for pagination.              |
| `DEFAULT_LIMIT`    | `10`             | Default limit for pagination.                    |
| `CACHE_TTL_PRODUCT` | `5m`          | Cache expiration time for products.              |
| `CACHE_TTL_COMMENT` | `10m`         | Cache expiration time for comments.              |

---

## Dockerization

### 1. Build and Run Containers
Use Docker Compose to build and run the application along with PostgreSQL and Redis:
```bash
docker-compose up --build
```

### 2. Access the Application
Once the containers are running, access the application at:
- **API**: `http://localhost:8080`
- **PostgreSQL**: `postgres://myuser:mypassword@localhost:5432/mydb`
- **Redis**: `redis://localhost:6379`

### 3. Stop Containers
To stop the application:
```bash
docker-compose down
```

---

## Contributing

To contribute to this project:
1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes and push them to your fork.
4. Submit a pull request.

---

## License

This project is licensed under the MIT License. See the [LICENSE](https://opensource.org/license/MIT) file for details.

---

## Contact

If you have any questions or need assistance, feel free to reach out:
- Email: leaderenv@gmail.com
- GitHub: [@alikhalafnejad](https://github.com/alikhalafnejad)

---
