.PHONY: clean test

lint:
	go fmt ./...

test:
	go test -v ./...

coverage:
	go test -v -cover ./...
