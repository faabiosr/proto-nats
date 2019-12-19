.DEFAULT_GOAL := test

# Build (dev only)
build:
	@CGO_ENABLE=0 go build -o ./build/protoc-gen-nats cmd/protoc-gen-nats/*.go
.PHONY: build

# Clean up
clean:
	@rm -fR ./cover.* ./build
.PHONY: clean

# Run tests and generates html coverage file
cover: test
	@go tool cover -html=./coverage.text -o ./cover.html
	@test -f ./coverage.text && rm ./coverage.text;
.PHONY: cover

# Format all go files
fmt:
	@gofmt -s -w -l $(shell go list -f {{.Dir}} ./...)
.PHONY: fmt

# Run pkger and embed the asserts
pack:
	@pkger -o cmd/protoc-gen-nats
.PHONY: pack

# Run tests
test:
	@go test -v -race -coverprofile=./coverage.text -covermode=atomic $(shell go list ./...)
.PHONY: test
