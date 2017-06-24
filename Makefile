all:
	CGO_ENABLED=0 GOOS=linux go build -a -o bin/awsreset -tags netgo -ldflags '-w' cmd/awsreset/*.go
