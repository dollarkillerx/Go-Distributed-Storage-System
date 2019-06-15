/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 下午1:06
* */
package oss

import (
	"Go-Distributed-Storage-System/config"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"log"
)

var (
	OssCli *oss.Client
	e error
)

func init() {
	ossCline()
}

func ossCline() {
	OssCli, e = oss.New(config.OssConfig.OssEndpoint,config.OssConfig.OssAccessKeyID,config.OssConfig.OssAccessKeySecret)
	if e != nil {
		panic(e.Error())
	}
}

// Bucket: 获取bucket存储空间
func Bucket() *oss.Bucket {
	bucket, err := OssCli.Bucket(config.OssConfig.OssBucket)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return bucket
}
