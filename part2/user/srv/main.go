package main

import (
	handler2 "github.com/entere/micro-examples/part2/user/srv/handler"
	"github.com/entere/micro-examples/part2/user/srv/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
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
	io_github_entere_srv_user.RegisterUserHandler(service.Server(), new(handler2.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
