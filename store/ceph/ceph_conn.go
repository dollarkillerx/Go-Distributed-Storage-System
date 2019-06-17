/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 上午10:39
* */
package ceph

import (
	"gopkg.in/amz.v1/aws"
	"gopkg.in/amz.v1/s3"
)

var (
	CephConn *s3.S3
)

func init() {
	// 初始化ceph信息
	auth := aws.Auth{
		AccessKey: "",
		SecretKey: "",
	}
	region := aws.Region{
		Name:                 "default",
		EC2Endpoint:          "http://127.0.0.1:9080",
		S3Endpoint:           "http://127.0.0.1:9080",
		S3BucketEndpoint:     "",
		S3LocationConstraint: false, // 没有区域限制
		S3LowercaseBucket:    false, // bucket没有大小写限制
		Sign:                 aws.SignV2,
	}
	// 创建🔓s3类型连接
	CephConn = s3.New(auth, region)
}

// 获取指定Bucker
func GetCephBucker(bucket string) *s3.Bucket {
	return CephConn.Bucket(bucket)
}
