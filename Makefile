CURDIR:=`pwd`
LOCAL_BIN:=$(CURDIR)/bin

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

PHONY: install-buf
install-buf:
	GOBIN=$(LOCAL_BIN) go install -mod=mod github.com/bufbuild/buf/cmd/buf@v1.7.0

test:
	go test -race ./internal/...

.PHONY: test-coverage
test-coverage:
	go test -race -coverprofile="coverage.out" -covermode=atomic ./...
	go tool cover -html="coverage.out"