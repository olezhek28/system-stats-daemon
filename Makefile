CURDIR:=`pwd`
LOCAL_BIN:=$(CURDIR)/bin

PHONY: generate
generate:
	$(LOCAL_BIN)/buf generate

PHONY: install-buf
install-buf:
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/bufbuild/buf/cmd/buf@v1.7.0

test:
	go test -race ./internal/...

.PHONY: test-coverage
test-coverage:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"