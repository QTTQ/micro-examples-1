package main

import (
	"fmt"
	handler2 "github.com/entere/micro-examples/part5/auth/srv/handler"
	model2 "github.com/entere/micro-examples/part5/auth/srv/model"
	"github.com/entere/micro-examples/part5/auth/srv/proto/auth"
	basic2 "github.com/entere/micro-examples/part5/basic"
	config2 "github.com/entere/micro-examples/part5/basic/config"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/etcd"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	// 初始化配置、数据库等信息
	basic2.Init()

	// 使用etcd注册
	micReg := etcd.NewRegistry(registryOptions)

	// 新建服务
	service := micro.NewService(
		micro.Name("io.github.entere.srv.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// 服务初始化
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化handler
			model2.Init()
			// 初始化handler
			handler2.Init()
		}),
	)

	// 注册服务
	io_github_entere_srv_auth.RegisterAuthServiceHandler(service.Server(), new(handler2.Service))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := config2.GetEtcdConfig()
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.GetHost(), etcdCfg.GetPort())}
}
