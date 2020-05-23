package main

import (
	"log"
	"os"

	"context"
	pb "github.com/mgeri/shippy/user-service/proto/user"

	micro "github.com/micro/go-micro/v2"
)

func main() {

	// Create a new service
	service := micro.NewService(micro.Name("shippy.user.cli"))
	// Initialise the client and parse command line flags
	service.Init()

	// Create new greeter client
	client := pb.NewUserService("shippy.user.service", service.Client())

	name := "Marco Geri"
	email := "marco.geri.pi@gmail.com"
	password := "test123"
	company := "MyOwnCompany"

	log.Println(name, email, password)

	r, err := client.Create(context.Background(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := client.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := client.Auth(context.Background(), &pb.User{
		Email:    email,
		Password: password,
	})

	if err != nil {
		log.Fatalf("Could not authenticate user: %s error: %v\n", email, err)
	}

	log.Printf("Your access token is: %s \n", authResponse.Token)

	// let's just exit because
	os.Exit(0)
}
