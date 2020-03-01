.PHONY: build proto docker-build docker-run
build:
	@echo "Run 'make proto' to generate protobuf files."
	@echo "Before to run make, be sure proitoc, go protobuf and micro service protobuf plugin is installaed as following:"
	@echo "  go get -u github.com/micro/protobuf/{proto,protoc-gen-go}"
proto:
	protoc -I. --micro_out=. --go_out=. ./proto/vessel/vessel.proto
docker-build:
	docker build -t shippy-vessel-service .
docker-run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 --name shippy-vessel-service shippy-vessel-service
