BINARY_NAME=kanban-dunkan
GOFLAGS = -ldflags="-s -w"

clean: 
	go clean
	rm -rf ./bin

build: clean
# SSH Main
	GOARCH=amd64 GOOS=darwin go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-darwin cmd/kb_ssh/main.go
	GOARCH=amd64 GOOS=linux go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-linux cmd/kb_ssh/main.go
	GOARCH=amd64 GOOS=windows go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-windows.exe cmd/kb_ssh/main.go
# Terminal
	GOARCH=amd64 GOOS=darwin go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-term-darwin cmd/kb_terminal/main.go
	GOARCH=amd64 GOOS=linux go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-term-linux cmd/kb_terminal/main.go
	GOARCH=amd64 GOOS=windows go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-term-windows.exe cmd/kb_terminal/main.go
# Desktop app
	# GOARCH=amd64 GOOS=darwin go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-win-darwin cmd/kb_win/main.go # not supported
	GOARCH=amd64 GOOS=linux go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-win-linux cmd/kb_win/main.go
	GOARCH=amd64 GOOS=windows go build $(GOFLAGS) -o ./bin/${BINARY_NAME}-win-windows.exe cmd/kb_win/main.go

run: build 
	./bin/${BINARY_NAME}-term-linux

# test:
#  go test ./...

# test_coverage:
#  go test ./... -coverprofile=coverage.out

dep:
	go mod download

assets:
	cp ./assets ./bin/

vet:
	go vet

lint:
	golangci-lint run --enable-all