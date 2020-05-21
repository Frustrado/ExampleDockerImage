.PHONY: build
build:
	CGO_ENABLED=0 GOOS=linux go build -o bin/auiapp -v main/*.go
modules:
	go mod download