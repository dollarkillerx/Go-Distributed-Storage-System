/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:09
* */
package config

import (
	"Go-Distributed-Storage-System/err"
	"encoding/json"
	"os"
)

// 基础配置
type configStruct struct {
	Host      string `json:"host"`
	MySQLDsn  string `json:"mysql_dsn"`
	PwdSalt   string `json:"pwd_salt"`
	RedisHost string `json:"redisHost"`
}

// oss配置
type OssCon struct {
	OssBucket   string `json:"oss_bucket"`
	OssEndpoint string `json:"oss_endpoint"`
	// oss访问key
	OssAccessKeyID string `json:"oss_access_key_id"`
	// oss访问key secret
	OssAccessKeySecret string `json:"oss_access_key_secret"`
}

//Mq配置
type MqConf struct {
	// AsyncTransferEnable:是否开启文件异步转移(默认同步)
	AsyncTransferEnable bool `json:"async_transfer_enable"`
	// RabbitUrl: rabbitMq服务入口
	RabbitURL string `json:"rabbit_url"` // 这个才是主要的,其他的配置的本业务的配置
	// TransExchangeName: 用于文件transfer的交换机
	TransExchangeName string `json:"trans_exchange_name"`
	// TransOSSQueueName:oss转移队列名称
	TransOSSQueueName string `json:"trans_oss_queue_name"`
	// TransOSSErrQueueName: oss转移失败后写入另一队列的队列名
	TransOSSErrQueueName string `json:"trans_oss_err_queue_name"`
	// TransOSSRoutingKey: routingKey
	TransOSSRoutingKey string `json:"trans_oss_routing_key"`
}

var (
	ConfigBase *configStruct
	OssConfig  *OssCon
	MqConfig   *MqConf
)

func init() {
	getBase()
	getOss()
	getMq()
}

func getBase() {
	path := "./config.json"
	file, e := os.Open(path)
	err.ErrPanic(e)
	decoder := json.NewDecoder(file)
	ConfigBase = &configStruct{}
	e = decoder.Decode(ConfigBase)
	err.ErrPanic(e)
}

func getOss() {
	path := "./ossconf.json"
	file, e := os.Open(path)
	err.ErrPanic(e)
	decoder := json.NewDecoder(file)
	OssConfig = &OssCon{}
	e = decoder.Decode(OssConfig)
	err.ErrPanic(e)
}

func getMq() {
	path := "./mqconfig.json"
	file, e := os.Open(path)
	err.ErrPanic(e)
	decoder := json.NewDecoder(file)
	MqConfig := &MqConf{}
	e = decoder.Decode(MqConfig)
	err.ErrPanic(e)
}
