# Agnos Backend API

Hospital Management System Backend built with Go, Gin, PostgreSQL, and Docker.

## 🏗️ Project Structure

```
Agnos-backend/
├── app/
│   ├── console/                    # CLI commands
│   │   ├── cmd.go                 # Command definitions
│   │   └── kernel.go              # Console kernel
│   ├── enum/                      # Enumerations
│   │   └── status.go              # Status Enum
│   ├── helper/                    # Helper utilities
│   │   └── helper.go             # General helpers
│   ├── message/                   # Message constants
│   │   └── main.go               # Message definitions
│   ├── middleware/                # HTTP middleware
│   │   └── auth.middleware.go    # Authentication middleware
│   ├── model/                     # Database models
│   │   ├── 0-base.go             # Base model
│   │   ├── patient.model.go      # Patient model
│   │   └── staff.model.go        # Staff model
│   ├── modules/                   # Feature modules
│   │   ├── modules.go            # Module registration
│   │   ├── patient/              # Patient module
│   │   │   ├── ctl.patient.go    # Patient controller
│   │   │   ├── mod.patient.go    # Patient module
│   │   │   ├── sv.patient.go     # Patient service
│   │   │   ├── inf.patient.go    # Patient interface
│   │   │   ├── controller_test.go # Patient tests
│   │   │   └── dto/              # Patient DTOs
│   │   │       └── dto.patient.go
│   │   └── staff/                # Staff module
│   │       ├── ctl.staff.go      # Staff controller
│   │       ├── mod.staff.go      # Staff module
│   │       ├── sv.staff.go       # Staff service
│   │       ├── inf.staff.go      # Staff interface
│   │       ├── controller_test.go # Staff tests
│   │       └── dto/              # Staff DTOs
│   │           └── dto.staff.go
│   ├── response/                  # HTTP response helpers
│   │   └── response.go           # Response formatting
│   ├── routes/                    # Route definitions
│   │   ├── routes.go             # Main routes
│   │   ├── patient.go            # Patient routes
│   │   └── staff.go              # Staff routes
│   └── util/                      # Utilities
│       ├── hashing/              # Password hashing
│       │   └── main.go
│       └── jwt/                  # JWT utilities
│           └── main.go
├── config/                        # Configuration
│   ├── config.go                 # Main config
│   ├── database.go               # Database config
│   └── helper.go                 # Config helpers
│   
├── database/                      # Database related
│   ├── migrations/               # Database migrations
│   │   └── Models.go
│   └── seeds/                    # Database seeds
│       ├── 0-base.go
│       └── mockUp.go
├── internal/                      # Internal packages
│   ├── interface.go              # Internal interfaces
│   ├── cmd/                      # Command framework
│   │   ├── cmd.go
│   │   ├── httpCmd.go
│   │   ├── migrateCmd.go
│   │   └── model.go
│   └── logger/                   # Logging system
│       └── logger.go
├── nginx/                         # Nginx configuration
│   ├── nginx.conf                # Main nginx config
│   └── default.conf              # Server config
├── docker-compose.yml             # Docker compose
├── Dockerfile                     # Docker build
├── Makefile                       # Build commands
├── main.go                        # Application entry point
├── go.mod                         # Go module
├── go.sum                         # Go dependencies
└── README.md                      # This file
```

## 🚀 Quick Start

### Prerequisites

- Docker & Docker Compose
- Go 1.21+ (for local development)

### Using Docker (Recommended)

```bash
# Clone the repository
git clone <repository-url>
cd Agnos-backend

# Start all services with Nginx
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Local Development

```bash
# Install dependencies
go mod download

# Set environment variables
export DB_HOST=localhost
export DB_PORT=5432
export DB_DATABASE=postgres
export DB_USER=root
export DB_PASSWORD=secret
export JWT_SECRET=secret

# Run database migrations
migrate-up:
	go run . migrate up

migrate-down:
	go run . migrate down

migrate-seed:
	go run . migrate seed

migrate-refresh:
	go run . migrate refresh


# Start the server
	go run . http

```

## 📡 API Endpoints

### Base URL

- **Docker**: `http://localhost/api`
- **Local**: `http://localhost:8080`

### Authentication

All authenticated endpoints require a JWT token in the Authorization header:

```
Authorization: Bearer <jwt-token>
```

### Staff Endpoints

#### Create Staff

```http
POST /staff
Content-Type: application/json

{
  "username": "doctor01",
  "password": "securepassword",
  "hospital": "hospital-a"
}
```

#### Staff Login

```http
POST /staff/login
Content-Type: application/json

{
  "username": "doctor01",
  "password": "securepassword",
  "hospital": "hospital-a"
}
```

**Response:**

```json
{
  "code": 200,
  "message": "Success",
  "data": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9...",
}
```

### Patient Endpoints

> **Note**: All patient endpoints require authentication

#### Get Patient by ID

```http
GET /patient/{id}
Authorization: Bearer <jwt-token>
```

#### List Patients

```http
GET /patients?page=1&size=10&orderBy=asc&sortBy=created_at
Authorization: Bearer <jwt-token>
```

**Query Parameters:**

- `page` (int): Page number (default: 1)
- `size` (int): Items per page (default: 10)
- `orderBy` (string): Sort order - "asc" or "desc" (default: "asc")
- `sortBy` (string): Sort field (default: "created_at")
- `firstName` (string): Filter by first name
- `lastName` (string): Filter by last name
- `middleName` (string): Filter by middle name
- `email` (string): Filter by email
- `phoneNumber` (string): Filter by phone number
- `nationalId` (string): Filter by national ID
- `passportId` (string): Filter by passport ID
- `dateOfBirth` (string): Filter by date of birth (YYYY-MM-DD)

**Response:**

```json
{
  "code": 200,
  "message": "Success",
  "data": [
    {
            "id": "65e08e33-9f57-45fe-b725-82242e3581ad",
            "first_name_th": "asdas",
            "middle_name_th": "asdsa",
            "last_name_th": "dsadsadsa",
            "first_name_en": "dsad",
            "middle_name_en": "sadsad",
            "last_name_en": "sadsa",
            "date_of_birth": "0001-01-01T00:00:00Z",
            "patient_hn": "asdsa",
            "national_id": "906976976",
            "passport_id": "976976976967",
            "phone_number": "8787567865756",
            "email": "asdasd@gdsgdsg.com",
            "gender": "1",
            "hospital": "a",
            "created_at": 1754363674,
            "updated_at": 1754363674,
            "deleted_at": null
        }
  ],
  "pagination": {
    "page": 1,
    "size": 10,
    "total": 1
  }
}
```

## 🧪 Testing

### Run Tests

```bash
# Run all tests
go test ./...

# Run specific module tests
go run . cmd test patient
go run . cmd test staff

# Run test summary
./simple_test_summary.sh
```

### Test Structure

- **Focus**: Success ✅ and Fail ❌ scenarios only
- **Coverage**: Controller unit tests with mocking
- **Pattern**: Interface-based dependency injection

## 🐳 Docker Services

| Service  | Port    | Description                     |
| -------- | ------- | ------------------------------- |
| nginx    | 80, 443 | Reverse proxy and load balancer |
| main     | 8080    | Go application server           |
| postgres | 5432    | PostgreSQL database             |

## 🔧 Environment Variables

| Variable           | Description            | Default      |
| ------------------ | ---------------------- | ------------ |
| `DEBUG`            | Enable debug mode      | `false`      |
| `PORT`             | Application port       | `8080`       |
| `DB_HOST`          | Database host          | `localhost`  |
| `DB_PORT`          | Database port          | `5432`       |
| `DB_DATABASE`      | Database name          | `postgres`   |
| `DB_USER`          | Database user          | `root`       |
| `DB_PASSWORD`      | Database password      | `secret`     |
| `JWT_SECRET`       | JWT signing secret     | `secret`     |
| `JWT_DURATION`     | JWT expiration (hours) | `720`        |
| `HTTP_JSON_NAMING` | JSON naming convention | `camel_case` |

## 📝 Development

### Adding New Features

1. **Create Module Structure**:

   ```
   app/modules/feature/
   ├── ctl.feature.go      # Controller
   ├── sv.feature.go       # Service
   ├── mod.feature.go      # Module registration
   ├── inf.feature.go      # Interface
   ├── controller_test.go  # Tests
   └── dto/
       └── dto.feature.go  # Data transfer objects
   ```

2. **Register Routes**: Add routes in `app/routes/`
3. **Add Tests**: Create controller tests focusing on success/fail scenarios
4. **Update Documentation**: Update this README with new endpoints

### Database Migrations

```bash
# Run migrations
go run . cmd migrate

# Create new migration
# Add your migration in database/migrations/
```

### CLI Commands

```bash
# Available commands
go run . cmd

# Test commands
go run . cmd test
go run . cmd test patient
go run . cmd test staff

# HTTP server
go run . cmd http

# Database migration
go run . cmd migrate

# Hello world
go run . cmd hello
```

## 🏥 Hospital System Features

### Staff Management

- Staff registration and authentication
- JWT-based session management
- Hospital-specific access control

### Patient Management

- Patient record retrieval
- Advanced filtering and search
- Pagination support
- Hospital-based data isolation

## 🔒 Security

- JWT authentication for API access
- Password hashing with bcrypt
- Hospital-based data segregation
- Nginx security headers
- Input validation and sanitization

## 📊 Monitoring

### Health Check

```http
GET /health
```

### Nginx Access Logs

```bash
docker-compose logs nginx
```

### Application Logs

```bash
docker-compose logs main
```

## 🚀 Deployment

### Production Setup

1. **Update Environment Variables**:

   ```bash
   # Set secure values
   JWT_SECRET=your-super-secure-secret
   DB_PASSWORD=strong-database-password
   ```

2. **SSL/HTTPS Setup**:

   - Update `nginx/default.conf` for SSL
   - Add SSL certificates to nginx volume

3. **Deploy**:
   ```bash
   docker-compose -f docker-compose.prod.yml up -d
   ```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

---

**Built with ❤️ using Go, Gin, PostgreSQL, Docker, and Nginx**
