all: proto-lint generate lint test docs

.PHONY: lint
lint:
	golangci-lint run --fix

.PHONY: test
test:
	go test ./...

.PHONY: bench
bench:
	go test ./... -bench=. -run=^$

.PHONY: cov
cov:
	mkdir -p coverage && \
	go test ./... -coverprofile=./coverage/coverage.out && \
	gcov2lcov -infile=./coverage/coverage.out -outfile=./coverage/lcov.info

.PHONY: generate
generate:
	buf generate

.PHONY: docs
docs:
	mkdir -p docs
	../proto-gen-md-diagrams/proto-gen-md-diagrams -d zaphiro -o docs -md true

.PHONY: proto-lint
proto-lint:
	buf lint

# Mandatory
ci-test: cov
# Mandatory if benchmark enabled
ci-bench: bench
# Mandatory for docker build but can be empty
ci-pre-build:
