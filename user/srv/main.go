package main

import (
    "github.com/entere/micro-examples/user/srv/handler"
    "github.com/micro/go-micro/util/log"

    user "github.com/entere/micro-examples/user/srv/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("io.github.entere.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
