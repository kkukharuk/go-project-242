build:
	go build -o bin/hexlet-path-size ./cmd/hexlet-path-size

run: build
	./bin/hexlet-path-size

test:
	go test ./...

clean:
	rm -rf bin/*

.PHONY: build run test clean
