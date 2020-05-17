module github.com/mgeri/shippy/consignment-service

go 1.14

// replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

require (
	github.com/golang/protobuf v1.3.5
	github.com/mgeri/shippy/vessel-service v0.0.0-20200229163553-864357e47907
	github.com/micro/go-micro/v2 v2.6.0
	github.com/pkg/errors v0.9.1
	go.mongodb.org/mongo-driver v1.3.3
)
