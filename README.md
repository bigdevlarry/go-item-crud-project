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
├── backend/                    # Go backend application
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
│       ├── create_item_test.go
│       ├── update_item_test.go
│       └── feature/           # Feature/integration tests
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
│   │   │       └── items.test.ts
│   │   ├── types/             # TypeScript type definitions
│   │   │   ├── entities.ts
│   │   │   ├── dto.ts
│   │   │   ├── enums.ts
│   │   │   ├── api.ts
│   │   │   ├── components.ts
│   │   │   ├── filters.ts
│   │   │   └── index.ts
│   │   ├── utils/             # Frontend utilities
│   │   │   └── formatters.ts
│   │   ├── views/             # Page components
│   │   │   └── HomeView.vue
│   │   ├── router/            # Vue Router configuration
│   │   │   └── index.ts
│   │   ├── assets/            # Static assets
│   │   ├── App.vue            # Root component
│   │   └── main.ts            # Application entry point
│   ├── Dockerfile             # Frontend container configuration
│   ├── Dockerfile.dev         # Development container
│   ├── package.json           # Node.js dependencies
│   ├── vite.config.ts         # Vite build configuration
│   ├── tailwind.config.js     # Tailwind CSS configuration
│   └── vitest.config.ts       # Test configuration
├── docker-compose.yml         # Multi-container deployment
├── Dockerfile                 # Backend container configuration
├── go.mod                     # Go module dependencies
├── go.sum                     # Go dependency checksums
├── go.work                    # Go workspace configuration
└── main.go                    # Backend application entry point
```

## 🎯 Design Decisions & Technical Choices

### Backend Architecture

#### **Domain-Driven Design (DDD)**
- **Rationale**: Clear separation of business logic from infrastructure concerns
- **Implementation**: 
  - `domain/models/` for business entities
  - `domain/dto/` for data transfer objects
  - `domain/validators/` for business rules
  - `repository/` for data access abstraction

#### **Custom Validation System**
- **Rationale**: Provide user-friendly error messages instead of technical validation errors
- **Implementation**:
  - Custom validators for `itemtype`, `itemstatus`, and `sortcode`
  - Structured error responses: `{"errors": {"field": "message"}}`
  - JSON parsing error handling for type mismatches

#### **Repository Pattern**
- **Rationale**: Abstract data access layer for testability and future extensibility
- **Implementation**: `ItemsStorage` interface with `InMemoryStore` implementation

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
- **Unit Tests**: Individual function testing
- **Feature Tests**: End-to-end API testing with real HTTP requests
- **Test Coverage**: Comprehensive coverage of all CRUD operations and edge cases

#### **Frontend Testing**
- **Unit Tests**: Pinia store testing with mocked API calls
- **Component Testing**: Vue Test Utils for component behavior verification
- **Mock Strategy**: Comprehensive fetch API mocking for isolated testing

## 🚀 Getting Started

### Prerequisites
- Go 1.21+
- Node.js 18+
- Docker & Docker Compose

### Development Setup

#### Backend Development
```bash
# Install dependencies
go mod download

# Run tests
go test ./... -v -cover

# Run application
go run main.go
```

#### Frontend Development
```bash
# Navigate to frontend directory
cd frontend

# Install dependencies
npm install

# Run development server
npm run dev

# Run tests
npm run test:unit

# Build for production
npm run build
```

### Docker Deployment

#### Single Command Deployment
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
- **Frontend**: http://localhost:3000
- **Backend API**: http://localhost:8080
- **API Documentation**: http://localhost:8080/items

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

### Items Management
- `GET /items` - List all items (with optional search query)
- `GET /items/:guid` - Get item by GUID
- `POST /items` - Create new item
- `PUT /items/:guid` - Update existing item
- `DELETE /items/:guid` - Delete item

## 🎨 Frontend Features

### User Interface
- **Responsive Design**: Works on desktop and mobile devices
- **Modern UI**: Clean, professional interface using Tailwind CSS
- **Interactive Elements**: Collapsible rows, modals, and toast notifications
- **Search Functionality**: Real-time search with debounced API calls

### State Management
- **Centralized State**: All items data managed in Pinia store
- **Reactive Updates**: Real-time UI updates on data changes
- **Error Handling**: Comprehensive error states and user feedback
- **Loading States**: Visual feedback during API operations

### Form Handling
- **Validation**: Client-side validation with server-side error display
- **User Experience**: Intuitive form interactions with proper feedback
- **Data Persistence**: Form state management during create/edit operations

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

## 📊 Performance Considerations

### Backend Performance
- **In-Memory Storage**: Fast data access for development/demo purposes
- **Efficient Validation**: Minimal overhead validation system
- **Structured Responses**: Optimized JSON response formats

### Frontend Performance
- **Component Lazy Loading**: Efficient component loading strategy
- **Reactive Updates**: Minimal DOM updates using Vue's reactivity
- **Bundle Optimization**: Vite-based build optimization
- **Caching Strategy**: Proper API response caching

## 🚀 Deployment & Production

### Docker Configuration
- **Multi-stage Builds**: Optimized container images
- **Health Checks**: Service health monitoring
- **Restart Policies**: Automatic service recovery
- **Environment Variables**: Configurable deployment settings

### Production Considerations
- **Database Integration**: Ready for PostgreSQL/MySQL integration
- **Authentication**: Prepared for JWT-based authentication
- **Logging**: Structured logging implementation
- **Monitoring**: Health check endpoints for monitoring

## 🧪 Test Coverage

### Backend Test Coverage
- **CRUD Operations**: 100% coverage of all endpoints
- **Validation Logic**: Comprehensive validation testing
- **Error Scenarios**: Edge case and error condition testing
- **Integration Tests**: End-to-end API testing

### Frontend Test Coverage
- **Store Logic**: Complete Pinia store testing
- **API Integration**: Mocked API interaction testing
- **Error Handling**: Error state and recovery testing
- **User Interactions**: Component behavior verification

## 🔮 Future Enhancements

### Potential Improvements
- **Database Integration**: Replace in-memory storage with persistent database
- **Authentication System**: JWT-based user authentication
- **Real-time Updates**: WebSocket integration for live updates
- **Advanced Search**: Full-text search with Elasticsearch
- **File Uploads**: Document and image upload capabilities
- **Audit Logging**: Comprehensive audit trail system

### Scalability Considerations
- **Microservices**: Break down into smaller, focused services
- **Caching Layer**: Redis integration for improved performance
- **Load Balancing**: Multiple backend instances with load balancing
- **CDN Integration**: Static asset delivery optimization

## 📝 Development Notes

### Code Quality
- **Linting**: ESLint and Prettier for consistent code formatting
- **Type Checking**: Comprehensive TypeScript type coverage
- **Documentation**: Inline code documentation and README
- **Git Workflow**: Proper version control practices

### Maintenance
- **Dependency Updates**: Regular dependency updates and security patches
- **Code Reviews**: Comprehensive code review process
- **Testing**: Continuous integration with automated testing
- **Monitoring**: Production monitoring and alerting setup

---

## 🎯 Summary

This application demonstrates modern full-stack development practices with:

- **Clean Architecture**: Domain-driven design with clear separation of concerns
- **Modern Frontend**: Vue 3 Composition API with TypeScript and Pinia
- **Robust Backend**: Go with Gin framework and comprehensive validation
- **Comprehensive Testing**: Both unit and integration tests
- **Production Ready**: Docker deployment with health checks and monitoring
- **Developer Experience**: Hot reload, TypeScript support, and comprehensive tooling

The codebase is maintainable, scalable, and follows industry best practices for both development and deployment.
