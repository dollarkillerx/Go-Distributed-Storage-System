/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 下午4:04
* */
package test

import "testing"

func TestOnes(t *testing.T)  {
	// 存储类型(表示文件存到哪里)
	type storeType int

	const (
		_ storeType = iota
		// StoreLocal : 节点本地
		storeLocal
		// StoreCeph : Ceph集群
		storeCeph
		// StoreOSS : 阿里OSS
		storeOSS
		// StoreMix : 混合(Ceph及OSS)
		storeMix
		// StoreAll : 所有类型的存储都存一份数据
		storeAll
	)
	t.Log(storeOSS)
}
