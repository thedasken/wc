run:
	@go run .

test:
	@go test -v ./...

build:
	@go build .

clean:
	@rm -f wc

.PHONY: run test build clean