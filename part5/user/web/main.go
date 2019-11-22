package main

import (
	"fmt"
	"github.com/entere/micro-examples/part5/basic"
	"github.com/entere/micro-examples/part5/basic/config"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"

	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"

	"github.com/entere/micro-examples/part5/user/web/handler"
	"github.com/micro/go-micro/web"
)

func main() {
	basic.Init()
	micReg := etcd.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		web.Name("io.github.entere.web.user"),
		web.Version("latest"),
		web.Registry(micReg),
		web.Address(":8088"),
	)

	// initialise service
	if err := service.Init(
		web.Action(func(context *cli.Context) {
			handler.Init()

		}),
	); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.HandleFunc("/user/login", handler.Login)
	service.HandleFunc("/user/logout", handler.Logout)

	// register call handler
	// service.HandleFunc("/user/call", handler.UserCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
func registryOptions(ops *registry.Options) {
	etcdCfg := config.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}