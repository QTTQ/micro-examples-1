# micro-examples
go-micro 微服务示例

本示例的目的是用go-micro完成一个用户系统微服务



## 创建micro-examples 项目

在开始之前，我们先在[github](http://github.com/)上创建 [`micro-examples`](http://github.com/entere/micro-examples) 项目，并使用 `go mod init` 初始化生成`go.mod` 文件，我们用 Go Modules 管理 Go 模块依赖 

```bash
cd micro-examples
go mod init github.com/entere/micro-examples

```



> 本节主要目标：实现根据用户id查询用户信息的接口，例如：

```bash
$ curl -s http://localhost:50734/user/info?user_id=1
# 返回结果

{
    "code": 200,
    "data": {
        "user_id": "1",
        "nickname": "匿名",
        "mobile": "13810895860",
        "avatar_url": "entere.png\n",
        "created_at": 1572310240
    }
}

```


实现用户服务，用户服务分为两层，web层（web）与服务层（srv）前者提供http接口，后者向web提供RPC服务。

web服务主要向用户提供http查询用户信息接口：

srv服务主要向所有内部服务提供RPC服务：根据userId查询用户

我们先来写服务层srv,后写web层web





## srv

Micro有提供代码生成器指令[new]()，它可以新建服务模板代码，把基本所需的目录结构建好，省去大家挖坑的时间。这个指令我们在[上一课]()安装过。

### micro new 指令生成 srv 骨架代码

```bash
cd micro-examples  
micro new user/srv --type=srv --alias=user --namespace=io.github.entere --gopath=false
```

我们看一个micro new 的参数说明：

|配置指令|作用|默认值|说明|
|---|---|---|---|
|--namespace|服务命令空间 |go.micro|可以根据自己的域名定义合适的空间前缀|
|--type|服务类型|srv|目前支持4种服务类型，分别是api、fnc(function)、srv(service)、web。|
|--alias|指定别名|声明则必填|使用单词，不要带任何标点符号，名称对Micro路由机制影响很大|
|--plugin|使用哪些插件|声明则必填|需要自选插件时使用|
|--gopath|是否使用GOPATH作为代码路径|true或者false|使用go modules 可以设置为false|
|--fqdn|服务定义域，API需要通过该域找到服务|默认是使用服务的命令空间加上类型再加上别名|服务定义域|

模板生成在srv目录，进入user/srv目录

```bash
$ cd user/srv
$ tree
```
生成的骨架代码结构如下：

```
├── main.go
├── plugin.go
├── handler
│   └── user.go
├── subscriber
│   └── user.go
├── proto/user
│   └── user.proto
├── Dockerfile
├── Makefile
└── README.md

```

protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto




## 在 proto/user/user.proto中定义User原型

修改user.proto中的内容

```proto
syntax = "proto3";

package io.github.entere.srv.user;
// import "google/protobuf/timestamp.proto";
service User {
	rpc QueryUserByID(QueryRequest) returns (QueryResponse) {}
}



message UserInfo {
    string user_id = 1;
    string nickname = 2;
    string mobile = 3;
    string avatar_url = 4;
    int32 gender = 5;
    uint32 created_at = 6;
    uint32 updated_at = 7;
    // google.protobuf.Timestamp created_at = 6;
    // google.protobuf.Timestamp updated_at = 7;
}

message Error {
    string msg = 1;
    string info = 2;
}

message QueryRequest {
    string user_id = 1;
    
}

message QueryResponse {
    uint32 code = 1;
    Error error = 2;
    UserInfo data = 3;
}

```

上面我们定义了User服务的基本原型结构，包含用户信息UserInfo，请求QueryRequest与响应结构QueryResponse，还定义了查询用户的方法QueryUserByID。

下面我们生成类型与服务方法：

```bash
$ protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto

```

## 编写用户模型服务

```bash
$ vi handler/user.go
```
调整QueryUserByID方法：

```golang

package handler

import (
	"context"

	"github.com/entere/micro-examples/user/srv/dao"
	"github.com/entere/micro-examples/user/srv/service"

	user "github.com/entere/micro-examples/user/srv/proto/user"
)

// User ...
type User struct{}

// QueryUserByID ...
func (e *User) QueryUserByID(ctx context.Context, req *user.QueryRequest, rsp *user.QueryResponse) error {
	userID := req.UserId
	userInfo := &user.UserInfo{
		UserId:    userID,
        Nickname:  "匿名",
        Mobile: "138********"
		AvatarUrl: "http://entere.github.io/images/avatar.jpg",
		Gender:    1,
	}
	rsp.Code = 200
	rsp.Error = nil
	rsp.Data = userInfo

	return nil
}


```

## 修改main.go 注册微服务

```bash

$ vi main.go

```

新建服务
初始化服务
注册服务
启动服务



```golang

package main

import (
	"github.com/entere/micro-examples/user/base"
	"github.com/entere/micro-examples/user/srv/handler"
	user "github.com/entere/micro-examples/user/srv/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	base.Init()
	// New Service
	service := micro.NewService(
		micro.Name("io.github.entere.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	// micro.RegisterSubscriber("io.github.entere.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	// micro.RegisterSubscriber("io.github.entere.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

```


## 运行测试

```bash
$ make build
$ micro  call io.github.entere.srv.user User.QueryUserByID '{"userId":"1"}'
```
 返回结果

```json
{
	"code": 200,
	"data": {
		"user_id": "1",
		"nickname": "匿名",
		"mobile": "138********",
		"avatar_url": "entere.png\n",
		"created_at": 1572310240
	}
}

```

