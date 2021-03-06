all:
	GOOS=linux GOARCH=amd64 go build -o bin/awsreset-linux-amd64-v1.0.0 cmd/awsreset/*.go
	GOOS=darwin GOARCH=amd64 go build -o bin/awsreset-darwin-amd64-v1.0.0 cmd/awsreset/*.go
test:
	go test -v cmd/awsreset/*.go
