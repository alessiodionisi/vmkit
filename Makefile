.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	--go-grpc_out=. --go-grpc_opt=paths=source_relative \
	proto/service.proto

.PHONY: build
build:
	goreleaser build --rm-dist

.PHONY: build-snapshot
build-snapshot:
	goreleaser build --rm-dist --snapshot

.PHONY: release
release:
	goreleaser release --rm-dist

.PHONY: release-snapshot
release-snapshot:
	goreleaser release --rm-dist --snapshot
