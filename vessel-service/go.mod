module github.com/mgeri/shippy/vessel-service

go 1.13

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

require (
	github.com/golang/protobuf v1.3.2
	github.com/mgeri/shippy/consignment-service v0.0.0-20200223173552-e3659c220068
	github.com/micro/go-micro/v2 v2.2.0
)
