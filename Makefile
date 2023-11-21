VERSION ?= latest

# Create a variable that contains the current date in UTC
# Different flow if this script is running on Darwin or Linux machines.
ifeq (Darwin,$(shell uname))
	NOW_DATE = $(shell date -u +%d-%m-%Y)
else
	NOW_DATE = $(shell date -u -I)
endif

all: test

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
	cd proto; buf generate

.PHONY: docs
docs:
	../proto-gen-md-diagrams/proto-gen-md-diagrams -d proto
	mkdir -p docs
	# mv -f proto/**/*.md docs #why it works on my mac but not in the action?
	mv -f proto/*/v1/*.md docs

.PHONY: proto-lint
proto-lint:
	cd proto; buf lint

# Mandatory
ci-test: cov
# Mandatory if benchmark enabled
ci-bench: bench
# Mandatory for docker build but can be empty
ci-pre-build:
