VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT  ?= $(shell git rev-parse --short HEAD)
LDFLAGS := -X github.com/muzaparoff/lorvi/cmd.Version=$(VERSION) -X github.com/muzaparoff/lorvi/cmd.Commit=$(COMMIT)

.PHONY: build test release

build:
	go build -ldflags "$(LDFLAGS)" -o lorvi .

test:
	go test -v ./...

release:
	git tag -a v0.1.0 -m "Initial release"
	git push origin v0.1.0
