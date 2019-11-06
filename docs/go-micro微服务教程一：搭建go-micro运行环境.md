# go-micro微服务教程一：搭建go-micro运行环境
Go Micro 是一个用于微服务开发的框架。[官方文档]（https://micro.mu/docs/go-micro.html）
Go Micro提供分布式系统开发的核心库，包含RPC与事件驱动的通信机制。micro的设计哲学是可插拔的架构理念，她提供可快速构建系统的组件，并且可以根据自身的需求剥离默认实现并自行定制。
本章节的目的是让大家最快速搭建好`go-micro` 运行环境

| 名称        | 说明   |
|:--------   | :--------------------  | 
|OS    | macOS Mojave 10.14.6 | 
|Golang     | go1.13 darwin/amd64 | 
|Micro       |micro version 1.11.3|
|G111MODULE |on|

### 安装micro和go-micro
```shell script
go get -u -v github.com/micro/go-micro
# 安装micro工具包，这个安装是可选项，micro提供了一系列的工具来帮助我们更好的使用go-micro。
go get -u -v github.com/micro/micro 
cd $GOPATH/src/github.com/micro/micro
go build -o micro main.go
sudo cp micro /usr/local/bin/
```

### 安装 protoc

[官方地址](https://github.com/protocolbuffers/protobuf/releases)
```shell script
# 源码安装，下载最新版，目前是 v3.10.1

wget https://github.com/protocolbuffers/protobuf/releases/download/v3.10.1/protobuf-all-3.10.1.tar.gz
tar zxvf protobuf-all-3.10.1.tar.gz
cd protobuf-3.10.1/
./autogen.sh
./configure
make
make install
protoc -h

# 或者
git clone https://github.com/google/protobuf
cd protobuf
./autogen.sh
./configure
make
make check
sudo make install
```

### 安装 Go 版本的 protobuf 编译器 protoc-gen-go
```shell script
go get -u -v github.com/golang/protobuf/protoc-gen-go
```



###  在 github 上创建 micro-examples 项目  
在开始之前，我们先在[github](http://github.com/)上创建 [`micro-examples`](http://github.com/entere/micro-examples) 项目，并使用 `go mod init` 初始化生成`go.mod` 文件，我们用 Go Modules 管理 Go 模块依赖  
```shell script
cd micro-examples
go mod init github.com/entere/micro-examples
```

### goland配置

我们用 goland 这个ide开发go应用

去掉"Use GOPATH that's defined in system environment"前面的√
选中Enable Go Modules(vgo) intergration，并在proxy中输入 https://goproxy.io

下一节，我们将会创建一个简单的微服务[go-micro微服务教程二：创建微服务]()



