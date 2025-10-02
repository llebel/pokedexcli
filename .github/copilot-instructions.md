# Pokedex CLI Development Guide

## Architecture Overview

This is a Go CLI application implementing a Pokemon location browser using the PokeAPI. The architecture follows a clean separation of concerns:

- **Main package**: Contains REPL loop and command definitions (`main.go`, `repl.go`, `command_*.go`)
- **pokeapi package**: HTTP client with caching for PokeAPI interactions (`internal/pokeapi/`)
- **pokecache package**: Thread-safe in-memory cache with TTL expiration (`internal/pokecache/`)

## Key Patterns

### Command Registration
Commands are registered in `getCommands()` in `repl.go` using this pattern:
```go
"command_name": {
    name: "command_name",
    description: "Description here", 
    callback: commandFunctionName,
}
```
All command functions must match signature: `func(*cliConfig) error`

### State Management
The `cliConfig` struct carries application state through the REPL:
- `pokeapiClient`: HTTP client instance
- `nextLocationsURL`/`prevLocationsURL`: Pagination state for API calls
- Future: Add `pokeCache` field when implementing caching in REPL

### HTTP Client with Caching
The pokeapi client follows this pattern:
1. Check cache first using URL as key
2. Make HTTP request if cache miss
3. Store response in cache before returning
4. Use `pokecache.Cache` for thread-safe operations

### Testing Conventions
- Use table-driven tests with `cases` slice containing input/expected pairs
- Test files named `*_test.go` in same package
- Cache tests use short intervals (5ms) for TTL testing
- Include edge cases like empty input and boundary conditions

## Development Workflow

### Running the Application
```bash
go run .
# Then use commands: help, map, mapb, exit
```

### Testing
```bash
go test ./...                    # Run all tests
go test ./internal/pokecache/    # Test specific package
go test -v                       # Verbose output
```

### Adding New Commands
1. Create `command_newname.go` with function matching `func(*cliConfig) error`
2. Add entry to `getCommands()` map in `repl.go`
3. Handle any new state in `cliConfig` struct

### Cache Integration
- Cache keys should be full URLs for API endpoints
- Use `c.cache.Get(url)` before HTTP calls
- Use `c.cache.Add(url, responseBytes)` after successful responses
- Cache operates on `[]byte` data, unmarshal after retrieval

## File Organization

- Root level: Main package files, each command in separate file
- `internal/pokeapi/`: API client, type definitions, endpoint methods
- `internal/pokecache/`: Generic cache implementation with operations
- Tests co-located with implementation files

## External Dependencies

- PokeAPI v2 (`https://pokeapi.co/api/v2`) for Pokemon location data
- Standard library only (no external packages required)
- Thread-safe cache using `sync.RWMutex` for concurrent access