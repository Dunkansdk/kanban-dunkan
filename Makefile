BINARY_NAME=kanban-dunkan
VERSION=$(shell git describe --tags --always --long --dirty)
GOFLAGS = -ldflags="-s -w -X main.version=$(VERSION)"

all: build

dep:
	go mod download

build: clean
# SSH Main
	GOARCH=amd64 GOOS=darwin go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-darwin cmd/ssh/main.go
	GOARCH=amd64 GOOS=linux go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-linux cmd/ssh/main.go
	GOARCH=amd64 GOOS=windows go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-windows.exe cmd/ssh/main.go
# Terminal
	GOARCH=amd64 GOOS=darwin go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-term-darwin cmd/term/main.go
	GOARCH=amd64 GOOS=linux go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-term-linux cmd/term/main.go
	GOARCH=amd64 GOOS=windows go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-term-windows.exe cmd/term/main.go
# Desktop app
	# GOARCH=amd64 GOOS=darwin go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-win-darwin cmd/desktop/main.go # not supported
	GOARCH=amd64 GOOS=linux go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-win-linux cmd/desktop/main.go
	GOARCH=amd64 GOOS=windows go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-win-windows.exe cmd/desktop/main.go
	@echo version: $(VERSION)

run: build 
	./bin/${BINARY_NAME}-term-linux

clean: 
	go clean
	rm -rf ./bin

# Live Reload
watch:
	@if command -v air > /dev/null; then \
	    air; \
	    echo "Watching...";\
	else \
	    read -p "Go's 'air' is not installed on your machine. Do you want to install it? [Y/n] " choice; \
	    if [ "$$choice" != "n" ] && [ "$$choice" != "N" ]; then \
	        go install github.com/air-verse/air@latest; \
			export GOROOT=$GOPATH; \
	        air; \
	        echo "Watching...";\
	    else \
	        echo "You chose not to install air. Exiting..."; \
	        exit 1; \
	    fi; \
	fi

.PHONY: all build run clean
