/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-11
* Time: 下午10:25
* */
package defs

type Message struct {
	Code int `json:"code"`
	Resp *Resp `json:"resp"`
}

type Resp struct {
	Message string `json:"message,omitempty"`
	Code string `json:"code"`
	Data interface{} `json:"data,omitempty"`
}

