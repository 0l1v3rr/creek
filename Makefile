.PHONY: run
run:
	@go run ./cmd/main/main.go

.PHONY: test
test:
	@go test -v ./test -bench=. -run=xxx -benchmem

.PHONY: test-detailed
test-detailed:
	@go test -v ./test

.PHONY: fmt
fmt:
	@go fmt