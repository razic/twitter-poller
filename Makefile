all:
	GOOS=linux GOARCH=amd64 go build -o bin/twitter-poller-linux-amd64-v1.0.0 *.go
	GOOS=darwin GOARCH=amd64 go build -o bin/twitter-poller-darwin-amd64-v1.0.0 *.go
test:
	go test -v
