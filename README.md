# Go Test Application

A full-stack application built with Go backend and Vue 3 frontend, demonstrating modern web development practices with comprehensive testing and Docker deployment.

## 🎯 Problem Statement

This project implements a CRUD API for managing items with debtor and beneficiary information. The goal is to demonstrate Go best practices including thread-safe in-memory storage, custom validation with user-friendly error messages, and comprehensive testing. The architecture applies Domain-Driven Design (DDD) principles, separating business logic from infrastructure concerns while leveraging Go's concurrency primitives for thread-safe operations.

**Key Challenges Addressed:**
- Thread-safe concurrent access using Go's `sync.RWMutex`
- User-friendly validation errors instead of technical validation messages
- Clean separation between HTTP handlers and business logic
- Comprehensive testing of edge cases (empty GUIDs, enum case insensitivity, negative amounts)

## 👤 Developer Notes

**Developer**: Olanrewaju Sule-Balogun

**Technology Experience**:
- **Vue.js**: Experienced with Vue 3 Composition API, Pinia, TypeScript, and modern frontend development
- **Go**: Relatively new to Go, but an experienced full-stack developer applying Go idioms and the Gin framework to build robust, tested, Dockerized applications
    - Leveraging Laravel/PHP experience (validation, DDD, repository pattern) in Go implementations
    - Utilizing Go concurrency primitives (`sync.RWMutex`) for thread-safe operations
    - Building RESTful APIs with Gin framework, custom validation, and comprehensive error handling

## 🏗️ Architecture Overview

This application follows a **Domain-Driven Design (DDD)** approach with clear separation of concerns:

- **Backend**: Go with Gin framework, custom validation, and in-memory storage
- **Frontend**: Vue 3 with Composition API, Pinia state management, Tailwind CSS and TypeScript
- **Deployment**: Docker Compose for containerized deployment

## 📁 Project Structure

```
go-test/
├── backend/                   # Go backend application
│   ├── bootstrap/             # Application initialization
│   │   └── validators.go      # Custom validator registration
│   ├── domain/                # Domain layer (DDD approach)
│   │   ├── dto/               # Data Transfer Objects
│   │   │   ├── item_create_dto.go
│   │   │   └── item_update_dto.go
│   │   ├── models/            # Domain entities
│   │   │   ├── item.go
│   │   │   ├── account.go
│   │   │   └── party.go
│   │   ├── enums/             # Domain enumerations
│   │   │   └── item_enum.go
│   │   ├── validators/        # Custom validation logic
│   │   │   └── validators.go
│   │   └── repository/        # Data access layer
│   │       └── items_repository.go
│   ├── handlers/              # HTTP request handlers
│   │   └── items_handler.go
│   ├── helpers/               # Utility functions
│   │   ├── response.go        # HTTP response helpers
│   │   ├── utils.go          # General utilities
│   │   └── validation.go     # Custom validation error formatting
│   └── tests/                 # Test suites
│       ├── setup.go           # Test setup and configuration
│       └── feature/           # Feature tests
│           ├── create_item_test.go
│           ├── update_item_test.go
│           ├── delete_item_test.go
│           ├── list_item_test.go
│           └── list_item_by_guid_test.go
├── frontend/                  # Vue 3 frontend application
│   ├── src/
│   │   ├── components/        # Reusable Vue components
│   │   │   ├── ItemsTable.vue
│   │   │   ├── ItemModal.vue
│   │   │   ├── ItemSearch.vue
│   │   │   └── ConfirmationModal.vue
│   │   ├── stores/            # Pinia state management
│   │   │   ├── items.ts
│   │   │   └── __tests__/
│   │   │       └── items.test.ts # Unit tests for CRUD operations
│   │   ├── types/             # TypeScript type definitions
│   │   │   ├── entities.ts
│   │   │   ├── dto.ts
│   │   │   ├── enums.ts
│   │   │   └── index.ts
│   │   ├── utils/             # Frontend utilities
│   │   │   └── formatters.ts
│   │   ├── router/            # Vue Router configuration
│   │   │   └── index.ts
│   │   ├── App.vue            # Root component
│   │   └── main.ts            # Application entry point
│   ├── Dockerfile             # Frontend container configuration
├── docker-compose.yml         # Multi-container deployment
├── Dockerfile                 # Backend container configuration
└── main.go                    # Backend application entry point
```

## 🎯 Design Decisions & Technical Choices

### Backend Architecture

#### **Domain-Driven Design (DDD)**
- **Rationale**: Clear separation of business logic from infrastructure concerns
- **Architecture Flow**:
```
HTTP Request
    ↓
Handler (HTTP concerns, validation)
    ↓
DTO (Data Transfer Objects)
    ↓
Helpers (Domain transformations, utilities)
    ↓
Repository (Data access layer)
    ↓
In-Memory Storage (ItemsStore with mutex)
```
- **Implementation**: 
  - `domain/models/` for business entities
  - `domain/dto/` for data transfer objects
  - `domain/enums/` for domain enumerations
  - `domain/validators/` for business rules
  - `repository/` for data access abstraction

#### **Custom Validation System**
- **Rationale**: Provide user-friendly error messages instead of technical validation errors
- **Implementation**:
  - Custom validators for `itemtype`, `itemstatus`, and `sortcode`
  - Structured error responses: `{"errors": {"field": "message"}}`
  - JSON parsing error handling for type mismatches

#### **Repository Pattern**
- **Rationale**: Centralized data access with thread-safe operations
- **Implementation**: `ItemsStore` struct with mutex for concurrent access using in-memory storage in `items_repository.go`

#### **Handler Layer**
- **Rationale**: Clean separation between HTTP concerns and business logic
- **Implementation**: `ItemsHandler` in `items_handler.go` for HTTP request/response handling

#### **Helper Utilities**
- **Rationale**: Reusable utilities for HTTP responses, validation, and domain transformations
- **Implementation**:
  - `helpers/response.go`: `Respond()`, `Error()`, `NoContent()` for consistent HTTP responses with proper status codes (200, 201, 204, 400, 404, 500)
  - `helpers/validation.go`: `ValidationErrorResponse()` with structured, user-friendly error messages
  - `helpers/utils.go`: Business logic utilities (`NewItemFromDTO()`, `ApplyUpdate()`, `ParseLimit()`) for domain transformations
  - Response formats: single objects for create/update operations, arrays for list operations

### Frontend Architecture

#### **Vue 3 Composition API**
- **Rationale**: Better TypeScript support, improved code organization, and reusability
- **Implementation**: `<script setup>` syntax throughout all components

#### **Pinia State Management**
- **Rationale**: Official Vue 3 state management with better TypeScript support than Vuex
- **Implementation**: 
  - Centralized store for items data
  - Reactive state with computed properties

#### **Component-Based Architecture**
- **Rationale**: Reusable, maintainable, and testable components
- **Implementation**:
  - `ItemsTable.vue`: Main data display with collapsible details
  - `ItemModal.vue`: Reusable create/edit form
  - `ItemSearch.vue`: Search functionality component
  - `ConfirmationModal.vue`: Reusable confirmation dialog

#### **TypeScript Integration**
- **Rationale**: Type safety, better IDE support, and reduced runtime errors
- **Implementation**: Comprehensive type definitions in `types/` folder

### Testing Strategy

#### **Backend Testing**
- **Feature Tests**: End-to-end API testing with real HTTP requests
  - CRUD operations (Create, Read, Update, Delete)
  - Validation edge cases (enum case insensitivity, GUID trimming, amount validation)
  - Error responses (400, 404, 422 status codes)
  - Concurrency handling via mutex protection
  - Empty/invalid GUID handling
  - Partial search and limit validation
- **Manual Testing**: Comprehensive manual testing of all CRUD operations and edge cases

#### **Frontend Testing**
- **Unit Tests**: Pinia store testing with mocked API calls
  - Store CRUD operations (create, update, delete, fetch)
  - Search functionality with query string encoding
  - Error handling and state management
  - Array updates for proper Vue reactivity

## 🚀 Getting Started

### Prerequisites
- Docker & Docker Compose

### Development Setup

#### Docker Development
```bash
# Build and start all services
docker compose up --build -d

# Check service status
docker compose ps

# View logs
docker compose logs -f

# Stop services
docker compose down
```

#### Service URLs
Once the services are running, access the application at:
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080

## 🧪 Testing Commands

### Backend Testing
```bash
# Run specific test package
go test ./backend/tests/feature -v
```

### Frontend Testing
```bash
# Navigate to frontend directory
cd frontend

# Run unit tests
npm run test:unit
```

## 🔧 API Endpoints

- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080

### Items Management

| Method | Path | Request Body | Response | Notes |
|--------|------|--------------|----------|-------|
| **POST** | `/items` | `{amount, type, status, attributes}` | `201` Created / `400` Validation Error / `422` Invalid Data | Creates a new item; validation errors return structured JSON |
| **GET** | `/items?query=&limit=` | - | `200` OK (array) | Lists all items; filtered by query string |
| **GET** | `/items/:guid` | - | `200` OK / `404` Not Found | Fetches item by GUID |
| **PUT** | `/items/:guid` | `{amount?, type?, status?, attributes?}` | `200` OK / `400` Validation Error / `404` Not Found | Updates existing item; partial updates supported |
| **DELETE** | `/items/:guid` | - | `204` No Content / `404` Not Found | Deletes an item by GUID |

### Validation Error Response Format

```json
{
  "errors": {
    "amount": "This field is required",
    "type": "Invalid item type. Must be ADMISSION, SUBMISSION, or REVERSAL",
    "sort_code": "Sort code must be in the format 00-00-00",
    "account_number": "Must be exactly 8 digits"
  }
}
```

## ⚠️ Edge Cases & Limitations

### Handled Edge Cases
- **GUID Trimming**: Automatic whitespace trimming for GUID paths (`strings.TrimSpace()`)
- **Empty GUIDs**: Returns 404 for empty or whitespace-only GUIDs
- **Enum Validation**: Case-insensitive validation (accepts "admission", "ADMISSION", "Admission")
- **Amount Validation**: Enforces positive values only (`gt=0` validation tag)
- **Partial Search**: Supports query string matching across all item fields
- **Limit Validation**: Handles negative values, non-numeric inputs, and enforces default limits
- **Concurrency**: Uses `sync.RWMutex` for thread-safe operations on in-memory storage

### Current Limitations
- **In-Memory Storage**: Data is lost on application restart (will be replaced with PostgreSQL)
- **No Authentication**: Endpoints are publicly accessible (JWT auth to be implemented)
- **Limited Search**: Basic string matching only (Elasticsearch integration planned)
- **No Pagination**: All matching results returned (pagination logic to be added)

## 🔮 Future Enhancements

### Potential Improvements
- **Database Integration**: Replace in-memory storage with PostgreSQL/Redis for persistence and caching
- **Authentication System**: JWT-based authentication for production API security
- **Logging & Monitoring**: Structured logging and application monitoring for production readiness
- **Audit Trail**: Comprehensive audit logging for compliance and debugging