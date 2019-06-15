/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 下午2:53
* */
package test

import (
	"Go-Distributed-Storage-System/store/oss"
	oss2 "github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	filename := "hes.text"
	ossPath := "oss/" + filename // 不能以/开头
	file, _ := os.Open(filename)
	err := oss.Bucket().PutObject(ossPath, file)
	if err != nil {
		t.Log("上传oss失败")
		return
	}
}

// 临时下载授权
func TestDownUrl(t *testing.T) {
	filepath := "oss/hes.text"
	url, e := oss.Bucket().SignURL(filepath, oss2.HTTPGet, 3600)
	if e != nil {
		t.Log(e.Error())
		return
	}
	t.Log(url)
}
