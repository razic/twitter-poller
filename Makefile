all:
	go build -o bin/twitter-poller *.go
test:
	go test -v
