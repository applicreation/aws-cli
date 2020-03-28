NAME=aws-cli
VERSION=0.4.0

GOOS=$(shell go env GOOS)
GOARCH=$(shell go env GOARCH)

clean:
	rm -rf ./build

build: clean
	go build -ldflags="-s -w" -o ./build/$(GOOS)/$(GOARCH)/$(NAME)

compress:
	upx --ultra-brute ./build/$(GOOS)/$(GOARCH)/$(NAME)

package: build compress
	cd ./build/$(GOOS)/$(GOARCH) && tar -czf ./$(NAME)-$(GOOS)-$(GOARCH)-$(VERSION).tar.gz ./$(NAME)
	cd ./build/$(GOOS)/$(GOARCH) && sha256sum ./$(NAME)-$(GOOS)-$(GOARCH)-$(VERSION).tar.gz
