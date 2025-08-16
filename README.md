# Monte Carlo Pi Simulation

A command-line application for calculating Pi (π) using the Monte Carlo method.

## What is the Monte Carlo Method?

The Monte Carlo method for calculating π works by:

1. Creating a square with side length 2r, centered at the origin
2. Inscribing a circle with radius r inside the square
3. Randomly generating points within the square
4. Counting how many points fall inside the circle
5. The ratio of points inside the circle to total points approaches π/4 as the number of points increases

So π ≈ 4 * (points inside circle / total points)

The more points you generate, the more accurate the approximation becomes.

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
- `-points`: Number of points to use in Monte Carlo simulation (default: 1000000)

### Examples

1. Calculate π using the default number of points (1,000,000):
   ```bash
   ./pi
   ```

2. Calculate π using a specific number of points:
   ```bash
   ./pi -points 10000000
   ```

3. Calculate π by passing the number of points as an argument:
   ```bash
   ./pi 5000000
   ```

4. Display help information:
   ```bash
   ./pi -help
   ```

5. Display version information:
   ```bash
   ./pi -version
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