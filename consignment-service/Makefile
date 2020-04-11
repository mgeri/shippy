.PHONY: build proto docker-build docker-run
build:
	@echo "Run 'make proto' to generate protobuf files."
	@echo "Before to run make, be sure proitoc, go protobuf and micro service protobuf plugin is installaed as following:"
	@echo "  go get -u github.com/micro/protobuf/{proto,protoc-gen-go}"
	@echo "  go get -u github.com/micro/protoc-gen-micro/v2"
proto:
	protoc -I. --micro_out=. --go_out=. ./proto/consignment/consignment.proto
docker-build:
	docker build -t shippy-consignment-service .
docker-run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 --name shippy-consignment-service shippy-consignment-service
