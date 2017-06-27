all:
	go build -o bin/awsreset cmd/awsreset/*.go
test:
	go test -v cmd/awsreset/*.go
