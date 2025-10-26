# Go Test Application

A full-stack application built with Go backend and Vue 3 frontend, demonstrating modern web development practices with comprehensive testing and Docker deployment.

## 👤 Developer Notes

**Developer**: Olanrewaju Sule-Balogun

**Technology Experience**:
- **Vue.js**: Experienced with Vue 3 Composition API, Pinia, TypeScript, and modern frontend development
- **Go**: Still relatively new to Go, but bringing experience from other ecosystems
  - Applying patterns from Laravel/PHP (validation, DDD, repository pattern) to Go
  - Learning Go idioms while leveraging established architectural principles
  - Exploring Gin framework, Go's type system, and idiomatic patterns

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

#### **Consistent API Responses**
- **Rationale**: Predictable response formats for frontend consumption
- **Implementation**:
  - Single objects for create/update operations
  - Arrays for list operations
  - Proper HTTP status codes (200, 201, 204, 400, 404, 500)

### Frontend Architecture

#### **Vue 3 Composition API**
- **Rationale**: Better TypeScript support, improved code organization, and reusability
- **Implementation**: `<script setup>` syntax throughout all components

#### **Pinia State Management**
- **Rationale**: Official Vue 3 state management with better TypeScript support than Vuex
- **Implementation**: 
  - Centralized store for items data
  - Reactive state with computed properties
  - Immutable array updates for proper Vue reactivity

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
- **Manual Testing**: Comprehensive manual testing of all CRUD operations and edge cases
- **Test Coverage**: Full coverage of API endpoints, validation, and error scenarios

#### **Frontend Testing**
- **Unit Tests**: Pinia store testing, for CRUD operations with mocked API calls

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
- **API Documentation**: http://localhost:8080/items

### Items Management
- `GET /items` - List all items (with optional search query)
- `GET /items/:guid` - Get item by GUID
- `POST /items` - Create new item
- `PUT /items/:guid` - Update existing item
- `DELETE /items/:guid` - Delete item

## 🔒 Security & Best Practices

### Backend Security
- **Input Validation**: Comprehensive validation on all inputs
- **CORS Configuration**: Proper cross-origin resource sharing setup
- **Error Handling**: Secure error responses without sensitive information
- **Type Safety**: Strong typing throughout the application

### Frontend Security
- **TypeScript**: Compile-time type checking
- **Input Sanitization**: Proper handling of user inputs
- **XSS Prevention**: Safe rendering of dynamic content
- **CSRF Protection**: Proper API request handling

## 🔮 Future Enhancements

### Potential Improvements
- **Database Integration**: Replace in-memory storage with persistent database
- **Authentication System**: JWT-based user authentication
- **Advanced Search**: Full-text search with Elasticsearch
- **Audit Logging**: Comprehensive audit trail system
---