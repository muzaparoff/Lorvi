VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT  ?= $(shell git rev-parse --short HEAD)
LDFLAGS := -X github.com/muzaparoff/lorvi/cmd.Version=$(VERSION) -X github.com/muzaparoff/lorvi/cmd.Commit=$(COMMIT)

.PHONY: build test release initial-release setup-test-env clean-test-env setup-monitoring grafana-dashboard clean-monitoring install

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

setup-test-env:
	chmod +x scripts/*.sh
	./scripts/setup_test_env.sh

clean-test-env:
	multipass delete k3s --purge
	pkill ollama || true
	rm -f ~/.kube/k3s-config

setup-monitoring:
	./scripts/install_monitoring.sh

grafana-dashboard:
	kubectl port-forward -n monitoring svc/monitoring-grafana 3000:80

clean-monitoring:
	helm uninstall -n monitoring monitoring loki
	kubectl delete namespace monitoring

install:
	chmod +x scripts/install_lorvi.sh
	./scripts/install_lorvi.sh
