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
	../proto-gen-md-diagrams/proto-gen-md-diagrams -d zaphiro -md true
	mkdir -p docs
	# mv -f zaphiro/**/*.md docs #why it works on my mac but not in the action?
	mv -f zaphiro/*/v1/*.md docs

.PHONY: proto-lint
proto-lint:
	buf lint

# Mandatory
ci-test: cov
# Mandatory if benchmark enabled
ci-bench: bench
# Mandatory for docker build but can be empty
ci-pre-build:
