package main

import (
	handler2 "github.com/entere/micro-examples/part1/hello/srv/handler"
	"github.com/micro/go-micro"

	"github.com/micro/go-micro/util/log"

	// "github.com/entere/micro-examples/hello/srv/subscriber"

	hello "github.com/entere/micro-examples/part1/hello/srv/proto/hello"
)

func main() {
	//microRegistry := etcdv3.NewRegistry(func(options *registry.Options) {
	//	options.Addrs = []string{
	//       "http://127.0.0.1:2379",
	//   }
	//})

	// New Service

	service := micro.NewService(

		micro.Name("io.github.entere.srv.hello"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler2.Hello))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("io.github.entere.srv.hello", service.Server(), new(subscriber.Hello))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("io.github.entere.srv.hello", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

//func registryOptions(ops *registry.Options) {
//
//	ops.Addrs = []string{fmt.Sprintf("%s:%d", "127.0.0.1", 2379)}
//}
