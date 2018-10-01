
install:
	go get ./...

test:
	go test ./...

coverage:
	go test -race -coverprofile=/tmp/cov ./... && go tool cover -html=/tmp/cov -o ./coverage.html

.PHONY: coverage.html
