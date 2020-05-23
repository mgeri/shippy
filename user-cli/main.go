package main

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/config/cmd"
	"log"
	"os"

	"context"
	pb "github.com/mgeri/shippy/user-service/proto/user"
)

func main() {

	// Micro command line tool init (read gomicro env)
	// Note: when it's the same of srv.Init() but for command cli
	cmd.Init()

	// Create new greeter client
	userSrvClient := pb.NewUserService("shippy.user.service", client.DefaultClient)

	name := "Marco Geri"
	email := "marco.geri.pi@gmail.com"
	password := "test123"
	company := "MyOwnCompany"

	log.Println(name, email, password)

	r, err := userSrvClient.Create(context.Background(), &pb.User{
		Name:     name,
		Email:    email,
		Password: password,
		Company:  company,
	})
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}
	log.Printf("Created: %s", r.User.Id)

	getAll, err := userSrvClient.GetAll(context.Background(), &pb.Request{})
	if err != nil {
		log.Fatalf("Could not list users: %v", err)
	}
	for _, v := range getAll.Users {
		log.Println(v)
	}

	authResponse, err := userSrvClient.Auth(context.Background(), &pb.User{
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
