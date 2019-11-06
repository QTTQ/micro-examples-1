# go-micro微服务教程二：创建微服务

本示例的目的是用go-micro完成一个简单微服务（srv）
Micro有提供代码生成器指令 `new`，它可以新建服务模板代码，把基本所需的目录结构建好，省去大家挖坑的时间。这个指令我们在[上一课]()安装过。

在开发之前，我们需要先理清 `micro api`, `api`, `srv` 三者之间的关系，这也是初学都容易迷惑的地方。
> * micro api 网关层，本质上是一个http协议的网关接口，它会把动态路由到转到后台服务中。默认监听8080端口。命令启动：micro api --handler=xx --namespace=xx
> * api api服务层,会把请求的参数组装成srv想要的形式，如果服务非常简单，可以不要api， 直接micro api -> srv
> * srv 后端微服务层，主要向所有内部服务提供接口

正常用 micro api->api->srv这样的调用
调用流程: 发起http请求到网关, 然后进入api层执行对应的handler, handler里面封装创建 具体微服务的rpc client 并向后端发起rpc请求.



### 生成微服务骨架代码

```shell script
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

```shell script
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


### 生成原型
在定义好原型后我们得使用`protoc`及`micro`的插件编译它，`micro`插件可以帮助生成`Go Micro`需要的原型文件
```shell script
protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/user/user.proto
```

### 编写服务
我们先不调整micro new 生成的骨架代码，只修改部分 import 让程序正常跑起来

### 运行程序

```shell script
cd micro-examples/usr/srv
make build
./user-srv 

# 查看是否启动
micro list services

```


### 测试

方式一：micro call 命令直接调用 srv 

```shell script
micro call io.github.entere.srv.user User.Call '{"name":"entere"}'
```
结果:
```json
{
    "msg": "Hello entere"
}
```

方式二：启用micao api 网关调用 srv

```shell script
micro api --handler=rpc --namespace=io.github.entere.srv
curl http://localhost:8080/user/call?name=entere
```

结果:
```json
{
    "msg": "Hello entere"
}
```

### 总结

本节课，我们基本没有写任何代码，简单的让go-micro 的服务层 srv 运行了起来，下一节，我们将结合api调用srv层