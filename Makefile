run:
	@echo "Starting app..."
	go run create_mr/cmd/main.go

test:
	go test ./... -v

lint:
	golangci-lint run -c .golangci.yaml

lint-fix:
	golangci-lint run -v -c .golangci.yaml --fix ./...