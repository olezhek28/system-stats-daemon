CURDIR:=`pwd`
LOCAL_BIN:=$(CURDIR)/bin
BIN_STAT := "./bin/stat"

build-stat:
	go build -v -o $(BIN_STAT) -ldflags "$(LDFLAGS)" ./cmd/system_stats_service/main.go

PHONY: generate
generate:
		mkdir -p pkg/stats_service_v1
		protoc --proto_path vendor.protogen --proto_path api/stats_service_v1 \
				--go_out=pkg/stats_service_v1 --go_opt=paths=import \
				--go-grpc_out=pkg/stats_service_v1 --go-grpc_opt=paths=import \
				--swagger_out=allow_merge=true,merge_file_name=api:pkg/stats_service_v1 \
				api/stats_service_v1/service.proto
		mv pkg/stats_service_v1/github.com/olezhek28/system-stats-daemon/pkg/stats_service_v1/* pkg/stats_service_v1/
		rm -rf pkg/stats_service_v1/github.com

PHONY: vendor-proto
vendor-proto:
		@if [ ! -d vendor.protogen/google/protobuf ]; then \
			git clone https://github.com/protocolbuffers/protobuf vendor.protogen/protobuf &&\
			mkdir -p  vendor.protogen/google/protobuf &&\
			mv vendor.protogen/protobuf/src/google/protobuf/*.proto vendor.protogen/google/protobuf &&\
			rm -rf vendor.protogen/protobuf ;\
		fi

PHONY: install-buf
install-buf:
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/bufbuild/buf/cmd/buf@v1.7.0

test:
	go test -race ./internal/...

.PHONY: test-coverage
test-coverage:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"

.PHONY: lint
lint:
	golangci-lint run --config=.golangci.yaml ./...