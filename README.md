# Monte Carlo Pi Simulation

A command-line application for calculating Pi using the Monte Carlo method.

## Requirements

- Go 1.x or higher

## Building and Running

This project uses a Makefile to simplify build and run operations.

### Available Commands

- `make help` - Show help information about available commands
- `make build` - Build the application (creates an executable named `pi`)
- `make run` - Build and run the application
- `make clean` - Remove build artifacts

### Examples

1. Build the application:
   ```bash
   make build
   ```

2. Run the application:
   ```bash
   make run
   ```
   
   Or after building:
   ```bash
   ./pi
   ```

3. Clean up build artifacts:
   ```bash
   make clean
   ```

## Application Usage

```
pi [options] [arguments]
```

### Options

- `-help`: Print help information
- `-version`: Print version information

### Examples

1. Run without arguments to see the welcome message:
   ```bash
   ./pi
   ```

2. Display help information:
   ```bash
   ./pi -help
   ```

3. Display version information:
   ```bash
   ./pi -version
   ```

4. Pass arguments to the program:
   ```bash
   ./pi arg1 arg2 arg3
   ```

## Development

This project uses Go modules for dependency management. To add a new dependency, use:

```bash
go get github.com/example/dependency
```

To run the program during development without using the Makefile:

```bash
go run main.go
```