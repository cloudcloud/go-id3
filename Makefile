
install:
	go get ./...

test:
	go test ./...

coverage: INT?=0
coverage: OUT?=../coverage
coverage:
	@if [ -f "$(OUT).json" ]; then rm $(OUT).json; fi
	@if [ -f "$(OUT).html" ]; then rm $(OUT).html; fi
	RUNINTEGRATION=$(INT) gocov test `go list ./... | grep -v /vendor/` > $(OUT).json
	gocov-html $(OUT).json > $(OUT).html
	go test -coverprofile=$(OUT).out -covermode=count -coverpkg "./..."
