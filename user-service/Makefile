.PHONY: build proto docker-build docker-run
build:
	@echo "Run 'make proto' to generate protobuf files."
	@echo "Before to run make, be sure protoc, go protobuf and micro service protobuf plugin is installed as following:"
	@echo "  go get -u github.com/micro/protobuf/{proto,protoc-gen-go}"
proto:
	protoc -I. --micro_out=. --go_out=. ./proto/user/user.proto
docker-build:
	docker build -t shippy-user-service .
docker-run:
	docker run -p 50052:50051 -e MICRO_SERVER_ADDRESS=:50051 --name shippy-user-service shippy-user-service
