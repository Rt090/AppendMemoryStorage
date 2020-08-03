.PHONY: test

test:
	@go test -cover $(shell go list ./...)
