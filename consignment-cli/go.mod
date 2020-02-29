module github.com/mgeri/shippy/consignment-cli

go 1.13

replace github.com/coreos/go-systemd => github.com/coreos/go-systemd/v22 v22.0.0

require (
	github.com/mgeri/shippy/consignment-service v0.0.0-20200223173552-e3659c220068
	github.com/micro/go-micro/v2 v2.2.0
)
