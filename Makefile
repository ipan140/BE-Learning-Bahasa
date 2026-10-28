.PHONY: run build build-linux tidy clean test

run:
	go run cmd/api/main.go

build:
	go build -o server.exe cmd/api/main.go

build-linux:
	GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o server cmd/api/main.go

tidy:
	go mod tidy

clean:
	rm -f server server.exe
