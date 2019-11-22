package main

import (
	"fmt"
	"github.com/entere/micro-examples/part5/basic"
	"github.com/entere/micro-examples/part5/basic/config"
	"github.com/entere/micro-examples/part5/user/srv/handler"
	"github.com/entere/micro-examples/part5/user/srv/model"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	s "github.com/entere/micro-examples/part5/user/srv/proto/user"
)

func main() {
	basic.Init()
	microReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("io.github.entere.srv.user"),
		micro.Version("latest"),
		micro.Registry(microReg),
	)

	// Initialise service
	service.Init(micro.Action(func(context *cli.Context) {
		model.Init()
		handler.Init()
	}))

	// Register Handler
	s.RegisterUserHandler(service.Server(), new(handler.Service))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("io.github.entere.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("io.github.entere.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {

	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}

}
