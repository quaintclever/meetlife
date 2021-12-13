SERVICE ?= meetlife-mosch

.PHONY: vendor
vendor:
	go mod tidy -v -compat=1.17
	go mod vendor -v

.PHONY: darwin
build: vendor
	go build -mod=vendor -trimpath -o bin/$(SERVICE)

.PHONY: run
run: build
	bin/$(SERVICE) config/config-local.ini

