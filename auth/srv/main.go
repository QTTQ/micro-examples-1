package main

import (
	"github.com/entere/micro-examples/auth/srv/handler"
	"github.com/entere/micro-examples/auth/srv/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	"github.com/entere/micro-examples/auth/srv/proto/auth"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("io.github.entere.srv.auth"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	auth.RegisterAuthHandler(service.Server(), new(handler.Auth))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("io.github.entere.srv.auth", service.Server(), new(subscriber.Auth))

	// Register Function as Subscriber
	micro.RegisterSubscriber("io.github.entere.srv.auth", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
