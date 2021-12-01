.PHONY: test vet lint ci

test:
	go test -cover ./...

vet:
	go vet ./...

lint:
	golint -set_exit_status ./...
	golangci-lint run

ci: vet lint test
