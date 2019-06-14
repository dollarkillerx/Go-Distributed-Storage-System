/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-14
* Time: 下午9:37
* */
package redis

import (
	"Go-Distributed-Storage-System/config"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"time"
)

var (
	RedisConn *redis.Pool
)

func init() {
	RedisConn = newRedisPool()
}

// newRedisPool:创建redis连接池
func newRedisPool()*redis.Pool {
	return &redis.Pool{
		MaxIdle:50, // 池中的最大空闲连接数
		MaxActive:30, // 最大连接数
		IdleTimeout:300*time.Second,// 超时回收
		Dial: func() (conn redis.Conn, e error) {
			// 1. 打开连接
			dial, e := redis.Dial("tcp", config.ConfigBase.RedisHost)
			if e != nil {
				fmt.Println(e.Error())
				return nil,e
			}
			// 2. 访问认证
			//dial.Do("AUTH",redisPassword)
			return dial,nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error { // 定时检查连接是否可用
			// time.Since(t) 获取离现在过了多少时间
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}
}
