all: gotool
	@go build -v
gotool:
	gofmt -w .
	go tool vet .
.PHONY: clean gotool


