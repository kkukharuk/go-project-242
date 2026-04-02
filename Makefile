build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run: build
	./bin/hexlet-path-size

test:
	go test -v ./...

lint:
	golangci-lint run

lint-fix:
	golangci-lint run --fix

clean:
	rm -rf bin/*

.PHONY: build run test lint lint-fix clean
