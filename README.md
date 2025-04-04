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
│   │   │   ├── service.go       # Core logic for stock operations
│   │   │   └── service_test.go  # Unit tests for stock operations
│   │   └── tax/         # Logic for tax calculations
│   │       ├── service.go       # Core logic for tax calculations
│   │       └── service_test.go  # Unit tests for tax calculations
│   ├── pkg/
│   │   └── models/      # Data models for operations and tax output
│   │       ├── operations.go     # Model for stock operations
│   │       └── tax.go    # Model for tax output
│   └── utils/           # Utility functions
│           └── math.go       # Helper functions for math operations
│           └── math_test.go       # Unit tests
├── payloads/            # Input files for testing
│   ├── case_1      # Example input for Use Case #1
│   ├── case_2      # Example input for Use Case #2
│   └── case_3      # Example input for Use Case #3
├── tests/
│   ├── integration_test.go # Integration tests for the application
├── Dockerfile           # Dockerfile for containerized execution
├── Makefile             # Makefile for automating common tasks
├── go.mod               # Go module file
├── go.sum               # Go dependencies checksum file
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

# Example Cases
All examples are based on cases located inside [payloads](./payloads/).

## Use Case #1
| Operation | Unit Cost | Quantity | Tax Paid | Explanation                                                                 |
|:----------|:----------|:---------|:---------|:----------------------------------------------------------------------------|
| buy       | 10.00     | 100      | 0        | Buying stocks does not incur taxes                                          |
| sell      | 15.00     | 50       | 0        | Total value is less than R$ 20,000                                          |
| sell      | 15.00     | 50       | 0        | Total value is less than R$ 20,000                                          |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 100},
{"operation":"sell", "unit-cost":15.00, "quantity": 50},
{"operation":"sell", "unit-cost":15.00, "quantity": 50}]
```
Output:
```json
[{"tax": 0},{"tax": 0},{"tax": 0}]
```

## Use Case #2
| Operation | Unit Cost | Quantity | Tax Paid | Explanation                                                                 |
|:----------|:----------|:---------|:---------|:----------------------------------------------------------------------------|
| buy       | 10.00     | 10000    | 0        | Buying stocks does not incur taxes                                          |
| sell      | 20.00     | 5000     | 10000    | Profit of R$ 50,000: 20% of the profit corresponds to R$ 10,000, no prior loss |
| sell      | 5.00      | 5000     | 0        | Loss of R$ 25,000: no taxes are paid                                       |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
{"operation":"sell", "unit-cost":20.00, "quantity": 5000},
{"operation":"sell", "unit-cost":5.00, "quantity": 5000}]
```
Output:
```json
[{"tax": 0.0},{"tax": 10000.0},{"tax": 0.0}]
```

## Use Case #3
| Operation | Unit Cost | Quantity | Tax Paid | Explanation                                                                 |
|:----------|:----------|:---------|:---------|:----------------------------------------------------------------------------|
| buy       | 10.00     | 10000    | 0        | Buying stocks does not incur taxes                                          |
| sell      | 5.00      | 5000     | 0        | Loss of R$ 25,000: no taxes are paid                                       |
| sell      | 20.00     | 3000     | 1000     | Profit of R$ 30,000: Deduct loss of R$ 25,000 and pay 20% of R$ 5,000 in taxes (R$ 1,000) |

Input:
```json
[{"operation":"buy", "unit-cost":10.00, "quantity": 10000},
{"operation":"sell", "unit-cost":5.00, "quantity": 5000},
{"operation":"sell", "unit-cost":20.00, "quantity": 3000}]
```
Output:
```json
[{"tax": 0.0},{"tax": 0.0},{"tax": 1000.0}]
```