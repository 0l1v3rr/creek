.PHONY: run
run:
	go run ./cmd/main/main.go

.PHONY: test
test:
	go test -v ./test -bench=. -run=xxx -benchmem

.PHONY: test_detailed
test_detailed:
	go test -v ./test

.PHONY: test_coverage
test_coverage:
	go test ./... -coverprofile=coverage.out

.PHONY: fmt
fmt:
	go fmt -x ./ ./test ./cmd/main

.PHONY: docs
docs:
	go doc Stream

.PHONY: vet
vet:
	go vet -x ./