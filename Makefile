GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

default: build

build: plugin test

plugin:
	go build -buildmode=plugin -o plugins/dog.so plugins/dog.go

test:
	go run main.go
