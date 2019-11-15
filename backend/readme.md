# SmartLocker

本项目为某智能储物柜的服务端程序

文件目录结构说明

```
.
├── Dockerfile docker构建文件
├── Dockerfile.CN 修改了镜像源的docker构建文件
├── Makefile 构建脚本
├── VERSION 项目版本
├── cmd 可执行文件
│   ├── seed 随机数据生成器
│   │   └── seed.go
│   └── server 服务端
│       ├── router 路由
│       │   ├── api.go
│       │   ├── v1 第1版
│       │   │   ├── article.go
│       │   │   ├── face.go
│       │   │   ├── middleware 认证
│       │   │   │   └── jwt.go
│       │   │   ├── user.go
│       │   │   └── wrapper.go
│       │   └── v1.go
│       └── server.go
├── config 配置文件处理包
│   ├── config.go
│   └── scheme.go
├── config_sqlite.yaml 使用sqlite的设置文件
├── config_mysql.yaml 使用mysql的设置文件
├── docker-startup.sh docker容器内的启动脚本
├── e 错误处理包
│   ├── code.go
│   └── msg.go
├── go.mod 项目依赖记录
├── go.sum 项目依赖记录
├── logger 日志记录包
│   └── logger.go
├── model 数据库交互模型
│   ├── cabinet.go
│   ├── db.go
│   ├── locker.go
│   └── user.go
├── readme.md 本文档
├── resources [预留]
├── service
│   ├── article 储物格信息
│   │   └── article.go
│   ├── auth 认证
│   │   └── auth.go
│   ├── cabinet 储物柜信息
│   │   └── cabinet.go
│   ├── user 用户信息
│   │   └── user.go
|   └── cache 缓存
|       ├── redis.go
│       └── cabinet.go
└── util 辅助工具包
    ├── hash.go
    └── img.go

```

