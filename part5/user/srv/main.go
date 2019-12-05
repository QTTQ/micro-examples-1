package main

import (
	"fmt"
	"github.com/entere/micro-examples/part5/basic"
	"github.com/entere/micro-examples/part5/basic/common"
	"github.com/entere/micro-examples/part5/basic/config"
	"github.com/entere/micro-examples/part5/user/srv/handler"
	"github.com/entere/micro-examples/part5/user/srv/model"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-plugins/config/source/grpc"

	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	s "github.com/entere/micro-examples/part5/user/srv/proto/user"
)

type userCfg struct {
	common.AppCfg
}

var (
	appName = "student-srv"
	cfg     = &userCfg{}
)

func main() {
	initCfg()
	microReg := etcd.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("io.github.entere.srv.student"),
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
	// micro.RegisterSubscriber("io.github.entere.srv.student", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("io.github.entere.srv.student", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func initCfg() {
	source := grpc.NewSource(
		grpc.WithAddress("127.0.0.1:9600"),
		grpc.WithPath("micro"),
	)

	basic.Init(config.WithSource(source))

	err := config.C().App(appName, cfg)
	if err != nil {
		panic(err)
	}

	log.Logf("[initCfg] 配置，cfg：%v", cfg)

	return
}

func registryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
}
