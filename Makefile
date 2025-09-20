PROTO_DIRS := proto/v1/simple
PATCH_PATH := $(shell go env GOPATH)/src/github.com/alta/protopatch/patch
PROTO_FILES := $(foreach dir,$(PROTO_DIRS),$(wildcard $(dir)/*.proto))

generate: check-plugins
	@echo "==> Generating Go structs"
	protoc \
	  --proto_path=. \
	  --proto_path=vendor.protogen \
	  --proto_path=$(PATCH_PATH) \
	  --go_out=paths=source_relative:. \
	  $(PROTO_FILES)

	@echo "==> Generating gRPC services"
	protoc \
	  --proto_path=. \
	  --proto_path=vendor.protogen \
	  --proto_path=$(PATCH_PATH) \
	  --go-grpc_out=paths=source_relative:. \
	  $(PROTO_FILES)

	@echo "==> Applying go-patch (custom tags)"
	protoc \
	  --proto_path=. \
	  --proto_path=vendor.protogen \
	  --proto_path=$(PATCH_PATH) \
	  --go-patch_out=plugin=go,paths=source_relative:. \
	  $(PROTO_FILES)

clean:
	@echo "==> Cleaning generated files"
	find $(PROTO_DIRS) -type f \( -name "*.pb.go" -o -name "*.grpc.pb.go" \) -delete

.PHONY: generate check-plugins clean
