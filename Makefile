default: lint test generate

lint:
	golangci-lint run ./...

test:
	go test ./...

# Generate copywrite headers
generate:
	cd tools; go generate ./...

.PHONY: default lint test
