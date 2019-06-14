# Go-Distributed-Storage-System
Go Distributed Storage System

### 基础架构
- 最小化上传接口
- 账户系统
- 秒传
- 分块上传
- 断点续传
- Ceph分布式文件系统
- OSS公有云
- 微服务化
- 采用GRPC内部通讯

### 开发环境
- Ubuntu18.4
- Go: 12.5
- IDE: Goland
- TestAPi: Postman
- SQL: Mysql
- NoSQL: Redis
- MQ: RabbitMQ
- RPC: GRPC
- Docker/Kubernets
- Ceph

### 敏捷开发
- v1版本
    - 简单的文件上传下载浏览服务
    - 用户系统
- V2版本
    - 实现秒传功能
    - 实现文件分块上传
    - 实现断点续传
    
    
### 错误码
- 0051 模板错误
- 0040 用户输入参数错误
- 0050 服务器内部错误
- 0052 数据库查询错误
- 0200 模板消息
- 0053 内部输入错误
- 0054 存储数据出错