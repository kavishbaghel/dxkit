.PHONY: build clean test lint

build:
	go build -o dist/ main.go

clean:
	rm -f dxkit

test:
	go test ./...

lint:
	golangci-lint run
