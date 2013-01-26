export GOPATH=$(shell pwd)
export GO=go
export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig/:/opt/X11/lib/pkgconfig/

all: *.go
	$(GO) build
	$(GO) vet
	$(GO) fmt
	$(GO) test
