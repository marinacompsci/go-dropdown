BIN_PATH="bin/app"

SOURCES=$(shell find . -name "*.go" -type f)

build: $(SOURCES)
	go build -o $(BIN_PATH)

