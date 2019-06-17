/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 下午11:00
* */
package main

import (
	"Go-Distributed-Storage-System/service/account/container"
	"Go-Distributed-Storage-System/service/account/proto"
	"github.com/micro/go-micro"
	"log"
	"time"
)

func main() {
	// 创建一个service
	service := micro.NewService(
		micro.Name("go.micro.service"),        // 服务的名称
		micro.RegisterTTL(time.Second*10),     // 10s检查等待时间
		micro.RegisterInterval(time.Second*5), // 服务美5s发一次心跳
	)
	// 初始化解析命令行参数
	service.Init()

	// 注册服务到rpc
	proto.RegisterUserServiceHandler(service.Server(), new(container.User))
	if err := service.Run(); err != nil {
		log.Println(err.Error())
	}
}
