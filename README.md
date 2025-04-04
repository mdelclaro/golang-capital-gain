# Golang Capital Gain Calculator

This project is a Golang-based application that calculates capital gains based on a series of stock operations (e.g., `BUY` and `SELL`). The application processes input data and outputs the calculated taxes for each operation based on predetermined rules.

The program receives lists of stock market operations in JSON format, one per line, through standard input (stdin). Each line represents an independent simulation, and the program does not maintain state between lines. For each line of input, the program returns a list containing the tax paid for each operation in JSON format.

## Requirements

- Go 1.23 or later
- Docker (optional, for containerized execution)

## Project Structure

```
.
├── cmd/
│   └── main.go          # Entry point of the application
├── internal/
│   ├── domain/
│   │   ├── stocks/      # Logic for handling stock operations
│   │   └── tax/         # Logic for tax calculations
│   ├── pkg/
│   │   └── models/      # Data models for operations and tax output
│   └── utils/           # Utility functions
├── payloads/            # Input files for testing
├── tests/
│   └── integration_test.go # Integration tests
├── Dockerfile           # Dockerfile for containerized execution
├── go.mod               # Go module file
└── README.md            # Project documentation
```

## Install Dependencies

```bash
make install
```

## Usage

### Run Locally

You can run the application locally by providing input via standard input:

```bash
make run case=payloads/case_x
```

### Run with Docker

1. Build the Docker image:

   ```bash
   make docker-build
   ```

2. Run the container with an input file:

   ```bash
   make docker-run case=payloads/case_x
   ```

## Testing

### Run All Tests

```bash
make test
```