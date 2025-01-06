# Golang Backend Boilerplate

This project is a boilerplate for building backend services using Go. It provides a structured foundation with essential features and configurations to kickstart your development process.

## Project Structure

- **cmd/**: Contains the main application entry point.
- **internals/**: Contains the core business logic and internal packages.
- **pkg/**: Contains utility packages and shared code.
- **test/**: Contains test files for unit and integration testing.
- **utils/**: Contains utility functions and helpers.

## Prerequisites

- Go 1.22.5 or later
- Docker (for containerization)
- PostgreSQL (as the database)

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/iacopoghilardi/golang-backend-boilerplate.git
   cd golang-backend-boilerplate
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set up environment variables:
   - Copy `.env.example` to `.env` and configure the necessary environment variables.

## Usage

- To run the application:
  ```bash
  go run cmd/main.go
  ```

- To build the application:
  ```bash
  go build -o bin/app cmd/main.go
  ```

## Testing

- Run tests using:
  ```bash
  go test ./...
  ```

## Dependencies

- [Gin Gonic](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [Viper](https://github.com/spf13/viper) - Configuration management
- [Testify](https://github.com/stretchr/testify) - Testing utilities
- [Testcontainers](https://github.com/testcontainers/testcontainers-go) - Containerized testing

## License

This project is licensed under the MIT License.

## Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

## Contact

For any inquiries, please contact [your-email@example.com].
