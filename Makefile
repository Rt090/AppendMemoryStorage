.PHONY: test

test:
  go list (./...) | xargs go test
