GOOS        ?= $(shell go env GOOS)
GOARCH      ?= $(shell go env GOARCH)
GOPATH      ?= $(shell go env GOPATH)
# Disable cgo by default.
CGO_ENABLED ?= 0

temporal-server:
	@printf $(COLOR) "Build temporal-server with CGO_ENABLED=$(CGO_ENABLED) for $(GOOS)/$(GOARCH)..."
	CGO_ENABLED=$(CGO_ENABLED) go build -o ./bin/temporal-server .

container:
	@printf $(COLOR) "Build temporal-server container using containerd..."
	nerdctl --namespace k8s.io build -t temporal-distribution/server:1.18.0 .

run-dev-migrations:
	./dev-migrations.sh