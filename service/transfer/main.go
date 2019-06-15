/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 下午4:55
* */
package transfer

import (
	"Go-Distributed-Storage-System/config"
	"Go-Distributed-Storage-System/mq"
	"Go-Distributed-Storage-System/store/oss"
	"bufio"
	"encoding/json"
	"log"
	"os"
)

func ProcessTransfer(msg []byte) error {
	// 1.解析msg
	pubData := mq.TransferData{}
	err := json.Unmarshal(msg, pubData)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	// 2.根据临时存储文件,创建文件句柄
	file, err := os.Open(pubData.CurLocation)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	// 3.通过文件句柄将文件写入到oos
	err = oss.Bucket().PutObject(pubData.DestLocation, bufio.NewReader(file))
	if err != nil {
		log.Println(err.Error())
		return err
	}
	// 4.更改文件的存储路径

	return nil
}

func main() {
	log.Println("开始监听转移任务队列...")
	mq.StartConsume(
		config.MqConfig.TransOSSQueueName,
		"transfer_oss",
		ProcessTransfer,
	)
}