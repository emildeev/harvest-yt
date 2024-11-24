
LINTBIN := $(CURBINDIR)/golangci-lint
LINT_VERSION := v1.61.0

run:
	@echo "Starting app..."
	go run create_mr/cmd/main.go

test:
	go test ./... -v

lint:
	$(LINTBIN) run -c .golangci.yaml

lint-fix:
	$(LINTBIN) run -v -c .golangci.yaml origin/master  --fix ./...


PHONY: .install-linter
.install-linter:
	@if ! [ -x "$(LINTBIN)" ] || [ "$$($(LINTBIN) --version | grep -oE 'v[0-9]+\.[0-9]+\.[0-9]+')" != $(LIN_VERSION) ]; then \
  		echo "Installing golangci-lint..."; \
    	GOBIN=$(CURBINDIR) go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(LIN_VERSION); \
    else \
    	echo "golangci-lint is already installed and up-to-date"; \
    fi
