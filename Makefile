test: lint
	go test -coverprofile=coverage.out ./...

lint:
	go fmt ./...
	golangci-lint run ./... --verbose 