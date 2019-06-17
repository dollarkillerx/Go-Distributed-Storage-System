/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-15
* Time: 下午4:00
* */
package mq

// 存储类型(表示文件存到哪里)
type StoreType int

const (
	_ StoreType = iota
	// StoreLocal : 节点本地
	StoreLocal
	// StoreCeph : Ceph集群
	StoreCeph
	// StoreOSS : 阿里OSS
	StoreOSS
	// StoreMix : 混合(Ceph及OSS)
	StoreMix
	// StoreAll : 所有类型的存储都存一份数据
	StoreAll
)

// 定义MQ消息体
type TransferData struct {
	FileHash      string
	CurLocation   string // 临时文件地址
	DestLocation  string // 目标文件地址
	DestStoreType StoreType
}
