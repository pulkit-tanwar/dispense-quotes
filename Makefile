COVERPROFILE=.cover.out
COVERDIR=.cover
PORT ?= 3000

run:
	@go run main.go dispense

run-server:
	@go run main.go serve


dep:
	@go get ./...

test:
	@go test -coverprofile=$(COVERPROFILE) ./...

tidy:
	@go mod tidy

cover: test
	@mkdir -p $(COVERDIR)
	@go tool cover -html=$(COVERPROFILE) -o $(COVERDIR)/index.html
	@cd $(COVERDIR) && python3 -m http.server $(PORT)


.PHONY: run dep test tidy