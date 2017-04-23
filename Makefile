GOOS=linux
GOARCH=amd64
CGO_ENABLED=0

default: build

build: plugin test

plugin:
	go build -buildmode=plugin -o plugins/animal_sounds.so plugins/animal.go
	go build -buildmode=plugin -o plugins/machine_sounds.so plugins/machine.go

test:
	go run main.go
