VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT  ?= $(shell git rev-parse --short HEAD)
LDFLAGS := -X github.com/muzaparoff/lorvi/cmd.Version=$(VERSION) -X github.com/muzaparoff/lorvi/cmd.Commit=$(COMMIT)

.PHONY: build test release initial-release

build:
	go build -ldflags "$(LDFLAGS)" -o lorvi .

test:
	go test -v ./...

release:
	@if [ -z "$(VERSION)" ]; then \
		echo "Usage: make release VERSION=v0.1.0"; \
		exit 1; \
	fi
	git tag -a $(VERSION) -m "Release $(VERSION)"
	git push origin $(VERSION)

# Create initial release
initial-release:
	git tag -a v0.1.0 -m "Initial release"
	git push origin v0.1.0
