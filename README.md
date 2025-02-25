Shopline - REST API Framework
This is a modular and scalable REST API framework built using Go , Chi router, GORM ORM, PostgreSQL , Redis , and JWT authentication . It supports features like user registration, product management, comments, ratings, dynamic roles, and caching.

Overview
This application serves as a base framework for building large-scale REST APIs. It includes:

User Management : Registration, login, and role-based access control.
Product Management : CRUD operations with pagination, filtering, and promoted products.
Comment System : Users can submit comments and ratings; admins can publish or reject comments.
Dynamic Roles : Admins can create new roles dynamically.
Caching : Redis is used to cache frequently accessed data (e.g., product lists, published comments).
Scalability : Modular design ensures easy extension and maintenance.

Architecture
The application follows a clean and modular architecture:

Layers:
Handlers : Handle HTTP requests and responses.
Services : Encapsulate business logic.
Repositories : Interact with the database using GORM.
Cache : Centralized caching mechanism using Redis.
Middleware : Handles authentication, authorization, and logging.
Models : Define database schemas and validation rules.
Packages:
cmd: Entry point of the application.
config: Load configuration settings from .env.
internal: Core application logic (handlers, services, repositories, models).
pkg: Utility packages (cache, logger, response, validation, etc.).
docker-compose.yml: Defines the application's services (app, PostgreSQL, Redis).
