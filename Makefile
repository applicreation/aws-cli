GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

clean:
	rm -rf ./build

build: clean
	go build -ldflags="-s -w" -o ./build/aws-cli-$(GOOS)-$(GOARCH)