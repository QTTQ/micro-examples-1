# go-micro微服务教程三：实现用户信息微服务




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

```shell script
protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto

```

## 编写用户模型服务

```shell script
vi handler/user.go
```
调整QueryUserByID方法：

```golang

package handler

import (
	"context"

	user "github.com/entere/micro-examples/user/srv/proto/user"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) QueryUserByID(ctx context.Context, req *user.QueryRequest, rsp *user.QueryResponse) error {

	userID := req.UserId
	if userID == "" {
		userID = "xxxxxx"
	}
	userInfo := &user.UserInfo{
		UserId:    userID,
		Nickname:  "匿名",
		Mobile: "138********",
		AvatarUrl: "avatar.jpg",
		Gender:    1,
	}
	rsp.Code = 200
	rsp.Error = nil
	rsp.Data = userInfo

	return nil
}

```

## 修改main.go 注册微服务

```shell script

vi main.go

```



```golang

package main

import (
	"github.com/entere/micro-examples/user/srv/handler"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"

	user "github.com/entere/micro-examples/user/srv/proto/user"
)

func main() {
	// 新建服务
	service := micro.NewService(
		micro.Name("io.github.entere.srv.user"),
		micro.Version("latest"),
	)

	// 初始化服务

	service.Init()

	// 注册服务
	user.RegisterUserHandler(service.Server(), new(handler.User))

	// 启动服务
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}


```


## 运行测试

```shell script
make build
./user-srv

micro  call io.github.entere.srv.user User.QueryUserByID '{"user_id":"543987654"}'

```
返回结果

```json
{
        "code": 200,
        "data": {
                "user_id": "543987654",
                "nickname": "匿名",
                "mobile": "138********",
                "avatar_url": "avatar.jpg",
                "gender": 1
        }
}


```