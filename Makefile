test:
	@echo "Testing"
	@go test -p 1 ./...

vet:
	@echo "Vetting"
	@go vet ./...

run:
	@echo "Running"
	@go run ./...

build:
	@echo "Building"
	@go build ./...

cover:
	@echo "testing with coverage"
	@go test -p 1 -cover ./...