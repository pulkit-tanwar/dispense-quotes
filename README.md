# Programming quotes despenser service

A Go microservice that will dispense programming wisdom quotes

## Usage

### To Run the application
```bash
  $ go run main.go dispense
```

### To Run unit tests
```bash
  $ go test ./...
```

## Development
```bash
  $ make dep           # install dependencies
  $ make test          # run unit tests
  $ make cover         # run code coverage report service (http://localhost:3000)
  $ make run           # run the service
