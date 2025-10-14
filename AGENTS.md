# AGENTS.md - AI Agent Instructions

This document provides guidance for AI coding agents working with the go-lubelogger-api project.

## Project Overview

**go-lubelogger-api** is a Go client library that provides bindings for the [LubeLogger API](https://docs.lubelogger.com/Advanced/API), a vehicle maintenance tracking system. The library enables Go developers to programmatically interact with LubeLogger instances to manage vehicle records, odometer readings, service records, repairs, and more.

- **Language**: Go (minimum version 1.11)
- **Architecture**: Package-based modular design with separate packages for each API resource
- **Authentication**: HTTP Basic Authentication
- **API Requirements**: LubeLogger v1.4.6+ (requires culture-invariant API support)
- **Development Methodology**: Test Driven Development (TDD)

## Repository Structure

```
.
├── lubelogger.go           # Main entry point with NewClient()
├── client/                 # HTTP client implementation and request handling
│   ├── client.go          # Core client with DoRequest() method
│   ├── query.go           # Request query builder
│   └── response.go        # Response handling
├── odometer/              # Odometer records API
├── servicerecord/         # Service records API
├── repairrecords/         # Repair records API
├── upgraderecords/        # Upgrade records API
├── taxrecords/            # Tax records API
├── gasrecords/            # Gas/fuel records API
├── reminders/             # Reminders API
├── vehicles/              # Vehicle listing API
├── vehicleinfo/           # Vehicle information API
├── calendar/              # Calendar export API
├── document/              # Document upload API
├── root/                  # Root-level operations (backup, cleanup)
├── parser/                # Data parsing utilities
├── debuglog/              # Debug logging support
├── example/               # Example usage with Docker Compose setup
└── test/                  # Test utilities
```

## Key Design Patterns

### 1. Client Initialization
```go
c := lubelogger.NewClient("https://demo.lubelogger.com", "username", "password")
```

### 2. Resource Operations Pattern
Each resource package (odometer, servicerecord, etc.) follows a consistent pattern:

- **GetRecords()**: Fetch all records for a vehicle
- **Add()**: Create a new record (method on struct)
- **Update()**: Update an existing record (method on struct)
- **Delete()**: Delete a record (method on struct)

### 3. Context-Based Operations
All API calls accept a `context.Context` as the first parameter for cancellation and timeout support.

### 4. Client-Vehicle ID Pattern
Most operations require both a `client.Client` and a `vehicleId int64` parameter.

## Code Conventions

### Naming
- Package names: lowercase, singular (e.g., `odometer`, `servicerecord`)
- Struct names: PascalCase (e.g., `OdometerRecord`, `ServiceRecord`)
- Method names: PascalCase (exported), camelCase (unexported)
- Constants: UPPER_SNAKE_CASE (e.g., `URGENCY_NOT_URGENT`)

### File Organization
Each resource package typically contains:
- `{resource}.go` - Main API functions (GetRecords, Add, Update, Delete)
- `structs.go` - Data structures and type definitions
- `{resource}_test.go` - Unit tests (where applicable)

### Error Handling
- Return errors explicitly; never panic in library code
- Use `fmt.Errorf()` for error wrapping with context
- HTTP errors include status code and response body

### HTTP Requests
- Always set `culture-invariant: true` header (required for v1.4.6+)
- Use HTTP Basic Authentication via Authorization header
- Support custom HTTP clients via `client.HttpClient()` method

## Development Guidelines

### Test Driven Development (TDD)

This project follows **Test Driven Development** practices. When adding new features or fixing bugs:

1. **Write tests first**: Create failing tests that define the expected behavior
2. **Implement the minimum code**: Write just enough code to make the tests pass
3. **Refactor**: Clean up the code while keeping tests green
4. **Repeat**: Continue the cycle for each new feature or change

### Adding a New Resource API

Follow the TDD approach:

1. **Create test file first**: `mkdir newresource && touch newresource/newresource_test.go`
2. **Write failing tests** for the expected API behavior
3. **Create structs**: Add `structs.go` with data structures
4. **Implement functionality**: Create `newresource.go` with API functions:
   ```go
   package newresource
   
   import (
       "context"
       "github.com/thaapaniemi/go-lubelogger-api/client"
   )
   
   func GetRecords(ctx context.Context, c client.Client, vehicleId int64) ([]Record, error) {
       // Implementation
   }
   
   func (r Record) Add(ctx context.Context, c client.Client, vehicleId int64) error {
       // Implementation
   }
   ```
5. **Run tests**: Ensure all tests pass with `go test ./newresource`
6. **Refactor**: Improve code quality while maintaining passing tests
7. **Update documentation**: Add new functions to README.md

### Testing

Following TDD principles:

- **Write tests first**: Before implementing any feature, write the test
- **Unit tests**: Should be self-contained where possible and test individual functions
- **Integration tests**: Require a running LubeLogger instance (see `example/docker-compose.yml`)
- **Test instance**: Use `http://127.0.0.1:8080` with credentials `test/1234`
- **Test coverage**: Aim for comprehensive coverage of all public APIs
- **Run tests frequently**: Execute `go test ./...` after every change
- **Test naming**: Use descriptive test names that explain what is being tested (e.g., `TestGetRecords_WithValidVehicleId`)

### Debug Logging
Enable debug output with:
```go
import "github.com/thaapaniemi/go-lubelogger-api/debuglog"
debuglog.Enabled = true
```

This logs:
- Request URLs, payloads, and content types
- Response status codes and bodies

### Custom HTTP Client
Override the default HTTP client for custom timeout, proxy, or TLS settings:
```go
httpClient := &http.Client{
    Timeout: 10 * time.Second,
}
c := lubelogger.NewClient(...)
c.HttpClient(httpClient)
```

## Common Pitfalls

1. **Culture-Invariant Header**: Always ensure the client sets `culture-invariant: true` header. This is critical for proper date/number parsing with LubeLogger v1.4.6+.

2. **Date Handling**: Use Go's `time.Time` type consistently. The parser package handles conversion to/from LubeLogger's date format.

3. **JSON Number Parsing**: The client uses `json.Decoder.UseNumber()` to handle numeric fields correctly during parsing.

4. **Vehicle ID**: Most API calls require a valid vehicle ID. Get vehicle IDs via `vehicles.GetRecords()` first.

5. **HTTP Method**: Ensure correct HTTP method (GET, POST, PUT, DELETE) is used for each operation.

## Testing Strategy

### Running the Example
```bash
cd example
docker-compose up -d
go run main.go
```

### Test Coverage Areas
- ✅ Vehicle listing and info retrieval
- ✅ CRUD operations for all record types
- ✅ Calendar export
- ✅ Document upload
- ✅ Reminder management
- ⚠️ Backup/cleanup operations (use with caution)

## API Reference Quick Guide

### Core Client
```go
lubelogger.NewClient(endpoint, username, password string) client.Client
client.HttpClient(newClient *http.Client)
client.DoRequest(ctx context.Context, r LubeLoggerRequest) ([]byte, error)
client.Decode(in []byte, out interface{}) error
```

### Vehicle Operations
```go
vehicles.GetRecords(ctx, c) ([]VehicleData, error)
vehicleinfo.GetRecords(ctx, c, vehicleId) ([]VehicleInfo, error)
```

### Record Operations (Pattern applies to all record types)
```go
// odometer, servicerecord, repairrecords, upgraderecords, taxrecords, gasrecords
GetRecords(ctx, c, vehicleId) ([]Record, error)
(r Record) Add(ctx, c, vehicleId) error
(r Record) Update(ctx, c) error
(r Record) Delete(ctx, c) error
```

### Special Operations
```go
reminders.GetRecords(ctx, c, vehicleId) ([]Reminder, error)
reminders.SendReminderEmails(ctx, c, urgencies []Urgency) error
calendar.GetCalendar(ctx, c) (string, error)
document.Upload(ctx, c) (string, error)  // Method on Document struct
root.MakeBackup(ctx, c) ([]byte, error)
root.Cleanup(ctx, c) ([]byte, error)
```

## Building and Dependencies

### Build Commands
```bash
go build                          # Build the library
go test ./...                     # Run all tests
go get -u                         # Update dependencies
```

### Dependencies
- **Standard library only**: No external dependencies required
- Uses only Go standard library packages (`net/http`, `encoding/json`, `context`, etc.)

## Contributing Guidelines

When contributing to this project:

1. **Follow TDD**: Write tests before implementation
2. **Minimal changes**: Make surgical, precise modifications
3. **Consistency**: Follow existing patterns in similar packages
4. **Testing**: Ensure all tests pass and changes don't break existing functionality
5. **Test coverage**: Add tests for new features and bug fixes
6. **Documentation**: Update README.md with new API functions
7. **No breaking changes**: This is an early-stage project, but avoid unnecessary API changes
8. **Error context**: Include meaningful error messages with context

## Known Issues & Limitations

1. **Development Status**: Project is in active development; stability not guaranteed
2. **API Version**: Requires LubeLogger v1.4.6+ for vehicles/vehicleinfo packages
3. **Error Handling**: Some API endpoints may return unexpected response formats
4. **Type Conversions**: Parser handles various data type conversions; edge cases may exist
5. **Date Formats**: Dependent on LubeLogger's date format; culture-invariant mode required

## Example Usage

See `example/main.go` for comprehensive usage examples covering:
- Client initialization with debug logging
- Fetching vehicle information
- CRUD operations for all record types
- Calendar export
- Document upload
- Reminder management

## Support Resources

- **LubeLogger Documentation**: https://docs.lubelogger.com/Advanced/API
- **LubeLogger Website**: https://lubelogger.com/
- **API Reference**: https://github.com/hargata/lubelog_scripts/blob/main/misc/LubeLogger.postman_collection.json
- **Repository**: https://github.com/thaapaniemi/go-lubelogger-api

## License

MIT License - See LICENSE file for details
