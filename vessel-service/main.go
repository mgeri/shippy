package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/mgeri/shippy/vessel-service/proto/vessel"

	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "mongodb://localhost:27017"
)

func createDummyData(repo *VesselRepository) {
	vessels := []*Vessel{
		{ID: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		_ = repo.Create(context.Background(), v)
	}
}

func main() {

	log.Println("starting vessel service...")

	// Create a new service. Optionally include some options here.
	srv := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("shippy.vessel.service"),
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

	vesselCollection := client.Database("shippy").Collection("vessel")
	repository := &VesselRepository{vesselCollection}

	// Create dummy vessel
	createDummyData(repository)

	// Init service handler
	sh := &handler{repository}

	// Register our implementation with
	_ = pb.RegisterVesselServiceHandler(srv.Server(), sh)

	// Run the server
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
