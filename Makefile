C = $(shell printf "\033[35;1m-->\033[0m")
V := $(if $V,,@)
VERSION := $(shell git rev-parse HEAD)

build: ; $(info $(C) building binary...)
	$(V) go build -v -ldflags="-X 'github.com/cloudcloud/go-id3/internal/version.buildNumber=$(VERSION)' -X 'github.com/cloudcloud/go-id3/internal/version.buildTime=$(shell date)'" ./cmd/go-id3

test: ; $(info $(C) running tests with race...)
	$(V) go test -race ./...

coverage: ; $(info $(C) generating coverage with race and outputting coverage.html...)
	$(V) go test -race -coverprofile=/tmp/cov ./... && go tool cover -html=/tmp/cov -o ./coverage.html

.PHONY: coverage.html
