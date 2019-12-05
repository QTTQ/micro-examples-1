package main

import (
	"fmt"
	"github.com/entere/micro-examples/part6/basic"
	"github.com/entere/micro-examples/part6/basic/common"
	"github.com/entere/micro-examples/part6/basic/config"
	"github.com/entere/micro-examples/part6/user-srv/handler"
	"github.com/entere/micro-examples/part6/user-srv/model"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-plugins/config/source/grpc"

	proto "github.com/entere/micro-examples/part6/user-srv/proto/user"
)

type userCfg struct {
	common.AppCfg
}

var (
	appName = "user-srv"
	cfg     = &userCfg{}
)

func main() {
	initCfg()
	// New Service
	microReg := etcd.NewRegistry(registryOptions)
	service := micro.NewService(
		micro.Name("io.github.entere.srv.user"),
		micro.Registry(microReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(micro.Action(func(context *cli.Context) {
		model.Init()
		handler.Init()
	}))

	// Register Handler
	proto.RegisterUserHandler(service.Server(), new(handler.User))

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
