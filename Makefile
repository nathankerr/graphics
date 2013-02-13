export GO=go

all: *.go
	$(GO) build
	$(GO) vet
	$(GO) fmt
	$(GO) test

clean:
	- rm test.*