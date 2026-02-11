# Go-Commerce Backend

A comprehensive e-commerce REST API built with Go (Golang) as a structured learning project. This repository demonstrates a progressive understanding of Go fundamentals through advanced concepts, culminating in a production-ready backend system with database integration, RESTful architecture, and clean code principles.

## Project Overview

This project was developed as a systematic learning journey through Go programming, following a semester-based curriculum that covers essential concepts from basic syntax to advanced web development patterns. The final application is a fully functional e-commerce backend API with user management, product inventory, and order processing capabilities.

## Learning Objectives & Mastered Concepts

### Semester 1: Fundamentals
**Weeks 1-8: Core Language Features**

- **Variables & Data Types**: Deep understanding of Go's type system, including explicit and implicit declarations, type conversion, and zero values
- **Operators & Control Flow**: Mastery of arithmetic, comparison, and logical operators; implementation of conditional logic with if-else and switch statements
- **Loops & Arrays**: Proficiency in Go's for loop variations, array manipulation, and iteration patterns
- **Functions**: Experience with multiple return values, variadic parameters, named returns, and function closures

**Key Learning Files**:
- `learning Path/learn_variables.go` - Variable declarations, constants, type conversion
- `learning Path/learn_operators.go` - Operators, conditionals, switch statements
- `learning Path/learn_loops.go` - Loop patterns, array operations

### Semester 2: Intermediate Concepts
**Weeks 9-14: Data Structures & Object-Oriented Design**

- **Slices & Maps**: Dynamic data structures, slice operations (append, copy, slicing), map manipulation and key-value storage patterns
- **Structs & Methods**: Custom type definitions, method receivers (value vs pointer), struct composition
- **Pointers**: Memory management, pointer dereferencing, understanding when to use pointers for efficiency and mutability
- **Interfaces**: Polymorphism in Go, interface satisfaction, type assertions, and the power of implicit implementation
- **Error Handling**: Go's explicit error handling pattern, custom error types, error wrapping, panic and recover mechanisms

**Key Learning Files**:
- `learning Path/learn_slices_maps.go` - Dynamic collections, CRUD operations on data structures
- `learning Path/learn_functions.go` - Struct methods, advanced function patterns
- `learning Path/learn_advanced.go` - Interfaces, pointers, custom errors, panic/recover

### Semester 3: Advanced & Production-Ready Concepts
**Weeks 15-20: Database Integration & Web Development**

- **File I/O Operations**: Reading and writing files, file handling patterns, resource management with defer
- **JSON Serialization**: Struct tags, marshaling/unmarshaling, working with JSON data
- **CSV Processing**: Reading and writing CSV files, data import/export functionality
- **Database Integration**: MySQL connectivity, SQL query execution, prepared statements, connection pooling
- **REST API Development**: HTTP server implementation, routing with Gorilla Mux, request/response handling
- **Middleware Pattern**: Logger middleware, CORS handling, request preprocessing
- **MVC Architecture**: Separation of concerns with Models, Controllers, and Routes
- **API Design**: RESTful conventions, proper HTTP status codes, JSON response standardization

**Key Learning Files**:
- `learning Path/learn_file_json_db.go` - File operations, JSON/CSV handling, database CRUD

## Technology Stack

- **Language**: Go 1.24
- **Database**: MySQL (via XAMPP)
- **Router**: Gorilla Mux
- **Database Driver**: go-sql-driver/mysql
- **Architecture**: MVC (Model-View-Controller)

## Project Structure

```
go-commerce-backend/
├── main.go                      # Application entry point
├── config/
│   └── database.go              # Database connection & configuration
├── models/
│   ├── user.go                  # User data model
│   ├── product.go               # Product data model
│   └── order.go                 # Order data model
├── controllers/
│   ├── user_controller.go       # User business logic & handlers
│   └── product_controller.go    # Product business logic & handlers
├── routes/
│   └── routes.go                # API route definitions
├── middlewares/
│   └── middleware.go            # Logger, CORS, and request processing
├── utils/
│   └── response.go              # Standardized JSON response helpers
└── learning Path/
    ├── learn_variables.go       # Variables & data types tutorial
    ├── learn_operators.go       # Operators & control flow tutorial
    ├── learn_loops.go           # Loops & arrays tutorial
    ├── learn_slices_maps.go     # Slices & maps tutorial
    ├── learn_functions.go       # Functions & structs tutorial
    ├── learn_advanced.go        # Pointers, interfaces, errors tutorial
    └── learn_file_json_db.go    # File I/O, JSON, database tutorial
```

## API Endpoints

### Health Check
- `GET /api/health` - Server status check

### User Management
- `GET /api/users` - Retrieve all users
- `GET /api/users/{id}` - Retrieve user by ID
- `POST /api/users` - Create new user
- `PUT /api/users/{id}` - Update user information
- `DELETE /api/users/{id}` - Delete user

### Product Management
- `GET /api/products` - Retrieve all products
- `GET /api/products/{id}` - Retrieve product by ID
- `GET /api/products/search?q={keyword}` - Search products by name or category
- `POST /api/products` - Create new product
- `PUT /api/products/{id}` - Update product information
- `DELETE /api/products/{id}` - Delete product

## Database Schema

### Users Table
```sql
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    balance INT DEFAULT 0,
    is_member BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Products Table
```sql
CREATE TABLE products (
    id VARCHAR(50) PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    price INT NOT NULL,
    stock INT DEFAULT 0,
    category VARCHAR(100),
    rating DECIMAL(3,2) DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Orders Table
```sql
CREATE TABLE orders (
    id VARCHAR(50) PRIMARY KEY,
    customer_id INT NOT NULL,
    total INT NOT NULL,
    status VARCHAR(50) DEFAULT 'pending',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES users(id)
);
```

### Order Items Table
```sql
CREATE TABLE order_items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    order_id VARCHAR(50) NOT NULL,
    product_id VARCHAR(50) NOT NULL,
    quantity INT NOT NULL,
    price INT NOT NULL,
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
```

## Installation & Setup

### Prerequisites
- Go 1.21 or higher
- MySQL (XAMPP recommended)
- Git

### Installation Steps

1. **Clone the repository**
```bash
git clone https://github.com/HHHAAAANNNNN/go-commerce-backend.git
cd go-commerce-backend
```

2. **Install dependencies**
```bash
go mod download
go mod tidy
```

3. **Configure database**
   - Start XAMPP and ensure MySQL is running
   - Create database:
```sql
CREATE DATABASE go_commerce;
```
   - Execute the schema SQL scripts (see Database Schema section)

4. **Update database credentials** (if necessary)
   - Edit `config/database.go` and update the DSN string:
```go
dsn := "root:@tcp(localhost:3306)/go_commerce?parseTime=true"
```

5. **Run the application**
```bash
go run main.go
```

The server will start at `http://localhost:8080`

## API Usage Examples

### Get all users
```bash
curl http://localhost:8080/api/users
```

### Create a new user
```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "John Doe",
    "email": "john@example.com",
    "balance": 5000000,
    "is_member": true
  }'
```

### Search products
```bash
curl "http://localhost:8080/api/products/search?q=laptop"
```

### Get product by ID
```bash
curl http://localhost:8080/api/products/LAP001
```

## Key Features Implemented

### 1. RESTful API Design
- Proper HTTP methods (GET, POST, PUT, DELETE)
- Meaningful endpoint naming conventions
- Appropriate HTTP status codes (200, 201, 400, 404, 500)
- JSON request/response format

### 2. Database Integration
- Connection pooling for efficiency
- Prepared statements to prevent SQL injection
- Transaction support for data integrity
- Proper error handling for database operations

### 3. Middleware Chain
- Request logging for debugging and monitoring
- CORS support for cross-origin requests
- Extensible middleware pattern for future additions (authentication, rate limiting, etc.)

### 4. Clean Architecture
- Separation of concerns (MVC pattern)
- Reusable components (utils, models)
- Modular design for easy maintenance and testing
- Consistent error handling across the application

### 5. Data Management
- CRUD operations for all entities
- Search and filter capabilities
- Data validation at controller level
- Standardized JSON responses

## Learning Progression

This project represents a structured learning path through Go development:

1. **Foundation Phase**: Understanding Go syntax, types, and control structures
2. **Intermediate Phase**: Working with complex data structures, functions, and OOP concepts
3. **Advanced Phase**: Implementing real-world patterns like interfaces, error handling, and pointer usage
4. **Production Phase**: Building a complete REST API with database integration and clean architecture

Each phase is documented in the `learning Path/` directory with standalone executable examples demonstrating specific concepts.

## Technical Highlights

- **Type Safety**: Leveraging Go's static typing for compile-time error detection
- **Explicit Error Handling**: Following Go idioms for robust error management
- **Interface-Based Design**: Using interfaces for flexible and testable code
- **Pointer Efficiency**: Understanding when to use pointers for performance and mutability
- **Goroutines Ready**: Architecture designed to easily incorporate concurrent processing
- **Database Best Practices**: Connection management, prepared statements, and proper resource cleanup

## Future Enhancements

- JWT authentication and authorization
- Unit and integration testing
- Rate limiting middleware
- Pagination for list endpoints
- Logging framework integration
- Docker containerization
- CI/CD pipeline setup
- API documentation with Swagger

## Development Environment

- **IDE**: Visual Studio Code
- **Go Extensions**: Go extension for VS Code
- **Database Tool**: phpMyAdmin (XAMPP)
- **Version Control**: Git
- **Testing Tools**: cURL, Postman

## Acknowledgments

This project was developed as a comprehensive learning exercise to master Go programming through practical application development. The structured approach from fundamentals to advanced concepts demonstrates a methodical understanding of the language and its ecosystem.

## License

This project is for educational purposes.

## Contact

GitHub: [@HHHAAAANNNNN](https://github.com/HHHAAAANNNNN)
