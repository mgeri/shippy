package main

import (
	"log"

	"github.com/golang/protobuf/proto"

	pb "github.com/mgeri/shippy/user-service/proto/user"
	micro "github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"

	_ "github.com/micro/go-plugins/broker/nats/v2"
)

const topic = "user.created"

func main() {
	srv := micro.NewService(
		micro.Name("shippy.email.service"),
		micro.Version("latest"),
	)

	srv.Init()

	// Get the broker instance using our environment variables
	pubSub := srv.Server().Options().Broker
	if err := pubSub.Connect(); err != nil {
		log.Fatal(err)
	}

	// Subscribe to messages on the broker
	_, err := pubSub.Subscribe(topic, func(e broker.Event) error {
		user := &pb.User{}
		if err := proto.Unmarshal(e.Message().Body, user); err != nil {
			return err
		}
		log.Println(user)
		go sendEmail(user)
		return nil
	})

	if err != nil {
		log.Println(err)
	}

	// Run the server
	if err := srv.Run(); err != nil {
		log.Println(err)
	}
}

func sendEmail(user *pb.User) error {
	log.Println("Sending email to:", user.Name)
	return nil
}
