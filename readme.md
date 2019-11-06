# SmartLocker

本项目为某智能储物柜的服务端程序

[![Build Status](https://travis-ci.org/panpaul/SmartLocker.svg?branch=master)](https://travis-ci.org/panpaul/SmartLocker)

文件目录结构说明

```
.
├── Dockerfile 用于部署的Dockerfile
├── Dockerfile.CN 修改了镜像的Dockerfile
├── LICENSE
├── Makefile
├── VERSION
├── cmd 本目录存放了最终生成的应用
│   ├── seed [TODO]测试数据生成器
│   │   └── seed.go
│   └── server 服务器本体
│       ├── router 路由
│       │   ├── api.go
│       │   ├── v1
│       │   │   └── wrapper.go
│       │   └── v1.go
│       └── server.go
├── config 一个简单的yaml配置读取器
│   ├── config.go
│   └── scheme.go
├── config.yaml 配置文件
├── docker-startup.sh 用于容器启动时的startup脚本
├── e 处理错误的包
│   ├── code.go
│   └── msg.go
├── go.mod
├── go.sum
├── logger 记录日志的包
│   └── logger.go
├── model 数据库模型
│   ├── cabinet.go
│   ├── db.go
│   ├── locker.go
│   └── user.go
├── readme.md
├── resources
├── service 本项目涉及的几个service
│   ├── article article即物品，处理储物相关细节
│   │   └── article.go
│   ├── cabinet 储物柜相关
│   │   └── cabinet.go
│   └── user 用户相关
│       └── user.go
└── util 辅助工具包
    └── hash.go
```

