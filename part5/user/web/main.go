package main

import (
	"fmt"
	"github.com/entere/micro-examples/part5/basic"
	"github.com/entere/micro-examples/part5/basic/common"
	"github.com/entere/micro-examples/part5/basic/config"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/config/source/grpc"
	"net/http"

	"github.com/micro/go-micro/registry/etcd"
	"github.com/micro/go-micro/util/log"

	"github.com/entere/micro-examples/part5/user/web/handler"
	"github.com/micro/go-micro/web"
)

type userCfg struct {
	common.AppCfg
}

var (
	appName = "student-web"
	cfg     = &userCfg{}
)

func main() {
	initCfg()
	micReg := etcd.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		web.Name("io.github.entere.web.student"),
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

	service.HandleFunc("/student/login", handler.Login)
	service.HandleFunc("/student/logout", handler.Logout)
	// service.HandleFunc("/student/info", handler.Info)

	infoHandler := http.HandlerFunc(handler.Info)
	service.Handle("/student/info", handler.JWTAuthWrapper(infoHandler))

	// register call handler
	// service.HandleFunc("/student/call", handler.UserCall)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	etcdCfg := &common.Etcd{}
	err := config.C().App("etcd", etcdCfg)
	if err != nil {
		panic(err)
	}
	ops.Addrs = []string{fmt.Sprintf("%s:%d", etcdCfg.Host, etcdCfg.Port)}
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
