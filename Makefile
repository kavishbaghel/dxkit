.PHONY: build clean test lint

build:
	go build -o dist/ cmd/

clean:
	rm -f dxkit

test:
	go test ./...

lint:
	golangci-lint run
