VERSION := "0.2"
GIT_COMMIT := $(shell git rev-parse --short HEAD)

LDFLAGS := -X ./internal/version.Vcs=$(GIT_COMMIT) -X ./internal/version.Version=$(VERSION)

build:
	go build -o bin/$(SERVICE) -ldflags "$(LDFLAGS)" main.go


run: build
	./bin/main -config config.yaml