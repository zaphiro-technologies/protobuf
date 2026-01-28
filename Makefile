# Copyright 2024 Zaphiro Technologies
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

PYPROJECT_FILE ?= pyproject.toml

.PHONY: install-proto-gen-md
install-proto-gen-md:
	@echo "Installing latest proto-gen-md-diagrams version..."
	@rm -rf /tmp/proto-gen-md-diagrams
	@git clone --depth 1 https://github.com/GoogleCloudPlatform/proto-gen-md-diagrams /tmp/proto-gen-md-diagrams
	@cd /tmp/proto-gen-md-diagrams && \
		go build -o proto-gen-md-diagrams . && \
		mkdir -p $(CURDIR)/bin && \
		cp -f proto-gen-md-diagrams $(CURDIR)/bin/proto-gen-md-diagrams && \
		echo "Installed to $(CURDIR)/bin/proto-gen-md-diagrams"

.PHONY: install-gomarkdoc
install-gomarkdoc:
	@echo "Installing latest gomarkdoc version..."
	go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@latest

.PHONY: install
install: install-proto-gen-md install-gomarkdoc

all: proto-lint install generate lint test docs

.PHONY: lint
lint:
	golangci-lint run --fix && golines . -w

.PHONY: test
test:
	go test ./...

.PHONY: bench
bench:
	go test ./... -benchmem -p 1 -run='^$$' -bench=.

.PHONY: cov
cov:
	mkdir -p coverage && \
	go test ./... -coverpkg=./... -coverprofile=./coverage/coverage.out && \
	gcov2lcov -infile=./coverage/coverage.out -outfile=./coverage/lcov.info

.PHONY: generate
generate:
	buf generate

.PHONY: docs
docs:
	mkdir -p docs
	bin/proto-gen-md-diagrams -d zaphiro -o docs -md true
	gomarkdoc ./go/constants > docs/constants.md

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

.PHONY: example-measurements-python
example-measurements-python:
	cd examples/python && poetry run python measurements/main.py

.PHONY: example-faults-python
example-faults-python:
	cd examples/python && poetry run python faults/main.py

.PHONY: release
release: validate-tag check-version update-pyproject
	@echo "✅ Version update completed successfully!"

# Validate that TAG parameter is provided and follows version format
validate-tag:
	@if [ -z "$(TAG)" ]; then \
		echo "❌ Error: TAG parameter is required"; \
		echo "Usage: make update-version TAG=1.2.3"; \
		exit 1; \
	fi
	@echo "Validating tag format: $(TAG)"
	@if ! echo "$(TAG)" | grep -E "^[0-9]+\.[0-9]+\.[0-9]+.*$$" > /dev/null; then \
		echo "❌ Tag $(TAG) does not follow version format (*.*.*)"; \
		echo "Skipping version file update."; \
		exit 1; \
	fi
	@echo "Valid version tag detected: $(TAG)"


# Update version in pyproject.toml
update-pyproject: validate-tag
	@echo "Updating version in $(PYPROJECT_FILE)"
	@if [ ! -f "$(PYPROJECT_FILE)" ]; then \
		echo "⚠️  Warning: $(PYPROJECT_FILE) not found, skipping pyproject.toml update"; \
	else \
		sed -i 's/version = ".*"/version = "$(TAG)"/g' $(PYPROJECT_FILE); \
		echo "✅ Updated $(PYPROJECT_FILE) with version $(TAG)"; \
		grep "version = " $(PYPROJECT_FILE) || true; \
	fi

# Check current versions
check-version:
	@echo "🔍 Current version information:"
	@if [ -f "$(PYPROJECT_FILE)" ]; then \
		echo "$(PYPROJECT_FILE):"; \
		grep "version = " $(PYPROJECT_FILE) || echo "No version found in $(PYPROJECT_FILE)"; \
	else \
		echo "$(PYPROJECT_FILE): not found"; \
	fi

# Mandatory
ci-test: cov
# Mandatory if benchmark enabled
ci-bench: bench
# Mandatory for docker build but can be empty
ci-pre-build:
