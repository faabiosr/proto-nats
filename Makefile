.DEFAULT_GOAL := test

# Build builds the app (only for testing purpose)
build:
	@go build -v -o ./build/protoc-gen-nats ./cmd/protoc-gen-nats
.PHONY: build

# Clean up
clean:
	@rm -fR ./cover.* ./build/ ./dist/
.PHONY: clean

# Creates folders and download dependencies
configure:
	@mkdir -p ./build
	@go mod download
.PHONY: configure

# Run tests and generates html coverage file
cover: test
	@go tool cover -html=./cover.out -o ./cover.html
	@rm ./cover.out
.PHONY: cover

# GolangCI Linter
lint:
	@golangci-lint run -v ./...
.PHONY: lint

# Run pkger and embed the asserts
pack:
	@pkger -o cmd/protoc-gen-nats
.PHONY: pack

# Release runs the goreleaser and creates the release for local testing.
release:
	@goreleaser release --snapshot
.PHONY: release

# Run tests
test:
	@go test -v -covermode=atomic -coverprofile=cover.out ./...
.PHONY: test
