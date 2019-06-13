/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-13
* Time: 上午10:35
* */
package defs

import (
	"errors"
	"os"
	"sync"
)

// FileMeta 文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt int
}

var fileMetas sync.Map

// UpdateFileMeate: 更新/新增文件元信息
func UpdateFileMeta(fmeta *FileMeta) {
	fileMetas.Store(fmeta.FileSha1,fmeta)
}

// GetFileMeta:通过sha1key获取文件元信息对象
func GetFileMeta(fileSha1 string) (*FileMeta,error) {
	value, ok := fileMetas.Load(fileSha1)
	if ok {
		meta,ok := value.(*FileMeta)
		if ok {
			return meta,nil
		}
	}
	return nil,errors.New("bad")
}

// 获取批量的文件元信息列表
func GetLastFileMetas(conunt int) []*FileMeta {
	metas := make([]*FileMeta, 10) // 这个地方就有点蛋疼了,go 说 map就不应该有len 和 size gg
	fileMetas.Range(func(key, value interface{}) bool {
		meta := value.(*FileMeta)
		metas = append(metas, meta)
		return true
	})
	// 这边对于[]的排序要实现接口的三个方法我就难得写了
	return metas[0:conunt]
}

func FileDeleteHandler(hash string) error {
	value, ok := fileMetas.Load(hash)
	if ok != true {
		return errors.New("load file not exits")
	}
	meta := value.(*FileMeta)
	path := meta.Location
	// 删除map中的数据
	fileMetas.Delete(hash)
	// 删除真实文件  其实这里可以开一个协程 异步删除
	err := os.Remove(path)
	if err != nil {
		return errors.New("load file not exits")
	}
	return nil
}