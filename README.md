# Pet Management API

A RESTful API built with Go and Gin framework for managing pets and their owners.

## Features

- Pet management (CRUD operations)
- Owner management (CRUD operations)
- RESTful API design
- CORS support
- Health check endpoint
- Environment configuration

## Prerequisites

- Go 1.21 or higher
- Git

## Installation

1. Clone the repository:

```bash
git clone <repository-url>
cd pet-manage-be
```

2. Install dependencies:

```bash
go mod tidy
```

3. Copy environment configuration:

```bash
cp config.env.example .env
```

4. Update the `.env` file with your configuration.

## Running the Application

1. Start the server:

```bash
go run main.go
```

The server will start on port 8080 by default.

## API Endpoints

### Health Check

- `GET /health` - Check if the API is running

### Pets

- `GET /api/v1/pets` - Get all pets
- `GET /api/v1/pets/:id` - Get a specific pet
- `POST /api/v1/pets` - Create a new pet
- `PUT /api/v1/pets/:id` - Update a pet
- `DELETE /api/v1/pets/:id` - Delete a pet

### Owners

- `GET /api/v1/owners` - Get all owners
- `GET /api/v1/owners/:id` - Get a specific owner
- `POST /api/v1/owners` - Create a new owner
- `PUT /api/v1/owners/:id` - Update an owner
- `DELETE /api/v1/owners/:id` - Delete an owner

## Project Structure

```
pet-manage-be/
├── handlers/          # HTTP handlers
│   ├── pet_handlers.go
│   └── owner_handlers.go
├── models/            # Data models
│   ├── pet.go
│   └── owner.go
├── main.go           # Application entry point
├── go.mod           # Go module file
├── go.sum           # Go dependencies checksum
└── README.md        # This file
```

## Development

### Adding New Features

1. Create models in the `models/` directory
2. Create handlers in the `handlers/` directory
3. Add routes in `main.go`
4. Update this README with new endpoints

### Database Integration

Currently, the API returns mock data. To integrate with a real database:

1. Add database driver to `go.mod`
2. Create database connection in a new `database/` package
3. Update handlers to use database queries
4. Add database migrations

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests if applicable
5. Submit a pull request

## License

This project is licensed under the MIT License.
