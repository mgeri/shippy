package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/metadata"
	"github.com/micro/go-micro/v2/server"
	"github.com/pkg/errors"
	"log"
	"os"
	// Import the generated protobuf code
	pb "github.com/mgeri/shippy/consignment-service/proto/consignment"
	userProto "github.com/mgeri/shippy/user-service/proto/user"
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
		micro.Version("latest"),
		micro.WrapHandler(AuthWrapper), // JWT auth management
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

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) error {

		// This skips our auth check if DISABLE_AUTH is set to true
		if os.Getenv("DISABLE_AUTH") == "true" {
			return fn(ctx, req, resp)
		}

		// Get JWT token from metadata
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		token := meta["Token"]
		log.Println("Authenticating with token: ", token)

		// Validate token using the user service
		authClient := userProto.NewUserService("shippy.user.service", client.DefaultClient)
		authResp, err := authClient.ValidateToken(ctx, &userProto.Token{
			Token: token,
		})
		log.Println("Auth resp:", authResp)
		log.Println("Err:", err)
		if err != nil {
			return err
		}

		err = fn(ctx, req, resp)
		return err
	}
}
