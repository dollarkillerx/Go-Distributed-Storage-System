/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 下午4:06
* */
package mq

// 生产者

import (
	"Go-Distributed-Storage-System/config"
	"errors"
	"github.com/streadway/amqp"
	"log"
)

var (
	MQConn  *amqp.Connection
	MQChann *amqp.Channel
	e       error
)

func init() {
	MQConn, e = amqp.Dial(config.MqConfig.RabbitURL)
	if e != nil {
		panic(e.Error())
	}
	MQChann, e = MQConn.Channel()
	if e != nil {
		log.Println(e.Error())
	}
}

// 发布消息
func Publish(exchange, routingKey string, msg []byte) error {
	// 1.判断channel是否是正常的
	if MQChann == nil {
		return errors.New("chan not ex")
	}
	// 2.执行消息投递
	err := MQChann.Publish(
		exchange,   // 交换机
		routingKey, // key
		false,      // 如钩没有合适的队列会丢弃此消息
		false,      // 过期参数
		amqp.Publishing{
			ContentType: "text/plain", // 明文发送
			Body:        msg,
		},
	)
	return err
}
