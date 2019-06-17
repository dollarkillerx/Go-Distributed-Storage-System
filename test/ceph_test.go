/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 上午10:51
* */
package test

import (
	"Go-Distributed-Storage-System/store/ceph"
	"fmt"
	"gopkg.in/amz.v1/s3"
	"io/ioutil"
	"testing"
)

func TestCeph(t *testing.T) {
	bucker := ceph.GetCephBucker("testbucket1")
	// 创建一个新的bucket
	err := bucker.PutBucket(s3.PublicRead) // 参数权限
	if err != nil {
		t.Log(err.Error())
	}
	// 查询bucket下面制定条件的object keys
	result, err := bucker.List("", "", "", 100) // 最多返回100条数据
	if err != nil {
		t.Log(err.Error())
	}
	fmt.Printf("object keys%v \n", result)
	// 上传新对象
	err = bucker.Put("/a.txt", []byte("just for test"), "octet-steam", s3.PublicRead)
	/**
	bucker.Put("/a.txt",[]byte("just for test"),"octet-steam",s3.PublicRead)
	put参数
	1.在文件系统的绝对路径
	2.内容
	3.context-type
	4.权限
	*/
	if err != nil {
		t.Log(err.Error())
	}
	// 查询下名制定条件的object keys
	result, err = bucker.List("", "", "", 100) // 最多返回100条数据
	if err != nil {
		t.Log(err.Error())
	}
	fmt.Printf("object keys%v \n", result)
}

func TestUp(t *testing.T) {
	filename := "test.txt"
	bucker := ceph.GetCephBucker("userfile")
	cephPath := "/ceph/" + filename
	bytes, _ := ioutil.ReadFile(filename)
	bucker.Put(cephPath, bytes, "octet-stream", s3.PublicRead)
}

func TestDow(t *testing.T) {
	bucker := ceph.GetCephBucker("userfile")
	cephPath := "/ceph/test.txt"
	data, _ := bucker.Get(cephPath)
	ioutil.WriteFile("hello.text", data, 00666)
}
