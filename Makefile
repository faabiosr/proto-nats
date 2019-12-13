.DEFAULT_GOAL := test

# Clean up
clean:
	@rm -fR ./cover.*
.PHONY: clean

# Run tests and generates html coverage file
cover: test
	@go tool cover -html=./cover.text -o ./cover.html
	@test -f ./cover.text && rm ./cover.text;
.PHONY: cover

# Format all go files
fmt:
	@gofmt -s -w -l $(shell go list -f {{.Dir}} ./...)
.PHONY: fmt

# Run tests
test:
	@go test -v -race -coverprofile=./cover.text -covermode=atomic $(shell go list ./...)
.PHONY: test
