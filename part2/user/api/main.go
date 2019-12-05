package main

import (
	client2 "github.com/entere/micro-examples/part2/user/api/client"
	handler2 "github.com/entere/micro-examples/part2/user/api/handler"
	"github.com/entere/micro-examples/part2/user/api/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("io.github.entere.api.student"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the User srv client
		micro.WrapHandler(client2.UserWrapper(service)),
	)

	// Register Handler
	io_github_entere_api_user.RegisterUserHandler(service.Server(), new(handler2.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
