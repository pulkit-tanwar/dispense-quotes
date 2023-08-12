# Programming quotes despenser service

A Go microservice that will dispense programming wisdom quotes

## Usage

### To Run the application
```bash
  $ go run main.go dispense
```

### To Run the HTTP server
```bash
  $ go run main.go serve
```

### To Run unit tests
```bash
  $ go test ./...
```

## Development
```bash
  $ make dep           # install dependencies
  $ make run-server    # run the HTTP server
  $ make tidy          # removes any unused dependencies and adds any missing dependencies to the go.mod  file.
  $ make test          # run unit tests
  $ make cover         # run code coverage report service (http://localhost:3000)
  $ make run           # run the service
