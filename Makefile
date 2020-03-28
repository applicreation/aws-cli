GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

clean:
	rm -rf ./bin/*

build: clean
	go build -ldflags="-s -w" -o ./bin/aws-cli-$(GOOS)-$(GOARCH)