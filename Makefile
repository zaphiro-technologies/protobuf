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

docker-compose:
	docker compose -f .docker/docker-compose.yml up -d

.PHONY: docker-start
docker-start: docker-compose sleep create-rabbit-defaults

.PHONY: docker-stop
docker-stop:
	docker compose -f .docker/docker-compose.yml down -v

.PHONY: sleep
sleep:
	sleep 10

.PHONY: create-rabbit-defaults
create-rabbit-defaults:
	docker exec rabbitmq ./init.sh

.PHONY: example-measurements-go
example-measurements-go:
	cd examples/go/measurements && go run main.go

.PHONY: example-faults-go
example-faults-go:
	cd examples/go/faults && go run main.go

# Mandatory
ci-test: cov
# Mandatory if benchmark enabled
ci-bench: bench
# Mandatory for docker build but can be empty
ci-pre-build:
