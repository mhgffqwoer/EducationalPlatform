.SILENT: build run test clean

BIN_NAME = app

build:
	go build -o ./bin/$(BIN_NAME) ./cmd/main.go

run: build
	./bin/$(BIN_NAME)

test:
	go test -v ./tests/...

clean:
	rm -rf bin
