#  go-micro微服务教程三：实现api
 
 本节课的主要目的是实现api网关
 
## micro new 指令生成 api 骨架代码
 
```shell script
cd micro-examples  
micro new user/api --type=api --alias=user --namespace=io.github.entere --gopath=false
```

### 修改生成原型

修改 proto/user/user.proto 为以下内容
```proto3
syntax = "proto3";

package io.github.entere.api.user;

import "github.com/micro/go-micro/api/proto/api.proto";

service User {
	rpc Info(go.api.Request) returns (go.api.Response) {}
}

```

编译生成原型

```shell script
cd user/api
protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto
```

### 修改 client/user.go

```golang
package client

import (
	"context"

	user "github.com/entere/micro-examples/user/srv/proto/user"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server"
)

type userKey struct{}

// FromContext retrieves the client from the Context
func UserFromContext(ctx context.Context) (user.UserService, bool) {
	c, ok := ctx.Value(userKey{}).(user.UserService)
	return c, ok
}

// Client returns a wrapper for the UserClient
func UserWrapper(service micro.Service) server.HandlerWrapper {
	client := user.NewUserService("io.github.entere.srv.user", service.Client())

	return func(fn server.HandlerFunc) server.HandlerFunc {
		return func(ctx context.Context, req server.Request, rsp interface{}) error {
			ctx = context.WithValue(ctx, userKey{}, client)
			return fn(ctx, req, rsp)
		}
	}
}

```

### 修改 handler/user.go

```golang
package handler

import (
	"context"
	"encoding/json"
	"github.com/micro/go-micro/util/log"

	"github.com/entere/micro-examples/user/api/client"
	userSrv "github.com/entere/micro-examples/user/srv/proto/user"
	api "github.com/micro/go-micro/api/proto"
	"github.com/micro/go-micro/errors"
)

type User struct{}

func extractValue(pair *api.Pair) string {
	if pair == nil {
		return ""
	}
	if len(pair.Values) == 0 {
		return ""
	}
	return pair.Values[0]
}

// User.Info is called by the API as /user/call with post body {"name": "foo"}
func (e *User) Info(ctx context.Context, req *api.Request, rsp *api.Response) error {
	log.Log("Received User.Info request")

	// extract the client from the context
	userClient, ok := client.UserFromContext(ctx)
	if !ok {
		return errors.InternalServerError("io.github.entere.api.user.info", "user client not found")
	}

	// make request
	response, err := userClient.QueryUserByID(ctx, &userSrv.QueryRequest{
		UserId: extractValue(req.Post["user_id"]),
	})
	if err != nil {
		return errors.InternalServerError("io.github.entere.api.user.user.call", err.Error())
	}

	b, _ := json.Marshal(response)

	rsp.StatusCode = 200
	rsp.Body = string(b)

	return nil
}


```


### 修改 main.go

```golang

package main

import (
	"github.com/micro/go-micro/util/log"

	"github.com/entere/micro-examples/user/api/client"
	"github.com/entere/micro-examples/user/api/handler"
	"github.com/micro/go-micro"

	user "github.com/entere/micro-examples/user/api/proto/user"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("io.github.entere.api.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		// create wrap for the User srv client
		micro.WrapHandler(client.UserWrapper(service)),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}



```

### 编译运行

```shell script

cd usr/api
make build
./user-api

```

### 启动micro api 网关

重新运行 srv api 并启动 micro api， 程序调用顺序：micro api->api->srv

```shell script
# 启动 srv
cd micro-examples/user/srv
./user-srv

#启动 api 
cd micro-examples/user/api
./user-api

#启动 micro-api api网关，实现restful api接口支持
cd micro-examples/
micro api --namespace=io.github.entere.api --handler=api

```

### 测试

curl 请求接口：

```shell script
curl -H "Content-Type:application/x-www-form-urlencoded" -X POST -d "user_id=f546e78a46284765" http://localhost:8080/user/info

```

返回结果：
```json

{
    "code": 200,
    "data": {
        "user_id": "f546e78a46284765",
        "nickname": "匿名",
        "mobile": "138********",
        "avatar_url": "avatar.jpg",
        "gender": 1
    }
}

```
