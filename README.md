# Ecom GO API

A high-performance e-commerce API built with Go, featuring a modular architecture, type-safe database queries, and containerized deployment.

## 🚀 Features

- RESTful API for e-commerce operations
- PostgreSQL database with type-safe queries
- Docker containerization for easy deployment
- Database migrations with Goose
- Structured logging with slog
- Modular project structure
- Environment-based configuration

## 🛠 Tech Stack

### Core Framework
- **Go 1.25.5** - Programming language
- **Chi Router** - Lightweight HTTP routing library for building REST APIs
  - `github.com/go-chi/chi v1.5.5`
  - `github.com/go-chi/chi/v5 v5.2.3`

### Database
- **PostgreSQL 16 (Alpine)** - Relational database
- **pgx v5** - PostgreSQL driver for Go with native support for pgx types
  - `github.com/jackc/pgx/v5 v5.8.0`
- **sqlc** - Type-safe SQL code generator
  - Generates Go code from SQL queries
  - Eliminates SQL injection vulnerabilities
  - Provides compile-time type checking

### Database Migrations
- **Goose** - Database migration management tool
- Supports versioned migrations stored in `internal/adapters/postgresql/migrations/`

### Infrastructure
- **Docker & Docker Compose** - Containerization and orchestration
- PostgreSQL running in Alpine Linux container
- Persistent volumes for data storage

### Logging
- **slog** - Structured logging (Go 1.21+)
- Text-based output to stdout

## 📁 Project Structure

```
Ecom_go_api/
├── cmd/
│   ├── main.go          # Application entry point
│   └── api.go           # API configuration and handlers
├── internal/
│   ├── adapter/
│   │   ├── env/
│   │   │   └── env.go   # Environment variable configuration
│   │   └── postgresql/
│   │       └── sqlc/    # Generated sqlc code
│   ├── adapters/
│   │   └── postgresql/
│   │       ├── migrations/
│   │       │   └── 00001_add_products.sql
│   │       └── sqlc/
│   │           └── queries.sql
│   ├── json/
│   │   └── json.go      # JSON utilities
│   └── products/
│       ├── handler.go   # HTTP handlers for products
│       └── service.go   # Business logic for products
├── docker-compose.yaml  # Docker services configuration
├── go.mod              # Go module definition
├── go.sum              # Go dependencies checksum
├── sqlc.yaml           # sqlc configuration
└── README.md           # This file
```

## 🔧 Setup & Installation

### Prerequisites
- Go 1.25.5 or higher
- Docker & Docker Compose
- PostgreSQL client tools (optional, for direct DB access)

### Environment Setup

1. **Clone the repository**
   ```bash
   git clone https://github.com/Abhishekkapoor98/Ecom_GO_API.git
   cd Ecom_GO_API
   ```

2. **Configure environment variables**
   
   Create a `.env` file in the root directory:
   ```env
   GOOSE_DBSTRING=host=localhost user=postgres password=postgres dbname=ecom sslmode=disable
   ```

3. **Start PostgreSQL with Docker Compose**
   ```bash
   docker compose up
   ```

4. **Run database migrations**
   ```bash
   goose -dir ./internal/adapters/postgresql/migrations up
   ```

5. **Run the application**
   ```bash
   go run cmd/*.go
   ```

The API will be available at `http://localhost:8080`

## 🗄️ Database Configuration

- **Host:** localhost
- **Port:** 5432
- **Database:** ecom
- **Username:** postgres
- **Password:** postgres (change in production)
- **SSL Mode:** disable (for development)

Connection string format:
```
host=localhost user=postgres password=postgres dbname=ecom sslmode=disable
```

## 📦 Available Endpoints

### Products
- `GET /api/products` - List all products

TBD: Future scope:
- `POST /api/products` - Create a new product
- `GET /api/products/{id}` - Get product by ID
- `PUT /api/products/{id}` - Update a product
- `DELETE /api/products/{id}` - Delete a product

## 🔄 Development Workflow

### Code Generation
When you modify SQL queries, regenerate Go code:
```bash
sqlc generate
```

### Database Migrations
Create new migrations:
```bash
goose -dir ./internal/adapters/postgresql/migrations create migration_name sql
```

Run migrations:
```bash
goose -dir ./internal/adapters/postgresql/migrations up
```

Rollback migrations:
```bash
goose -dir ./internal/adapters/postgresql/migrations down
```

### Building Docker Image
```bash
docker build -t ecom-go-api:latest .
```

## 📝 Configuration

The application uses environment variables for configuration:

| Variable | Default | Description |
|----------|---------|-------------|
| `GOOSE_DBSTRING` | `host=localhost user=postgres password=postgres dbname=ecom sslmode=disable` | PostgreSQL connection string |

## 🧪 Testing

```bash
go run cmd/*.go
using the postman for calling the API.
```

## 👤 Author

[Abhishek Kapoor](https://github.com/Abhishekkapoor98)

## 🤝 Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
  
