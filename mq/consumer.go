/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 下午4:33
* */
package mq

import (
	"log"
)

// 消费者

// 开始监听队列获取消息
// qName: 队列名称
// cName: 消费者名称
// callback: 处理函数
func StartConsume(qName,cName string,callback func(msg []byte)error){
	// 1.通过channel.Consume 获得消息的信道
	if MQChann == nil {
		log.Println("chann is not")
		return
	}
	deliveries, err := MQChann.Consume(
		qName, // 队列名称
		cName, // 消费者名称
		true,  // 收到消息自动回复ack
		false, // 不是队列的唯一消费者
		false, // nolocal 是不支持的
		false, // 等待返回信息
		nil,
	)
	if err != nil {
		log.Println(err.Error())
		return
	}

	// 2.循环获取队列的消息
	done := make(chan bool)
	go func() {
		for {
			select {
			case data :=<-deliveries :
				// 3.调用callback处理消息
				err := callback(data.Body)
				if err != nil {
					// TODO:当前消息写道另一个队列,用于异常情况的重试
				}
			}
		}
	}()

	<-done
}