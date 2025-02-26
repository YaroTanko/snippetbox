# SnippetBox Development Guide

## Build & Run Commands
- Run application: `task run` or `go run .`
- Build application: `task build`
- Clean build artifacts: `task clean`

## Test Commands
- Run all tests: `task test` or `go test -v ./...`
- Run single test: `go test -v -run ^TestName$ ./...`
- Run specific test file: `go test -v ./path/to/file_test.go`

## Code Style Guidelines
- **Imports**: Group standard library imports first, then third-party packages
- **Formatting**: Use `gofmt` for consistent formatting
- **Naming**: Use camelCase for variables, PascalCase for exported functions/types
- **Error Handling**: Check all errors, avoid using `_` for errors
- **Functions**: Start HTTP handler functions with lowercase unless exported
- **Testing**: Name test functions using descriptive names (Test + what it tests)
- **HTTP**: Use http.ResponseWriter before *http.Request in handler signatures
- **Comments**: Add meaningful comments for exported functions

## Project Structure
This is a Go web application (SnippetBox) using the standard library's net/http package without external frameworks.