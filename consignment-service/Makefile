.PHONY: build proto
build:
	@echo "Run 'make proto' to generate protobuf files."
	@echo "Before to run make, be sure proitoc, go protobuf and micro service protobuf plugin is installaed as following:"
	@echo "  go get -u github.com/micro/protobuf/{proto,protoc-gen-go}"
proto:
	protoc -I. --micro_out=. --go_out=. ./proto/consignment/consignment.proto