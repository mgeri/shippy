package main

import (
	"context"
	"fmt"
	"log"
	"os"
	// Import the generated protobuf code
	pb "github.com/mgeri/shippy/consignment-service/proto/consignment"
	vesselProto "github.com/mgeri/shippy/vessel-service/proto/vessel"

	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func main() {

	log.Println("starting consignment service...")

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.consignment.service"),
	)

	// Init will parse the command line flags.
	srv.Init()

	// Init Mongo data store
	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	consignmentCollection := client.Database("shippy").Collection("consignment")
	repository := &ConsignmentRepository{consignmentCollection}

	// Create vessel service client
	vesselService := vesselProto.NewVesselService("shippy.vessel.service", srv.Client())

	// Init service handler
	sh := &handler{repository, vesselService}
	// Register handler
	_ = pb.RegisterShippingServiceHandler(srv.Server(), sh)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
