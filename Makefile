.PHONY: build clean test lint

build:
	go build -o dxkit main.go

clean:
	rm -f dxkit

test:
	go test ./...

lint:
	golangci-lint run
